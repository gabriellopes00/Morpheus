import { Mailer } from '@/application/mail/mailer'
import { AccountData } from '@/domain/account'
import { MailQueue, MailQueueProcess } from '@/ports/mail-queue'
import Queue from 'bull'
import { Sentry } from '../utils/sentry'

const { REDIS_PORT, REDIS_HOST } = process.env

export class BullMailQueue implements MailQueue {
  private readonly queue: Queue.Queue<MailQueueProcess> = null

  constructor(private readonly mailer: Mailer) {
    this.queue = new Queue<MailQueueProcess>('mail-queue', {
      defaultJobOptions: { attempts: 5, backoff: 1000 * 60 },
      redis: { port: Number(REDIS_PORT), host: REDIS_HOST }
    })
  }

  public async addProcess(process: MailQueueProcess): Promise<void> {
    try {
      await this.queue.add(process, { removeOnComplete: true })
    } catch (error) {
      console.error(error)
    }
  }

  public async process(): Promise<void> {
    return await this.queue.process(async (job, done) => {
      try {
        // await this.mailer.sendEmail(job.data)
        await job.data.process()
        done(null)
      } catch (err) {
        done(err)
      }
    })
  }

  public handleFailedJobs(): void {
    this.queue.on('failed', async (job, err) => {
      Sentry.captureMessage(`Job id: ${job.id} failed with error: ${err.message}`)
      Sentry.captureException(err)
    })

    this.queue.on('completed', async (job, _) => {
      console.error(`Job id: ${job.id} completed successfully`)
    })
  }
}

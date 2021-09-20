import { Mailer } from '@/application/mail/mailer'
import { AccountData } from '@/domain/account'
import { MailQueue } from '@/ports/mail-queue'
import Queue from 'bull'

const { REDIS_PORT, REDIS_HOST } = process.env

export class BullMailQueue implements MailQueue {
  private readonly queue: Queue.Queue<AccountData> = null

  constructor(private readonly mailer: Mailer) {
    this.queue = new Queue<AccountData>('mail-queue', {
      defaultJobOptions: { attempts: 5, backoff: 1000 * 60 },
      redis: { port: Number(REDIS_PORT), host: REDIS_HOST }
    })
  }

  public async addProcess(data: AccountData): Promise<void> {
    try {
      await this.queue.add(data, { removeOnComplete: true })
    } catch (error) {
      console.error(error)
    }
  }

  public async process(): Promise<void> {
    return await this.queue.process(async (job, done) => {
      try {
        await this.mailer.sendEmail(job.data)
        done(null)
      } catch (err) {
        done(err)
      }
    })
  }

  public handleFailedJobs(): void {
    this.queue.on('failed', async (job, err) => {
      console.error(`Job id: ${job.id} failed with error: ${err.message}`)
    })

    this.queue.on('completed', async (job, _) => {
      console.error(`Job id: ${job.id} completed successfully`)
    })
  }
}

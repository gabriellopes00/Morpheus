import { MailQueue } from '@/ports/mail-queue'
import Queue from 'bull'

const { REDIS_PORT, REDIS_HOST, REDIS_PASSWORD } = process.env

export class BullMailQueue implements MailQueue {
  private readonly queue = new Queue<() => Promise<void>>('mail-queue', {
    defaultJobOptions: { attempts: 5, backoff: 1000 * 60 * 5 },
    redis: {
      port: Number(REDIS_PORT),
      host: REDIS_HOST,
      password: REDIS_PASSWORD
    }
  })

  public async addProcess(process: () => Promise<void>): Promise<void> {
    await this.queue.add(process, { removeOnComplete: true })
  }

  public async process(): Promise<void> {
    this.queue.process(async (job, done) => {
      try {
        await job.data()
        done(null)
      } catch (err) {
        done(err)
      }
    })
  }

  public handleFailedJobs(err: Error): void {
    this.queue.on('failed', async (job, err) => {
      console.error(`Job id: ${job.id} failed with error: ${err.message}`)
    })
  }
}

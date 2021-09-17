import 'dotenv/config'
import 'module-alias/register'
import { Mailer } from '../application/mail/mailer'
import { Queue } from '../application/message-queue/queue'
import { createBullMailQueue } from './factories/mail-queue'

const mailQueue = createBullMailQueue()
const mailer = new Mailer(mailQueue)

;(async () => {
  try {
    // start processing email submissions
    mailQueue.process()
    mailQueue.handleFailedJobs()

    // start consuming message queue
    const queue = new Queue(mailer)
    await queue.consume()
  } catch (error) {
    console.error(error)
    process.exit(1)
  }
})()

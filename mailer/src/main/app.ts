import 'dotenv/config'
import 'module-alias/register'
import { Mailer } from '../application/mail/mailer'
import { MessageQueue } from '../application/message-queue/message-queue'
import { BullMailQueue } from '../framework/mail-queue/bull-mail-queue'
import { NodemailerMailProvider } from '../framework/mail-provider/nodemailer-mail-provider'
import { Sentry } from '../framework/utils/sentry'
;(async () => {
  try {
    const mailer = new Mailer(new NodemailerMailProvider())
    const mailQueue = new BullMailQueue(mailer)

    // start consuming message queue
    const queue = new MessageQueue(mailQueue)
    await queue.connect()
    await queue.consume()

    // start processing email submissions
    mailQueue.handleFailedJobs()
    await mailQueue.process()
  } catch (error) {
    Sentry.captureException(error)
    console.error(error)
    process.exit(1)
  }
})()

import 'dotenv/config'
import 'module-alias/register'
import { Mailer } from './mail/mailer'
import { Queue } from './queue/queue'

const mailer = new Mailer()

;(async () => {
  try {
    const queue = new Queue(mailer)
    await queue.consume()
  } catch (error) {
    console.error(error)
    process.exit(1)
  }
})()

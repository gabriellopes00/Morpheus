import 'dotenv/config'
import { Mailer } from '../services/mail/mailer'
import { Queue } from '../services/queue/queue';

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
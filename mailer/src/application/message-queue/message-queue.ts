import { AccountData } from '@/domain/account'
import { MailQueue } from '@/ports/mail-queue'
import amqp from 'amqplib/callback_api'

const { RABBITMQ_PORT, RABBITMQ_USER, RABBITMQ_VHOST, RABBITMQ_HOST, RABBITMQ_PASS } = process.env

const connectionOptions: amqp.Options.Connect = {
  port: Number(RABBITMQ_PORT),
  hostname: RABBITMQ_HOST,
  vhost: RABBITMQ_VHOST,
  username: RABBITMQ_USER,
  password: RABBITMQ_PASS
}

export class MessageQueue {
  private conn: amqp.Connection = null

  constructor(private readonly mailQueue: MailQueue) {}

  public async consume(): Promise<void> {
    amqp.connect(connectionOptions, (err, connection) => {
      if (err) {
        console.log(err)
        throw err
      }

      this.conn = connection

      this.conn.createChannel((err, ch) => {
        if (err) throw err

        const queues = ['account_created', 'account_deleted']

        for (const queue of queues) {
          ch.assertQueue(queue, { durable: true })
          ch.prefetch(1)

          ch.consume(queue, async msg => await this.handleMessage(msg, queue), { noAck: true })
        }
      })
    })
  }

  private async handleMessage(msg: amqp.Message, queue: string): Promise<void> {
    const account: AccountData = JSON.parse(msg.content.toString())

    switch (queue) {
      case 'account_created':
        this.mailQueue.addProcess(account)
        break

      default:
        break
    }
  }
}

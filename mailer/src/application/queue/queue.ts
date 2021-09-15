import amqp from 'amqplib/callback_api'
import { Account } from '@/domain/account'
import { Mailer } from '../mail/mailer'

const { RABBITMQ_PORT, RABBITMQ_USER, RABBITMQ_VHOST, RABBITMQ_HOST, RABBITMQ_PASS } = process.env

const connectionOptions: amqp.Options.Connect = {
  port: Number(RABBITMQ_PORT),
  hostname: RABBITMQ_HOST,
  vhost: RABBITMQ_VHOST,
  username: RABBITMQ_USER,
  password: RABBITMQ_PASS
}

export class Queue {
  private conn: amqp.Connection = null

  constructor(private readonly mailer: Mailer) {}

  public async connect(): Promise<Error> {
    amqp.connect(connectionOptions, (err, connection) => {
      if (err) return err
      this.conn = connection
    })

    return null
  }

  public async consume(): Promise<void> {
    if (this.conn === null) throw new Error('Amqp connection unavailable')

    this.conn.createChannel((err, ch) => {
      if (err) throw err

      const queues = ['account_created', 'account_deleted']

      for (const queue of queues) {
        ch.assertQueue(queue, { durable: true })
        ch.prefetch(1)

        ch.consume(
          queue,
          async msg => {
            const account: Account = JSON.parse(msg.content.toString())

            switch (queue) {
              case 'account_created':
                this.mailer.sendWelcomeEmail(account)
                break

              default:
                break
            }
          },
          { noAck: true }
        )
      }
    })
  }
}

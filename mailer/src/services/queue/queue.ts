import amqp from 'amqplib/callback_api'
import { Account } from '../../entities/account'
import { Mailer } from '../mail/mailer'

const { RABBITMQ_PORT, RABBITMQ_USER, RABBITMQ_VHOST, RABBITMQ_HOST, RABBITMQ_PASS } = process.env

const connectionOptions: amqp.Options.Connect = {
  port: Number(RABBITMQ_PORT),
  hostname: RABBITMQ_HOST,
  vhost: RABBITMQ_VHOST,
  username: RABBITMQ_USER,
  password: RABBITMQ_PASS,
}

export class Queue {
  constructor(private readonly mailer: Mailer){}

  public async consume(): Promise<void> {
    amqp.connect(connectionOptions, (err, conn) => {
    if (err) throw err

    conn.createChannel((err, ch) => {
      if (err) throw err
      
      const queue = 'account_created'

      ch.assertQueue(queue, { durable: true })
      ch.prefetch(1)
      
      ch.consume(queue, async (msg) => {
        const account: Account = JSON.parse(msg.content.toString())
        console.log(account);
        
        this.mailer.sendWelcomeEmail(account)
      }, { noAck: true })
      
    })
  })
  }
}
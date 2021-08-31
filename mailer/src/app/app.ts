import 'dotenv/config'
import amqp from 'amqplib/callback_api'

const {
  RABBITMQ_PORT,
  RABBITMQ_USER,
  RABBITMQ_VHOST,
  RABBITMQ_HOST,
  RABBITMQ_PASS,
} = process.env

const options: amqp.Options.Connect = {
  port: Number(RABBITMQ_PORT),
  hostname: RABBITMQ_HOST,
  vhost: RABBITMQ_VHOST,
  username: RABBITMQ_USER,
  password: RABBITMQ_PASS,
}

interface Account {
  id: string
  name: string
  email: string
  avatar_url: string
}

amqp.connect(options, (err, conn) => {
  if (err) {
    console.error(err)
    return
  }

  conn.createChannel((err, ch) => {
    if (err) {
      console.error(err)
      return
    }

    const queue = 'account_created'

    ch.assertQueue(queue, { durable: false })
    ch.prefetch(1)
    ch.consume(queue, (msg) => console.log(msg.content.toString()), {
      noAck: true,
    })
  })
})

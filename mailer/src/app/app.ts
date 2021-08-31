import 'dotenv/config'
import amqp from 'amqplib/callback_api'
import nodemailer from 'nodemailer'

const {
  RABBITMQ_PORT,
  RABBITMQ_USER,
  RABBITMQ_VHOST,
  RABBITMQ_HOST,
  RABBITMQ_PASS,

  SMTP_USER,
  SMTP_PASS,
  SMTP_PORT,
  SMTP_HOST,
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

    ch.assertQueue(queue, { durable: true })
    ch.prefetch(1)
    ch.consume(queue, async (msg) => {
      const account: Account = JSON.parse(msg.content.toString())

      const transporter = nodemailer.createTransport({
        host: SMTP_HOST,
        port: Number(SMTP_PORT),
        auth: {
          user: SMTP_USER,
          pass: SMTP_PASS
        }
      })

      await transporter.sendMail({
        from: "gabriellopes@morpheus.io", 
        to: account.email,
        text: `Well Come to Morpheus, ${account.name}<${account.email}>! Lorem ipsum dolor sit amet consectetur adipisicing elit. Maiores quasi eveniet dolore id recusandae accusantium molestias soluta expedita labore quia porro sint atque ducimus, voluptas veniam laudantium quaerat quo libero! Lorem ipsum dolor sit amet consectetur adipisicing elit. Officiis veniam neque minus ipsum deleniti exercitationem maxime dolore est temporibus pariatur laborum nemo libero, modi, porro sunt. Maiores esse error doloremque.`
      })

    }, {
      noAck: true,
    })
  })
})

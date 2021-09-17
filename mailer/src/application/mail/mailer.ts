import { MailQueue } from '@/ports/mail-queue'
import { readFile } from 'fs'
import { compile } from 'handlebars'
import { createTransport } from 'nodemailer'

const { SMTP_USER, SMTP_PASS, SMTP_PORT, SMTP_HOST } = process.env
const TEMPLATE_PATH = __dirname + '/templates/welcome.hbs'

export class Mailer {
  constructor(private readonly mailQueue: MailQueue){}

  private readonly transporter = createTransport({
    host: SMTP_HOST,
    port: Number(SMTP_PORT),
    auth: {      user: SMTP_USER,      pass: SMTP_PASS    }
  })

  public async sendEmail<T = any>(data: T): Promise<void> {
    readFile(TEMPLATE_PATH, { encoding: 'utf-8' }, async (err, file) => {
      if (err) throw err

      const template = compile(file)
      const html = template(data)

      const mailProcess = async (): Promise<void> => {
        await this.transporter.sendMail({
          sender: 'gabriellopes@morpheus.io',
          from: 'gabriellopes@morpheus.io',
          to: 'data.email',
          subject: 'Welcome to Morpheus',
          text: `Welcome to Morpheus, ${'account.name'}!

          Your account was successfully created using the email ${'account.email'}.

          Lorem ipsum dolor sit amet consectetur adipisicing elit. Necessitatibus facere numquam officia odio accusantium soluta excepturi reiciendis ipsam fugiat? Ipsum atque culpa expedita omnis itaque animi, inventore dolores laboriosam a. Lorem ipsum dolor sit amet consectetur adipisicing elit. Porro omnis debitis explicabo similique aliquam corrupti doloribus qui sequi tempora consequuntur delectus nostrum, ea vero beatae, facere doloremque commodi non totam?

          Lorem ipsum dolor, sit amet consectetur adipisicing elit. Modi quisquam quis, facere quibusdam quod amet perferendis tempora deserunt recusandae voluptates! Fuga reprehenderit natus ducimus rerum fugiat iure sunt iusto voluptas.`,
          html,
          encoding: 'utf-8'
        })
      }

      return await this.mailQueue.addProcess(mailProcess)
    })
  }
}

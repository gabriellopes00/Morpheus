import { AccountData } from '@/domain/account'
import { EmailProvider } from '@/ports/email-provider'
import { readFile } from 'fs'
import { compile } from 'handlebars'

const TEMPLATE_PATH = __dirname + '/templates/welcome.hbs'

export class Mailer {
  constructor(private readonly mailProvider: EmailProvider) {}

  public async sendEmail(data: AccountData): Promise<void> {
    readFile(TEMPLATE_PATH, { encoding: 'utf-8' }, async (err, file) => {
      if (err) throw err

      const template = compile(file)
      const html = template(data)

      await this.mailProvider.sendMail({
        sender: 'gabriellopes@morpheus.io',
        from: 'gabriellopes@morpheus.io',
        to: data.email,
        subject: 'Welcome to Morpheus',
        text: `Welcome to Morpheus, ${data.name}!

          Your account was successfully created using the email ${data.email}.

          Lorem ipsum dolor sit amet consectetur adipisicing elit. Necessitatibus facere numquam officia odio accusantium soluta excepturi reiciendis ipsam fugiat? Ipsum atque culpa expedita omnis itaque animi, inventore dolores laboriosam a. Lorem ipsum dolor sit amet consectetur adipisicing elit. Porro omnis debitis explicabo similique aliquam corrupti doloribus qui sequi tempora consequuntur delectus nostrum, ea vero beatae, facere doloremque commodi non totam?

          Lorem ipsum dolor, sit amet consectetur adipisicing elit. Modi quisquam quis, facere quibusdam quod amet perferendis tempora deserunt recusandae voluptates! Fuga reprehenderit natus ducimus rerum fugiat iure sunt iusto voluptas.`,
        html
      })
    })
  }
}

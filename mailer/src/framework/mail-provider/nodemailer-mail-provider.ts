import { EmailProvider, EmailProviderProps } from '@/ports/email-provider'
import { createTransport } from 'nodemailer'
import { MailOptions } from 'nodemailer/lib/json-transport'

const { SMTP_USER, SMTP_PASS, SMTP_PORT, SMTP_HOST } = process.env

export class NodemailerMailProvider implements EmailProvider {
  private readonly transporter = createTransport({
    host: SMTP_HOST,
    port: Number(SMTP_PORT),
    auth: { user: SMTP_USER, pass: SMTP_PASS }
  })

  public async sendMail(props: EmailProviderProps): Promise<void> {
    const mailOptions: MailOptions = {
      sender: props.sender,
      from: props.from,
      to: props.to,
      subject: props.subject,
      text: props.text,
      html: props.html,
      encoding: 'utf-8'
    }

    this.transporter.sendMail(mailOptions, (err, _) => {
      if (err) throw err
    })
  }
}

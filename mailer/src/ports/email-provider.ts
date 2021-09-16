export interface EmailProvider {
  sendMail(from: string, to: string, subject: string, text: string, html: string): Promise<void>
}

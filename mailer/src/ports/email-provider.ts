export interface EmailProviderProps {
  sender: string
  from: string
  to: string
  subject: string
  text: string
  html: string
}

export interface EmailProvider {
  sendMail(props: EmailProviderProps): Promise<void>
}

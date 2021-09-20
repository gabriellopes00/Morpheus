import { Mailer } from '@/application/mail/mailer'
import { BullMailQueue } from '@/framework/mail-queue/bull-mail-queue'

export function createBullMailQueue(mailer: Mailer): BullMailQueue {
  return new BullMailQueue(mailer)
}

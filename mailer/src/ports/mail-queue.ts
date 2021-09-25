import { AccountData } from '@/domain/account'

export interface MailQueueProcess {
  process(): Promise<void>
}

export interface MailQueue {
  addProcess(data: MailQueueProcess): Promise<void>
  process(): Promise<void>
}

import { AccountData } from '@/domain/account'

export interface MailQueue {
  addProcess(data: AccountData): Promise<void>
  process(): Promise<void>
}

import {BullMailQueue} from '@/framework/mail-queue/bull-mail-queue'

export function createBullMailQueue(): BullMailQueue {
  return new BullMailQueue()
}

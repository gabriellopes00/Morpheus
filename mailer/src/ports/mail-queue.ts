export interface MailQueue {
  addProcess(process: () => Promise<void>): Promise<void>
  process(): void
}

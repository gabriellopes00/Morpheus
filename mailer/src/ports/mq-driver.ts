export interface MQDriver {
  consumeMessage(queue: string): void
}

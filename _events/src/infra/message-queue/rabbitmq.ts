import { CreateAccount } from '@/modules/accounts/usecases/create-account'
import { DeleteAccount } from '@/modules/accounts/usecases/delete-account'
import amqp from 'amqplib'
const { RABBITMQ_PORT, RABBITMQ_USER, RABBITMQ_VHOST, RABBITMQ_HOST, RABBITMQ_PASS } = process.env

const connectionOptions: amqp.Options.Connect = {
  port: Number(RABBITMQ_PORT),
  hostname: RABBITMQ_HOST,
  vhost: RABBITMQ_VHOST,
  username: RABBITMQ_USER,
  password: RABBITMQ_PASS
}

export interface AccountMessage {
  id: string
  name: string
  email: string
  document: string
  avatarUrl: string
  phoneNumber: string
  gender: string
  birthDate: Date
  createdAt: Date
  updatedAt: Date
}

export class MessageQueue {
  private conn: amqp.Connection = null
  private readonly queues = ['account_created_events', 'account_deleted_events']

  constructor(
    private readonly createAccount: CreateAccount,
    private readonly deleteAccount: DeleteAccount
  ) {}

  public async connect(): Promise<void> {
    this.conn = await amqp.connect(connectionOptions)
  }

  public async consume(): Promise<void> {
    if (this.conn == null) throw new Error('Amqp connection unavailable')

    const channel = await this.conn.createChannel()

    for (const queue of this.queues) {
      channel.assertQueue(queue, { durable: true })
      channel.bindQueue(
        queue,
        'accounts_ex',
        queue === this.queues[0] ? 'account_created' : 'account_deleted'
      )
      channel.prefetch(1)
      channel.consume(queue, async (msg: amqp.Message) => await this.handleMessage(msg, queue), {
        noAck: true
      })
    }
  }

  private async handleMessage(msg: amqp.Message, queue: string): Promise<void> {
    const account: AccountMessage = JSON.parse(msg.content.toString())

    switch (queue) {
      case 'account_created_events':
        await this.createAccount.execute(account.id)
        break

      case 'account_deleted_events':
        await this.deleteAccount.execute(account.id)
        break

      default:
        break
    }
  }
}

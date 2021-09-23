import { AccountData } from '@/domain/account'
import { MailQueue } from '@/ports/mail-queue'
import amqp from 'amqplib'
import { DeleteAccount } from '../accounts/delete-account'
import { SaveAccount } from '../accounts/save-account'

const { RABBITMQ_PORT, RABBITMQ_USER, RABBITMQ_VHOST, RABBITMQ_HOST, RABBITMQ_PASS } = process.env

const connectionOptions: amqp.Options.Connect = {
  port: Number(RABBITMQ_PORT),
  hostname: RABBITMQ_HOST,
  vhost: RABBITMQ_VHOST,
  username: RABBITMQ_USER,
  password: RABBITMQ_PASS
}

export class MessageQueue {
  private conn: amqp.Connection = null
  private readonly queues = ['account_created_mail', 'account_deleted_mail']

  constructor(
    private readonly mailQueue: MailQueue,
    private readonly saveAccount: SaveAccount,
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
      channel.consume(queue, async msg => await this.handleMessage(msg, queue), { noAck: true })
    }
  }

  private async handleMessage(msg: amqp.Message, queue: string): Promise<void> {
    const account: AccountData = JSON.parse(msg.content.toString())

    switch (queue) {
      case 'account_created_mail':
        await this.mailQueue.addProcess(account)
        await this.saveAccount.save(account)
        break

      case 'account_deleted_mail':
        // await this.mailQueue.addProcess(account)
        await this.deleteAccount.delete(account.id)
        break

      default:
        break
    }
  }
}

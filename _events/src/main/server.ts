import 'dotenv/config'
import 'module-alias/register'
import { UidGenerator } from '@/infra/crypto/uuid-generator'
import { PostgresDataSource } from '@/infra/db/data-source'
import { PgAccountRepository } from '@/infra/db/repositories/account-repository'
import { MessageQueue } from '@/infra/message-queue/rabbitmq'
import { CreateAccount } from '@/modules/accounts/usecases/create-account'
import { DeleteAccount } from '@/modules/accounts/usecases/delete-account'
;(async () => {
  try {
    const repo = new PgAccountRepository(PostgresDataSource)

    const saveAccount = new CreateAccount(new UidGenerator(), repo)
    const deleteAccount = new DeleteAccount(repo)

    // start consuming message queues
    const queue = new MessageQueue(saveAccount, deleteAccount)
    await queue.connect()
    await queue.consume()
  } catch (error) {
    console.error(error)
    process.exit(1)
  }
})()

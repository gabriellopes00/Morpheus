import { Account } from '@/domain/account'
import { AccountRepository } from '@/ports/repositories/account-repository'
import { getRepository } from 'typeorm'
import { AccountEntity } from '../entities/account'

export class PgAccountRepository implements AccountRepository {
  public async save(data: Account): Promise<void> {
    const repository = getRepository(AccountEntity)
    await repository.save(repository.create(data))
  }

  public async findByEmail(email: string): Promise<Account> {
    const repository = getRepository(AccountEntity)
    const { id, name, email: _email, created_at } = await repository.findOne({ where: { email } })
    return null
  }
}

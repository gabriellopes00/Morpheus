import { Account } from '@/modules/accounts/domain/account'
import { CreateAccountRepository } from '@/modules/accounts/repositories/create-account-repository'
import { DeleteAccountRepository } from '@/modules/accounts/repositories/delete-account-repository'
import { DataSource } from 'typeorm'
import { AccountEntity } from '../entities/account-entity'

export class PgAccountRepository implements CreateAccountRepository, DeleteAccountRepository {
  constructor(private readonly dataSource: DataSource) {}

  public async create(data: Account): Promise<void> {
    const repository = this.dataSource.getRepository(AccountEntity)
    await repository.save(repository.create(data))
  }

  public async delete(referencedId: string): Promise<void> {
    const repository = this.dataSource.getRepository(AccountEntity)
    await repository.delete(referencedId)
  }
}

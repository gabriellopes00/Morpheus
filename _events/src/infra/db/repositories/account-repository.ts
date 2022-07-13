import { Account } from '@/modules/accounts/domain/account'
import { FindRepository, SaveRepository } from '@/shared/repositories'
import { DataSource } from 'typeorm'
import { AccountEntity } from '../entities/account-entity'

export class PgAccountRepository
implements Partial<FindRepository<Account>>, Partial<SaveRepository<Account>> {
  constructor(private readonly dataSource: DataSource) {}

  public async findBy(key: keyof Account, value: string): Promise<Account> {
    const repository = this.dataSource.getRepository(AccountEntity)
    const entity = await repository.findOneBy({ [key]: value })
    return new Account({ ...entity }, entity.id)
  }

  public async save(data: Account): Promise<void> {
    const repository = this.dataSource.getRepository(AccountEntity)
    await repository.save(repository.create({ ...data }))
  }
}

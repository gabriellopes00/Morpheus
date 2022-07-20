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
    if (!entity) return null
    else return new Account({ ...entity }, entity.id)
  }

  public async save(data: Account): Promise<void> {
    const repository = this.dataSource.getRepository(AccountEntity)
    await repository.save(
      repository.create({
        id: data.id,
        name: data.name,
        email: data.email,
        gender: data.gender as 'male',
        password: data.password,
        birthDate: data.birthDate,
        avatarUrl: data.avatarUrl,
        document: data.document,
        createdAt: data.createdAt
      })
    )
  }
}

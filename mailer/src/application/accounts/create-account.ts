import { Account, AccountData } from '@/domain/account'
import { AccountRepository } from '@/ports/repositories/account-repository'
import { UUIDGenerator } from '@/ports/uuid-generator'

export class CreateAccount {
  constructor(
    private readonly uuidGenerator: UUIDGenerator,
    private readonly repository: AccountRepository
  ) {}

  public async create(data: AccountData): Promise<void> {
    const account = new Account(data, this.uuidGenerator.generate())
    return await this.repository.create(account)
  }
}

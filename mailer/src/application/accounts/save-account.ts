import { Account, AccountData } from '@/domain/account'
import { AccountRepository } from '@/ports/repositories/account-repository'

export class SaveAccount {
  constructor(private readonly repository: AccountRepository) {}

  public async save(data: AccountData): Promise<void> {
    const account = new Account(data)
    return await this.repository.save(account)
  }
}

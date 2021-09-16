import { AccountRepository } from '@/ports/repositories/account-repository'

export class DeleteAccount {
  constructor(private readonly repository: AccountRepository) {}

  public async delete(accountId: string): Promise<void> {
    return await this.repository.delete(accountId)
  }
}

import { Account } from '@/modules/accounts/domain/account'
import { CreateAccountRepository } from '@/modules/accounts/repositories/create-account-repository'
import { DeleteAccountRepository } from '@/modules/accounts/repositories/delete-account-repository'

export class MockAccountRepository implements CreateAccountRepository, DeleteAccountRepository {
  private readonly _accounts: Account[] = []

  get rows() {
    return this._accounts
  }

  public async create(data: Account): Promise<void> {
    if (
      this._accounts.some(a => {
        return a.referencedId === data.referencedId || a.id === data.id
      })
    ) {
      throw new Error('inexistent referenced_id')
    }

    this._accounts.push(data)
  }

  public async delete(referencedId: string): Promise<void> {
    if (!this._accounts.some(a => a.referencedId === referencedId)) {
      throw new Error('inexistent referenced_id')
    }

    this._accounts.splice(
      this._accounts.findIndex(a => a.referencedId === referencedId),
      1
    )
  }
}

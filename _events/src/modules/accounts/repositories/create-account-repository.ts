import { Account } from '../domain/account'

export interface CreateAccountRepository {
  create(data: Account): Promise<void>
}

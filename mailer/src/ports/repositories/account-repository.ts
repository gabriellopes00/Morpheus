import { Account } from '@/domain/account'

export interface AccountRepository {
  create(data: Account): Promise<void>
  findByEmail(email: string): Promise<Account>
}

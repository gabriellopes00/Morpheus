import { Account } from '@/domain/account'

export interface AccountRepository {
  save(data: Account): Promise<void>
  findByEmail(email: string): Promise<Account>
}

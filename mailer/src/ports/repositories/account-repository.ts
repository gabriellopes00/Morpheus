import { Account } from '@/domain/account'

export interface AccountRepository {
  save(data: Account): Promise<void>
  delete(id: string): Promise<void>
}

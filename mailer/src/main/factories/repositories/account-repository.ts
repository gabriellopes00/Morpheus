import { PgAccountRepository } from '@/framework/database/repositories/account-repository'
import { AccountRepository } from '@/ports/repositories/account-repository'

export function createAccountRepository(): AccountRepository {
  return new PgAccountRepository()
}

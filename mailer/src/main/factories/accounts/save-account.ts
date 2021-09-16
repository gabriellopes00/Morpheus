import { SaveAccount } from '@/application/accounts/save-account'
import { createAccountRepository } from '../repositories/account-repository'

export function createCreateAccount(): SaveAccount {
  return new SaveAccount(createAccountRepository())
}

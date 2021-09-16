import { CreateAccount } from '@/application/accounts/create-account'
import { createAccountRepository } from '../repositories/account-repository'
import { createUUIDGenerator } from '../utils/uuid-generator'

export function createCreateAccount(): CreateAccount {
  return new CreateAccount(createUUIDGenerator(), createAccountRepository())
}

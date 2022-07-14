import { FindRepository, SaveRepository } from '@/shared/repositories'
import { Account, AccountData } from '../domain/account'
import { v4 as uuid } from 'uuid'
import { Hasher } from '@/core/infra/hasher'

export interface CreateAccountParams extends AccountData {}

export interface TokenData {
  email: string
}

export class CreateAccount {
  constructor(
    private readonly repository: FindRepository<Account> & SaveRepository<Account>,
    private readonly hasher: Hasher
  ) {}

  public async execute(params: CreateAccountParams): Promise<Account | Error> {
    const { name, document, email, gender, avatarUrl, birthDate, password } = params

    const account = new Account(
      { avatarUrl, gender, name, email, document, password, birthDate },
      uuid()
    )

    let existentCredentials = await this.repository.findBy('email', email)
    if (existentCredentials) {
      return new Error('Email do usuário já registrado')
    }

    existentCredentials = await this.repository.findBy('document', document)
    if (existentCredentials) {
      return new Error('CPF do usuário já registrado')
    }

    account.password = await this.hasher.generate(account.password)
    await this.repository.save(account)

    return account
  }
}

import { FindRepository, SaveRepository } from '@/shared/repositories'
import { Account, AccountData } from '../domain/account'

export interface CreateAccountParams extends AccountData {}

export interface TokenData {
  email: string
}

export class CreateAccount {
  constructor(private readonly repository: FindRepository<Account> & SaveRepository<Account>) {}

  public async execute(params: CreateAccountParams): Promise<Account | Error> {
    const { name, document, email, gender, avatarUrl, birthDate, password } = params

    const account = new Account(
      { avatarUrl, gender, name, email, document, password, birthDate },
      crypto.randomUUID()
    )

    let existentCredentials = await this.repository.findBy('email', email)
    if (existentCredentials) {
      return new Error('Email do usuário já registrado')
    }

    existentCredentials = await this.repository.findBy('document', document)
    if (existentCredentials) {
      return new Error('CPF do usuário já registrado')
    }

    await this.repository.save(account)

    return account
  }
}

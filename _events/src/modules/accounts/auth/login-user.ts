import { Encrypter } from '@/core/infra/encrypter'
import { Hasher } from '@/core/infra/hasher'
import { Account } from '@/modules/accounts/domain/account'
import { FindRepository } from '@/shared/repositories'
import { env } from 'process'

export interface LoginAccountResult {
  accessToken: string
  account: Account
}

export interface LoginAccountParams {
  email: string
  password: string
}

export class LoginAccount {
  constructor(
    private readonly repository: FindRepository<Account>,
    private readonly encrypter: Encrypter,
    private readonly hasher: Hasher
  ) {}

  public async execute(params: LoginAccountParams): Promise<LoginAccountResult | Error> {
    const account = await this.repository.findBy('email', params.email)
    if (!account) return new Error('E-mail, login ou senha inválido(s)')

    if (!account.password) return new Error('Conta de usuário ainda inativa')

    const isValidPassword = await this.hasher.compare(params.password, account.password)
    if (!isValidPassword) return new Error('E-mail, login ou senha inválido(s)')

    const accessToken = await this.encrypter.encrypt(
      { id: account.id, isRefreshToken: false },
      env.ACCESS_TOKEN_EXPIRATION
    )

    return { account, accessToken }
  }
}

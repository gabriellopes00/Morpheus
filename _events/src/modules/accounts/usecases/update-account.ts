import { FindRepository, SaveRepository } from '@/shared/repositories'
import { Account, AccountGender } from '../domain/account'

export interface UpdateAccountParams {
  id: string
  name?: string
  avatarUrl?: string
  gender?: AccountGender
}

export class UpdateAccount {
  constructor(private readonly repository: FindRepository<Account> & SaveRepository<Account>) {}

  public async execute(params: UpdateAccountParams): Promise<Account | Error> {
    const { id, name, avatarUrl, gender } = params

    const account = await this.repository.findBy('id', id)
    if (!account) return new Error('Usuário não encontrado')

    if (name) account.name = name
    if (avatarUrl) account.avatarUrl = avatarUrl
    if (gender) account.gender = gender

    await this.repository.save(account)
    return account
  }
}

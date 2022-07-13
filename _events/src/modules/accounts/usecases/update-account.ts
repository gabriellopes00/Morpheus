import { FindRepository, SaveRepository } from '@/shared/repositories'
import { Account } from '../domain/account'

export interface UpdateAccountParams {
  id: string
  name?: string
  avatarUrl?: string
}

export class UpdateAccount {
  constructor(private readonly repository: FindRepository<Account> & SaveRepository<Account>) {}

  public async execute(params: UpdateAccountParams): Promise<Account | Error> {
    const { id, name, avatarUrl } = params

    const account = await this.repository.findBy('id', id)
    if (!account) return new Error('Usuário não encontrado')

    if (name) account.name = name
    if (avatarUrl) account.name = avatarUrl

    await this.repository.save(account)
    return account
  }
}

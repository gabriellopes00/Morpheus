import { FindRepository } from '@/shared/repositories'
import { Account } from '../domain/account'

export interface FindAccountByIdParams {
  id: string
}

export class FindAccountById {
  constructor(private readonly repository: FindRepository<Account>) {}

  public async execute(params: FindAccountByIdParams): Promise<Account | Error> {
    return await this.repository.findBy('id', params.id)
  }
}

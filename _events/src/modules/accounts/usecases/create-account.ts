import { UUIDGenerator } from '@/core/infra/uuid-generator'
import { Account } from '../domain/account'
import { CreateAccountRepository } from '../repositories/create-account-repository'

export class CreateAccount {
  constructor(
    private readonly uuidGenerator: UUIDGenerator,
    private readonly repository: CreateAccountRepository
  ) {}

  public async execute(referencedId: string): Promise<void> {
    const account = new Account({ referencedId }, this.uuidGenerator.generate())
    return this.repository.create(account)
  }
}

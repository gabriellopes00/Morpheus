import { DeleteAccountRepository } from '../repositories/delete-account-repository'

export class DeleteAccount {
  constructor(private readonly repository: DeleteAccountRepository) {}

  public async execute(referencedId: string): Promise<void> {
    return this.repository.delete(referencedId)
  }
}

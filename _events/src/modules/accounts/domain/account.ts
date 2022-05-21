import { Entity } from '@/shared/entity'

export interface AccountData {
  referencedId: string
}

export class Account extends Entity<AccountData> {
  constructor(data: AccountData, id: string) {
    super(data, id)
  }

  public get referencedId(): string {
    return this.data.referencedId
  }
}

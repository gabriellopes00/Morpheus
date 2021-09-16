import { Entity } from '@/shared/entity'

export interface AccountData {
  id: string
  name: string
  email: string
}

export class Account extends Entity<AccountData> {
  constructor(data: AccountData) {
    super(data, data.id)
  }

  get name() {
    return this.data.name
  }

  set name(name: string) {
    this.data.name = name
  }

  get email() {
    return this.data.email
  }

  set email(email: string) {
    this.data.email = email
  }
}

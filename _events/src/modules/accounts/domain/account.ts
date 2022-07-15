import { Entity } from '@/shared/entity'

export type AccountGender = 'male' | 'female' | 'unspecified'

export interface AccountData {
  name: string
  email: string
  document: string
  password: string
  avatarUrl: string
  gender: AccountGender
  birthDate: string
}

export class Account extends Entity<AccountData> {
  constructor(data: AccountData, id: string) {
    super(data, id)
  }

  public get name(): string {
    return this.data.name
  }

  public get email(): string {
    return this.data.email
  }

  public get document(): string {
    return this.data.document
  }

  public get avatarUrl(): string {
    return this.data.avatarUrl
  }

  public get password(): string {
    return this.data.password
  }

  public set password(value: string) {
    this.data.password = value
  }

  public set name(value: string) {
    this.data.name = value
  }

  public set avatarUrl(value: string) {
    this.data.avatarUrl = value
  }

  public get gender(): AccountGender {
    return this.data.gender
  }

  public set gender(value: AccountGender) {
    this.data.gender = value
  }

  public get birthDate(): string {
    return this.data.birthDate
  }
}

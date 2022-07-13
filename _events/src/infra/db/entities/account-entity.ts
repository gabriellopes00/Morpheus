import { AccountData, AccountGender } from '@/modules/accounts/domain/account'
import { Column, CreateDateColumn, Entity, PrimaryColumn } from 'typeorm'

@Entity({ name: 'accountss' })
export class AccountEntity implements AccountData {
  @PrimaryColumn({ type: 'uuid' })
  public id: string

  @Column({ type: 'varchar' })
  public name: string

  @Column({ type: 'varchar' })
  public email: string

  @Column({ type: 'varchar' })
  public document: string

  @Column({ type: 'varchar' })
  public password: string

  @Column({ type: 'varchar' })
  public avatarUrl: string

  @Column({ type: 'varchar' })
  public gender: AccountGender

  @Column({ type: 'varchar' })
  public birthDate: string

  @CreateDateColumn({ name: 'created_at', type: 'timestamptz' })
  public createdAt: Date

  @CreateDateColumn({ name: 'updated_at', type: 'timestamptz' })
  public updatedAt: Date
}

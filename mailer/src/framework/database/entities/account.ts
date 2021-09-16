import { AccountData } from '@/domain/account'
import { Column, CreateDateColumn, Entity, PrimaryColumn, UpdateDateColumn } from 'typeorm'

@Entity('account')
export class AccountEntity implements AccountData {
  @PrimaryColumn()
  public id: string

  @Column()
  public name: string

  @Column()
  public email: string

  @CreateDateColumn()
  public created_at: Date
}

import { AccountData } from '@/domain/account'
import { Column, CreateDateColumn, Entity, PrimaryColumn } from 'typeorm'

@Entity('accounts')
export class AccountEntity implements AccountData {
  @PrimaryColumn()
  public id: string

  @Column()
  public name: string

  @Column()
  public email: string

  @CreateDateColumn({ name: 'created_at' })
  public createdAt: Date
}

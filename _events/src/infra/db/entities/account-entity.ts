import { Column, CreateDateColumn, Entity, PrimaryColumn } from 'typeorm'

@Entity({ name: 'test_table' })
export class AccountEntity {
  @PrimaryColumn()
  public id: string

  @Column()
  public referencedId: string

  @CreateDateColumn({ name: 'created_at', type: 'timestamptz' })
  public createdAt: Date
}

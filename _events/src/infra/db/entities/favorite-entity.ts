import { Column, CreateDateColumn, Entity, PrimaryColumn } from 'typeorm'

@Entity({ name: 'favorites' })
export class FavoriteEntity {
  @PrimaryColumn({ type: 'uuid' })
  public id: string

  @Column({ name: 'account_id', type: 'uuid' })
  public accountId: string

  @Column({ name: 'event_id', type: 'uuid' })
  public eventId: string
}

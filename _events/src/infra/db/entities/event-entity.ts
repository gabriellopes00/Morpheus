import {
  EventAgeGroup,
  EventData,
  EventStatus,
  EventVisibility
} from '@/modules/events/domain/event'
import { Column, CreateDateColumn, Entity, PrimaryColumn } from 'typeorm'

@Entity({ name: 'events' })
export class EventEntity implements EventData {
  @PrimaryColumn({ type: 'uuid' })
  public id: string

  @Column({ type: 'varchar' })
  public name: string

  @Column({ type: 'varchar' })
  public description: string

  @Column({ type: 'varchar', name: 'cover_url' })
  public coverUrl: string

  @Column({ type: 'uuid', name: 'organizer_account_id' })
  public organizerAccountId: string

  @Column({ type: 'int', name: 'age_group', enum: [0, 10, 12, 14, 16, 18] })
  public ageGroup: EventAgeGroup

  @Column({
    type: 'varchar',
    name: 'status',
    enum: ['available', 'finished', 'sold_out', 'canceled']
  })
  public status: EventStatus

  @Column({ type: 'timestamp', name: 'start_date_time' })
  public startDateTime: Date

  @Column({ type: 'timestamp', name: 'end_date_time' })
  public endDateTime: Date

  @Column({ type: 'uuid', name: 'category_id' })
  public categoryId: string

  @Column({ type: 'uuid', name: 'subject_id' })
  public subjectId: string

  @Column({ type: 'varchar', name: 'visibility', enum: ['private', 'public'] })
  public visibility: EventVisibility

  @CreateDateColumn({ name: 'created_at', type: 'timestamptz' })
  public createdAt: Date

  @CreateDateColumn({ name: 'updated_at', type: 'timestamptz' })
  public updatedAt: Date
}

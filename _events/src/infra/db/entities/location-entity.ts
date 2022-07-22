import { Location, LocationData } from '@/modules/events/domain/location'
import { Column, CreateDateColumn, Entity, PrimaryColumn, UpdateDateColumn } from 'typeorm'

@Entity({ name: 'locations' })
export class LocationEntity implements LocationData {
  @PrimaryColumn({ type: 'uuid' })
  public id: string

  @Column({ type: 'uuid', name: 'event_id' })
  public eventId: string

  @Column({ type: 'varchar' })
  public description: string

  @Column({ type: 'varchar' })
  public street: string

  @Column({ type: 'varchar' })
  public district: string

  @Column({ type: 'char', length: 2 })
  public state: string

  @Column({ type: 'varchar' })
  public city: string

  @Column({ type: 'int' })
  public number: number

  @Column({ type: 'float', precision: 7, scale: 2 })
  latitude: number

  @Column({ type: 'float', precision: 7, scale: 2 })
  longitude: number

  @Column({ type: 'varchar', name: 'postal_code' })
  public postalCode: string

  @CreateDateColumn({ name: 'created_at', type: 'timestamptz' })
  public createdAt: Date

  @UpdateDateColumn({ name: 'updated_at', type: 'timestamptz' })
  public updatedAt: Date

  public map(): Location {
    return Object.assign(new Location(this, this.id), this)
  }
}

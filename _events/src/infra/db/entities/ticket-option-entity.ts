import { TicketOption, TicketOptionData } from '@/modules/events/domain/ticket-option'
import { Column, Entity, PrimaryColumn } from 'typeorm'

@Entity({ name: 'ticket_options' })
export class TicketOptionEntity implements TicketOptionData {
  @PrimaryColumn({ type: 'uuid' })
  public id: string

  @Column({ type: 'varchar' })
  title: string

  @Column({ type: 'varchar' })
  description: string

  @Column({ type: 'int', name: 'maximum_buys_quantity' })
  minimumBuysQuantity: number

  @Column({ type: 'int', name: 'minimum_buys_quantity' })
  maximumBuysQuantity: number

  @Column({ type: 'float', precision: 7, scale: 2 })
  price: number

  @Column({ type: 'int', name: 'total_quantity' })
  totalQuantity: number

  @Column({ type: 'int', name: 'remaining_quantity' })
  remainingQuantity: number

  @Column({ name: 'event_id', type: 'uuid' })
  public eventId: string

  public map(): TicketOption {
    return new TicketOption(
      {
        eventId: this.eventId,
        description: this.description,
        maximumBuysQuantity: this.maximumBuysQuantity,
        minimumBuysQuantity: this.minimumBuysQuantity,
        price: this.price,
        remainingQuantity: this.remainingQuantity,
        totalQuantity: this.totalQuantity,
        title: this.title
      },
      this.id
    )
  }
}

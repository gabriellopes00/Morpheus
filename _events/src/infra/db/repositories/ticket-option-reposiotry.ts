import { TicketOption } from '@/modules/events/domain/ticket-option'
import { SaveRepository } from '@/shared/repositories'
import { DataSource, Repository } from 'typeorm'
import { TicketOptionEntity } from '../entities/ticket-option-entity'

export class PgTicketOptionsRepository implements SaveRepository<TicketOption> {
  private readonly repository: Repository<TicketOptionEntity>

  constructor(private readonly dataSource: DataSource) {
    this.repository = this.dataSource.getRepository(TicketOptionEntity)
  }

  public async save(t: TicketOption): Promise<void> {
    await this.repository.save(
      this.repository.create({
        id: t.id,
        title: t.title,
        description: t.description,
        price: t.price,
        eventId: t.eventId,
        maximumBuysQuantity: t.maximumBuysQuantity,
        minimumBuysQuantity: t.minimumBuysQuantity,
        remainingQuantity: t.remainingQuantity,
        totalQuantity: t.totalQuantity
      })
    )
  }

  public async saveAll(tickets: TicketOption[]): Promise<void> {
    const entities: TicketOptionEntity[] = tickets.map(t =>
      this.repository.create({
        id: t.id,
        title: t.title,
        description: t.description,
        price: t.price,
        eventId: t.eventId,
        maximumBuysQuantity: t.maximumBuysQuantity,
        minimumBuysQuantity: t.minimumBuysQuantity,
        remainingQuantity: t.remainingQuantity,
        totalQuantity: t.totalQuantity
      })
    )
    await this.repository.save(entities)
  }
}

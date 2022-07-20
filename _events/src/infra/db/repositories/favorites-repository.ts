import { Event } from '@/modules/events/domain/event'
import { DataSource, In, Repository } from 'typeorm'
import { v4 as uuid } from 'uuid'
import { EventEntity } from '../entities/event-entity'
import { FavoriteEntity } from '../entities/favorite-entity'

export class PgFavoritesRepository {
  private readonly repository: Repository<FavoriteEntity>

  constructor(private readonly dataSource: DataSource) {
    this.repository = this.dataSource.getRepository(FavoriteEntity)
  }

  public async findByAccount(accountId: string): Promise<Event[]> {
    const relations = await this.repository.findBy({ accountId })
    if (relations.length === 0) return []

    const eventRepository = this.dataSource.getRepository(EventEntity)
    const entities = await eventRepository.findBy({ id: In(relations.map(r => r.eventId)) })
    return entities.map<Event>(entity => new Event({ ...entity }, entity.id))
  }

  public async count(eventId: string): Promise<number> {
    const total = await this.repository.countBy({ eventId })
    return total
  }

  public async delete(accountId: string, eventId: string): Promise<void> {
    await this.repository.delete({ accountId, eventId })
  }

  public async save(accountId: string, eventId: string): Promise<void> {
    await this.repository.save(this.repository.create({ id: uuid(), accountId, eventId }))
  }
}

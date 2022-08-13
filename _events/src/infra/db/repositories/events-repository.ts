import { Event } from '@/modules/events/domain/event'
import { FindRepository, SaveRepository } from '@/shared/repositories'
import { DataSource, In, Repository } from 'typeorm'
import { EventEntity } from '../entities/event-entity'

export class PgEventsRepository implements SaveRepository<Event>, FindRepository<Event> {
  private readonly repository: Repository<EventEntity>

  constructor(private readonly dataSource: DataSource) {
    this.repository = this.dataSource.getRepository(EventEntity)
  }

  public async save(event: Event): Promise<void> {
    await this.repository.save(
      this.repository.create({
        id: event.id,
        ageGroup: event.agrGroup as 0,
        name: event.name,
        description: event.description,
        startDateTime: event.startDateTime,
        endDateTime: event.endDateTime,
        status: event.status as 'available' | 'sold_out' | 'canceled',
        categoryId: event.categoryId,
        visibility: event.visibility as 'private' | 'public',
        coverUrl: event.coverUrl,
        organizerAccountId: event.organizerAccountId,
        createdAt: event.createdAt,
        updatedAt: new Date()
      })
    )
  }

  public async findBy?(key: keyof Event, value: any): Promise<Event> {
    const entity = await this.repository.findOneBy({ [key]: value })
    if (!entity) return null

    return entity.map()
  }

  public async exists?(params: { id: string }): Promise<boolean> {
    return !!(await this.repository.findOneBy({ id: params.id }))
  }

  public async findAll?(): Promise<Event[]> {
    const entities = await this.repository.find()
    return entities.map(e => e.map())
  }

  public async findAllBy(key: keyof Event, value: any): Promise<Event[]> {
    const entities = await this.repository.find({ where: { categoryId: value } })
    return entities.map(e => e.map())
  }

  public async findManyById(ids: string[]): Promise<Event[]> {
    const entities = await this.repository.find({ where: { id: In(ids) } })
    return entities.map(e => e.map())
  }
}

import { Location } from '@/modules/events/domain/location'
import { SaveRepository } from '@/shared/repositories'
import { DataSource, Repository } from 'typeorm'
import { LocationEntity } from '../entities/location-entity'

export class PgLocationsRepository implements SaveRepository<Location> {
  private readonly repository: Repository<LocationEntity>

  constructor(private readonly dataSource: DataSource) {
    this.repository = this.dataSource.getRepository(LocationEntity)
  }

  public async save(location: Location): Promise<void> {
    await this.repository.save(
      this.repository.create({
        id: location.id,
        street: location.street,
        city: location.city,
        state: location.state,
        description: location.description,
        createdAt: location.createdAt,
        updatedAt: new Date(),
        eventId: location.eventId,
        district: location.district,
        postalCode: location.postalCode.replace('-', ''),
        number: location.number,
        latitude: location.latitude,
        longitude: location.longitude
      })
    )
  }

  public async saveAll(ticket: Location[]): Promise<void> {
    const entities: LocationEntity[] = ticket.map(t => this.repository.create({ ...t }))
    await this.repository.save(entities)
  }
}

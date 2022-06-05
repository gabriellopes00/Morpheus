import { UUIDGenerator } from '@/core/infra/uuid-generator'
import { Location, LocationData } from '../../domain/location'
import { SaveRepository } from '../../repositories/generic'

export interface CreateLocationCredentials extends LocationData {}

export class CreateLocationUseCase {
  constructor(
    private readonly repository: SaveRepository<Location>,
    private readonly uuidGenerator: UUIDGenerator
  ) {}

  public async execute(params: CreateLocationCredentials): Promise<Location | Error> {
    const id = this.uuidGenerator.generate()
    const location = Location.create(params, id)
    if (location instanceof Error) {
      return location
    }

    await this.repository.save(location)
    return location
  }
}

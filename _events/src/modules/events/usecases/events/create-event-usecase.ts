import { UUIDGenerator } from '@/core/infra/uuid-generator'
import { Event, EventData, EventStatus } from '../../domain/event'
import { SaveEventRepository } from '../../repositories/events'

export interface CreateEventCredentials extends Omit<EventData, 'status'> {}

export class CreateEventUseCase {
  constructor(
    private readonly repository: SaveEventRepository,
    private readonly uuidGenerator: UUIDGenerator
  ) {}

  public async execute(params: CreateEventCredentials): Promise<Event | Error> {
    const defaultStatus: EventStatus = 'available'
    const id = this.uuidGenerator.generate()
    const event = Event.create({ ...params, status: defaultStatus }, id)
    if (event instanceof Error) {
      return event
    }

    await this.repository.save(event)
    return event
  }
}

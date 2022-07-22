import { UUIDGenerator } from '@/core/infra/uuid-generator'
import { Event, EventData, EventStatus } from '../../domain/event'
import { SaveRepository } from '@/shared/repositories'

export interface CreateEventCredentials
  extends Modify<
    Omit<EventData, 'status'>,
    {
      startDateTime: string
      endDateTime: string
    }
  > {}

export class CreateEventUseCase {
  constructor(
    private readonly repository: SaveRepository<Event>,
    private readonly uuidGenerator: UUIDGenerator
  ) {}

  public async execute(params: CreateEventCredentials): Promise<Event | Error> {
    const defaultStatus: EventStatus = 'available'
    const id = this.uuidGenerator.generate()
    const event = Event.create(
      {
        ...params,
        status: defaultStatus,
        startDateTime: new Date(params.startDateTime),
        endDateTime: new Date(params.endDateTime)
      },
      id
    )
    if (event instanceof Error) {
      return event
    }

    await this.repository.save(event)
    return event
  }
}

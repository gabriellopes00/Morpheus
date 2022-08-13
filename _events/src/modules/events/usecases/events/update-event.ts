import { FindRepository, SaveRepository } from '@/shared/repositories'
import { Event, EventAgeGroup } from '../../domain/event'

export interface UpdateEventParams {
  id: string
  name?: string
  description?: string
  coverUrl?: string
  ageGroup?: number
}

export class UpdateEvent {
  constructor(private readonly repository: SaveRepository<Event> & FindRepository<Event>) {}

  public async execute(params: UpdateEventParams): Promise<Event> {
    const event = await this.repository.findBy('id', params.id)
    if (params.name) event.name = params.name
    if (params.description) event.description = params.description
    if (params.coverUrl) event.coverUrl = params.coverUrl
    if (params.ageGroup) event.ageGroup = params.ageGroup as EventAgeGroup
    await this.repository.save(event)
    return event
  }
}

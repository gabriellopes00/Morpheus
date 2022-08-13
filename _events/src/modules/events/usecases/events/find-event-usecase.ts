import { FindRepository } from '@/shared/repositories'
import { Event } from '../../domain/event'
import { TicketOption } from '../../domain/ticket-option'

export class FindEvent {
  constructor(
    private readonly repository: FindRepository<Event>,
    private readonly ticketsRepository: FindRepository<TicketOption>
  ) {}

  public async execute(params: {
    id: string
  }): Promise<Modify<Event, { tickets: TicketOption[] }>> {
    const event = await this.repository.findBy('id', params.id)
    const tickets = await this.ticketsRepository.findAllBy('eventId', params.id)

    return Object.assign(event, { tickets })
  }
}

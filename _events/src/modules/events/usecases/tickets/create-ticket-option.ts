import { UUIDGenerator } from '@/core/infra/uuid-generator'
import { Event } from '../../domain/event'
import { TicketOption, TicketOptionData } from '../../domain/ticket-option'
import { FindRepository, SaveRepository } from '@/shared/repositories'

export interface CreateTicketOptionCredentials extends TicketOptionData {}

export class CreateTicketOption {
  constructor(
    private readonly ticketRepository: SaveRepository<TicketOption>,
    private readonly eventRepository: FindRepository<Event>,
    private readonly uuidGenerator: UUIDGenerator
  ) {}

  public async execute(params: CreateTicketOptionCredentials[]): Promise<TicketOption[] | Error> {
    if (!params.every((t, _, arr) => t.eventId === arr[0].eventId)) {
      return new Error('All ticket options must have the same event id')
    }

    const eventExists = await this.eventRepository.exists({ id: params[0].eventId })
    if (!eventExists) return new Error('Event not found')

    const options = params.map(option => {
      const id = this.uuidGenerator.generate()
      const ticketOption = TicketOption.create(option, id)
      return ticketOption
    })

    const error = options.find(t => t instanceof Error)
    if (error) return error as Error

    await this.ticketRepository.saveAll(options as TicketOption[])
    return options as TicketOption[]
  }
}

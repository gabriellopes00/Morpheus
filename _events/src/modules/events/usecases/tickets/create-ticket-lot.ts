import { UUIDGenerator } from '@/core/infra/uuid-generator'
import { TicketOption } from '../../domain/ticket-option'
import { TicketOptionLot, TicketOptionLotData } from '../../domain/ticket-option-lot'
import { FindRepository, SaveRepository } from '../../repositories/generic'

export interface CreateTicketLotCredentials {
  lots: TicketOptionLotData[]
}

export class CreateTicketLot {
  constructor(
    private readonly lotRepository: SaveRepository<TicketOptionLot>,
    private readonly optionRepository: FindRepository<TicketOption>,
    private readonly uuidGenerator: UUIDGenerator
  ) {}

  public async execute(params: CreateTicketLotCredentials): Promise<TicketOptionLot[] | Error> {
    if (!params.lots.every((t, _, arr) => t.ticketOptionId === arr[0].ticketOptionId)) {
      return new Error('All ticket lots must have the same ticket option id')
    }

    const optionExists = await this.optionRepository.exists({ id: params.lots[0].ticketOptionId })
    if (!optionExists) return new Error('Event not found')

    const lots = params.lots.map(lot => {
      const id = this.uuidGenerator.generate()
      const ticketLot = TicketOptionLot.create(lot, id)
      return ticketLot
    })

    const error = lots.find(t => t instanceof Error)
    if (error) return error as Error

    await this.lotRepository.saveAll(lots as TicketOptionLot[])
    return lots as TicketOptionLot[]
  }
}

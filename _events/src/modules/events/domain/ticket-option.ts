import { Entity } from '@/shared/entity'
import { TicketOptionLot } from './ticket-option-lot'

export interface TicketOptionData {
  eventId: string
  title: string
  description: string
  salesStartDateTime: Date
  salesEndDateTime: Date
  minimumBuysQuantity: number
  maximumBuysQuantity: number
  lots: TicketOptionLot[]
}

export class TicketOption extends Entity<TicketOptionData> {}

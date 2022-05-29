import { Entity } from '@/shared/entity'

export interface TicketOptionLotData {
  number: number
  ticketOptionId: string
  price: number
  quantity: number
}

export class TicketOptionLot extends Entity<TicketOptionLotData> {}

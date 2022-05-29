import { Entity } from '@/shared/entity'

export interface TicketOptionLotData {
  number: number
  ticketOptionId: string
  price: number
  total_quantity: number
  remaining_quantity: number
}

export class TicketOptionLot extends Entity<TicketOptionLotData> {}

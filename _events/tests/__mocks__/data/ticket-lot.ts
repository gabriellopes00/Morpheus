import { CreateTicketLotCredentials } from '@/modules/events/usecases/tickets/create-ticket-lot'

export const ticketLot: CreateTicketLotCredentials = {
  number: 1,
  price: 149.99,
  ticketOptionId: crypto.randomUUID(),
  remainingQuantity: 200,
  totalQuantity: 200
}

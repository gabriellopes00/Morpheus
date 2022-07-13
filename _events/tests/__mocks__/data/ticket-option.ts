import { CreateTicketOptionCredentials } from '@/modules/events/usecases/tickets/create-ticket-option'

export const ticketOption: CreateTicketOptionCredentials = {
  title: 'V.I.P',
  description: 'V.I.P tickets provider a great experience to the user',
  minimumBuysQuantity: 1,
  maximumBuysQuantity: 4,
  eventId: crypto.randomUUID()
}

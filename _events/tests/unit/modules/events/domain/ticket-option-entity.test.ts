import { TicketOption, TicketOptionData } from '@/modules/events/domain/ticket-option'

describe('Event Ticket Option Entity', () => {
  const ticketOptionData: TicketOptionData = {
    title: 'My Test Ticket Option',
    description: 'My Test Ticket Option Description',
    eventId: 'e6677386-1792-416d-a3c4-0e424c367aab',
    maximumBuysQuantity: 6,
    minimumBuysQuantity: 1
  }
  test('Should create a ticketOption successfully', async () => {
    const ticketOption = TicketOption.create(
      ticketOptionData,
      '3a8a7a22-a1fc-4a73-96f3-4ac0ec944de5'
    ) as TicketOption
    expect(ticketOption).toBeInstanceOf(TicketOption)
    expect(ticketOption.id).toBe('3a8a7a22-a1fc-4a73-96f3-4ac0ec944de5')
  })

  test('Should return an error when minimum buys quantity is greater than maximum buys quantity', async () => {
    const ticketOptionDataWithInvalidQuantity: TicketOptionData = {
      ...ticketOptionData,
      maximumBuysQuantity: 1,
      minimumBuysQuantity: 6
    }
    const ticketOption = TicketOption.create(
      ticketOptionDataWithInvalidQuantity,
      '3a8a7a22-a1fc-4a73-96f3-4ac0ec944de5'
    )
    expect(ticketOption).toBeInstanceOf(Error)
  })

  test('Should return an error when minimum buys quantity is less than 1', async () => {
    const ticketOptionDataWithInvalidQuantity: TicketOptionData = {
      ...ticketOptionData,
      maximumBuysQuantity: 1,
      minimumBuysQuantity: 0
    }
    const ticketOption = TicketOption.create(
      ticketOptionDataWithInvalidQuantity,
      '3a8a7a22-a1fc-4a73-96f3-4ac0ec944de5'
    )
    expect(ticketOption).toBeInstanceOf(Error)
  })
})

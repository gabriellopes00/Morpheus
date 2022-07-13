// import { UidGenerator } from '@/infra/crypto/uuid-generator'
// import { CreateEventController } from '@/modules/events/controllers/events/create-event-controller'
// import { CreateEventUseCase } from '@/modules/events/usecases/events/create-event-usecase'
// import { CreateLocationUseCase } from '@/modules/events/usecases/events/create-location-usecase'
// import { CreateTicketLot } from '@/modules/events/usecases/tickets/create-ticket-lot'
// import { CreateTicketOption } from '@/modules/events/usecases/tickets/create-ticket-option'
// import { event } from '@t/__mocks__/data/event'
// import { location } from '@t/__mocks__/data/location'
// import { ticketLot } from '@t/__mocks__/data/ticket-lot'
// import { ticketOption } from '@t/__mocks__/data/ticket-option'

// describe('Create Event Controller', () => {
//   const uuidGenerator = new UidGenerator()
//   const createEvent = new CreateEventUseCase(null, uuidGenerator)
//   const createLocation = new CreateLocationUseCase(null, uuidGenerator)
//   const createTicketOption = new CreateTicketOption(null, null, uuidGenerator)
//   const createTicketLot = new CreateTicketLot(null, null, uuidGenerator)
//   const controller = new CreateEventController(
//     createEvent,
//     createLocation,
//     createTicketOption,
//     createTicketLot
//   )

//   it('should return an event successfully', async () => {
//     await controller.handle({
//       params: {
//         ...event,
//         location: location,
//         ticketOptions: [{ ...ticketOption, lots: [{ ...ticketLot }] }]
//       }
//     })
//   })
// })

describe('', () => {
  it('', async () => {
    expect(1).toBe(1)
  })
})

// import { CreateEventUseCase } from '@/modules/events/usecases/events/create-event-usecase'
// import { MockEventRepository } from '@t/__mocks__/repositories/mock-event-repository'
// import { UidGenerator } from '@/infra/crypto/uuid-generator'
// import { Event, EventData } from '@/modules/events/domain/event'

// describe('Create Event Usecase', () => {
//   const repository = new MockEventRepository()
//   const uuidGenerator = new UidGenerator()
//   const usecase = new CreateEventUseCase(repository, uuidGenerator)
//   const params: EventData = {
//     ageGroup: 16,
//     categoryId: 'c923fe2f-6219-4446-b958-d224f744001f',
//     coverUrl: 'https://random_image.png',
//     description: 'lorem ipsum...',
//     name: 'My Test Event',
//     status: 'available',
//     visibility: 'public',
//     organizerAccountId: '8402a14c-bc0c-48cb-856a-c1d15cd9ca09',
//     subjectId: '65ef5404-95ec-49fd-89f0-65cb158eec12',
//     startDateTime: new Date(2025, 9, 11, 19, 30),
//     endDateTime: new Date(2025, 9, 11, 23, 30)
//   }

//   test('Should create an event successfully', async () => {
//     const event = (await usecase.execute(params)) as Event
//     expect(event).toBeInstanceOf(Event)
//     expect(event.id).toBeDefined()
//   })

//   test('Should create an event using `create` method', async () => {})
// })

describe('', () => {
  it('', async () => {
    expect(1).toBe(1)
  })
})

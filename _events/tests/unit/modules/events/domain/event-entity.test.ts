import { Event, EventData } from '@/modules/events/domain/event'

/* eslint-disable */
describe('Event', () => {
  const eventData: EventData = {
    ageGroup: 16,
    categoryId: 'c923fe2f-6219-4446-b958-d224f744001f',
    coverUrl: 'https://random_image.png',
    description: 'lorem ipsum...',
    name: 'My Test Event',
    status: 'available',
    visibility: 'public',
    organizerAccountId: '8402a14c-bc0c-48cb-856a-c1d15cd9ca9',
    subjectId: '65ef5404-95ec-49fd-89f0-65cb158eec12',
    startDateTime: new Date(2025, 9, 11, 19, 30),
    endDateTime: new Date(2025, 9, 11, 23, 30)
  }

  it('Should create an event successfully', async () => {
    const event = Event.create(eventData, 'e6677386-1792-416d-a3c4-0e424c367aab')
    expect(event).toBeInstanceOf(Event)
  })

  test('Should not create an event if start date is in the past', async () => {
    const params: EventData = {
      ...eventData,
      startDateTime: new Date(2020, 9, 10, 19, 30),
      endDateTime: new Date(2020, 9, 10, 23, 30)
    }
    const error = Event.create(params, 'aebd7c0b-3a16-4d5c-b4de-68db8c98dc58')
    expect(error).toEqual(new Error('Event cannot start in the past'))
  })

  test('Should not create an event if start date is after end date', async () => {
    const params: EventData = {
      ...eventData,
      startDateTime: new Date(2025, 9, 10, 23, 30),
      endDateTime: new Date(2025, 9, 10, 19, 30)
    }
    const error = Event.create(params, 'aebd7c0b-3a16-4d5c-b4de-68db8c98dc58')
    expect(error).toEqual(new Error('Event cannot start after it ends'))
  })

  test('Should not create an event if duration is less than one hour', async () => {
    const params: EventData = {
      ...eventData,
      startDateTime: new Date(2025, 9, 10, 19, 30),
      endDateTime: new Date(2025, 9, 10, 19, 31)
    }
    const error = Event.create(params, 'aebd7c0b-3a16-4d5c-b4de-68db8c98dc58')
    expect(error).toEqual(new Error('Event must have at least one hour of duration'))
  })

  test('Should not create an event if status is not available', async () => {
    const params: EventData = {
      ...eventData,
      status: 'sold_out'
    }
    const error = Event.create(params, 'aebd7c0b-3a16-4d5c-b4de-68db8c98dc58')
    expect(error).toEqual(new Error("New events' status must be 'available'"))
  })

  test('Should not create an event if age group is invalid', async () => {
    const params: EventData = {
      ...eventData,
      ageGroup: 890 as any // invalid age group
    }
    const error = Event.create(params, 'aebd7c0b-3a16-4d5c-b4de-68db8c98dc58')
    expect(error).toEqual(new Error('Invalid age group'))
  })
})

/* eslint-enable */

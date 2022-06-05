import { Location, LocationData } from '@/modules/events/domain/location'

/* eslint-disable */
describe('Location', () => {
  const locationData: LocationData = {
    city: 'Jundiaí',
    district: 'Anhangabaú',
    description: 'My event location',
    latitude: 40.73061,
    longitude: -73.935242,
    number: 123,
    postalCode: '12345-678',
    state: 'SP',
    street: 'Av. 9 de Julho',
    eventId: 'b134f25b-892a-428c-bcc9-8adbb47b5d22'
  }

  it('Should create an location successfully', async () => {
    const location = Location.create(locationData, 'e6677386-1792-416d-a3c4-0e424c367aab')
    expect(location).toBeInstanceOf(Location)
  })

  test('Should not create an location with an invalid postal code', async () => {
    const params: LocationData = {
      ...locationData,
      postalCode: '890890'
    }
    const error = Location.create(params, 'aebd7c0b-3a16-4d5c-b4de-68db8c98dc58')
    expect(error).toEqual(new Error('Invalid postal code'))
  })

  test('Should not create an location with an invalid state abbr', async () => {
    const params: LocationData = {
      ...locationData,
      state: 'São Paulo'
    }
    const error = Location.create(params, 'aebd7c0b-3a16-4d5c-b4de-68db8c98dc58')
    expect(error).toEqual(new Error('Invalid state abbr'))
  })
})

/* eslint-enable */

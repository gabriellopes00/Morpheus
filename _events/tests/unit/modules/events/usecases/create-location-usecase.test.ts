import { UidGenerator } from '@/infra/crypto/uuid-generator'
import { Location, LocationData } from '@/modules/events/domain/location'
import { CreateLocationUseCase } from '@/modules/events/usecases/events/create-location-usecase'
import { MockLocationRepository } from '@t/__mocks__/repositories/mock-location-repository'

describe('Create Location Usecase', () => {
  const repository = new MockLocationRepository()
  const uuidGenerator = new UidGenerator()
  const usecase = new CreateLocationUseCase(repository, uuidGenerator)
  const params: LocationData = {
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

  test('Should create an event successfully', async () => {
    const event = (await usecase.execute(params)) as Location
    expect(event).toBeInstanceOf(Location)
    expect(event.id).toBeDefined()
  })

  test('Should save a location in repository using `save` method', async () => {
    const saveSpy = jest.spyOn(repository, 'save')
    await usecase.execute(params)
    expect(saveSpy).toHaveBeenCalledTimes(1)
    expect(saveSpy).toHaveBeenCalledWith(expect.any(Location))
  })

  test('Should create a location using `create` method', async () => {
    const createSpy = jest.spyOn(Location, 'create')
    await usecase.execute(params)
    expect(createSpy).toHaveReturnedTimes(1)
    expect(createSpy).toHaveBeenCalledWith(
      params,
      expect.stringMatching(/^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$/)
    )
  })
})

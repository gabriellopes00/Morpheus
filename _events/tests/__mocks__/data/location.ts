import { CreateLocationCredentials } from '@/modules/events/usecases/events/create-location-usecase'

export const location: CreateLocationCredentials = {
  eventId: crypto.randomUUID(),
  street: 'Beiramar ipanema',
  district: 'Ipanema district',
  state: 'RJ',
  city: 'Rio de Janeiro',
  number: 645,
  postalCode: '13216-570',
  description: '',
  latitude: -22.986851946094028,
  longitude: -43.20215535253323
}

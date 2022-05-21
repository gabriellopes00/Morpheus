import { Entity } from '@/shared/entity'

export interface LocationData {
  eventId: string
  street: string
  district: string
  state: string
  city: string
  number: number
  postalCode: string
  description: string
  latitude: number
  longitude: number
}

export class Location extends Entity<LocationData> {}

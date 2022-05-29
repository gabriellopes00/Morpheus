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

export class Location extends Entity<LocationData> {
  constructor(data: LocationData, id: string) {
    super(data, id)
  }

  static create(data: LocationData): Location | Error {
    return new Location(data, 'uuid')
  }

  public get eventId(): string {
    return this.data.eventId
  }

  public get street(): string {
    return this.data.street
  }

  public get district(): string {
    return this.data.district
  }

  public get state(): string {
    return this.data.state
  }

  public get city(): string {
    return this.data.city
  }

  public get number(): number {
    return this.data.number
  }

  public get postalCode(): string {
    return this.data.postalCode
  }

  public get description(): string {
    return this.data.description
  }

  public get latitude(): number {
    return this.data.latitude
  }

  public get longitude(): number {
    return this.data.longitude
  }

  public get address(): string {
    return `${this.street}, ${this.number}, ${this.city}`
  }

  public get location(): string {
    return `${this.latitude},${this.longitude}`
  }
}

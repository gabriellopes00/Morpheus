import { FindRepository } from '@/shared/repositories'
import axios from 'axios'
import { Event } from '../../domain/event'
import { Location } from '../../domain/location'

export interface FindNearbyEventsParams {
  latitude: number
  longitude: number
}

interface GeoCodeResponse {
  latitude: number
  longitude: number
  continent: string
  lookupSource: string
  continentCode: string
  city: string
  countryName: string
  postcode: string
  countryCode: string
  principalSubdivision: string
  principalSubdivisionCode: string
}

export class FindNearbyEvents {
  constructor(
    private readonly repository: FindRepository<Event>,
    private readonly locationRepository: FindRepository<Location>
  ) {}

  public async execute(params: FindNearbyEventsParams): Promise<Event[]> {
    const { data } = await axios.get<GeoCodeResponse>(
      `https://api.bigdatacloud.net/data/reverse-geocode-client?latitude=${params.latitude}&longitude=${params.longitude}&localityLanguage=en`
    )

    const locations = await this.locationRepository.findAllBy('city', data.city)
    const events = await this.repository.findManyById(locations.map(l => l.eventId))

    return events
  }
}

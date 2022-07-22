import { Location } from '@/modules/events/domain/location'
import { SaveRepository } from '@/modules/events/repositories/generic'

export class MockLocationRepository implements SaveRepository {
  private readonly _events: Location[] = []

  get rows() {
    return this._events
  }

  public truncate() {
    this._events.length = 0
  }

  public async save<T = Location>(data: T): Promise<void> {
    const location = data as unknown as Location
    if (this._events.some(a => a.id === location.id)) {
      throw new Error('id must be unique')
    }

    this._events.push(location)
  }

  public async saveAll<I = any>(events: never[] | I[]): Promise<void> {
    throw new Error('Method not implemented.')
  }
}

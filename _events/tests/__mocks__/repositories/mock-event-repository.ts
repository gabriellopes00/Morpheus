import { Event } from '@/modules/events/domain/event'
import { SaveEventRepository } from '@/modules/events/repositories/events'

export class MockEventRepository implements SaveEventRepository {
  private readonly _events: Event[] = []

  get rows() {
    return this._events
  }

  public truncate() {
    this._events.length = 0
  }

  public async save(data: Event): Promise<void> {
    if (this._events.some(a => a.id === data.id)) {
      throw new Error('id must be unique')
    }

    this._events.push(data)
  }
}

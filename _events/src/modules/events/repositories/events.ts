import { Event } from '../domain/event'

export interface SaveEventRepository {
  save(event: Event): Promise<void>
}

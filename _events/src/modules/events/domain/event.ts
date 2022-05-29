import { Entity } from '@/shared/entity'

export type EventStatus = 'available' | 'finished' | 'sold_out' | 'canceled'

export type EventVisibility = 'public' | 'private' | 'invited_only'

export interface EventData {
  name: string
  description: string
  coverUrl: string
  organizerAccountId: string
  ageGroup: number
  status: EventStatus
  locationId: string
  startDateTime: Date
  endDateTime: Date
  categoryId: string
  subjectId: string
  visibility: EventVisibility
}

export class Event extends Entity<EventData> {
  constructor(data: EventData, id: string) {
    super(data, id)
  }

  public get name(): string {
    return this.data.name
  }

  public get description(): string {
    return this.data.description
  }

  public get coverUrl(): string {
    return this.data.coverUrl
  }

  public get organizerAccountId(): string {
    return this.data.organizerAccountId
  }

  public get agrGroup(): number {
    return this.data.ageGroup
  }

  public get status(): string {
    return this.data.status
  }

  public get locationId(): string {
    return this.data.locationId
  }

  public get startDateTime(): Date {
    return this.data.startDateTime
  }

  public get endDateTime(): Date {
    return this.data.endDateTime
  }

  public get categoryId(): string {
    return this.data.categoryId
  }

  public get subjectId(): string {
    return this.data.subjectId
  }

  public get visibility(): string {
    return this.data.visibility
  }
}

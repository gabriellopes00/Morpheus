import { Entity } from '@/shared/entity'

export type EventStatus = 'available' | 'finished' | 'sold_out' | 'canceled'
export type EventAgeGroup = 0 | 10 | 12 | 14 | 16 | 18
export type EventVisibility = 'public' | 'private' | 'invited_only'

export interface EventData {
  name: string
  description: string
  coverUrl: string
  organizerAccountId: string
  ageGroup: EventAgeGroup
  status: EventStatus
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

  static create(data: EventData, id: string): Event | Error {
    const timeUntilStart = data.startDateTime.getTime() - new Date().getTime()
    if (timeUntilStart < 0) {
      return new Error('Event cannot start in the past')
    }

    if (data.startDateTime.getTime() > data.endDateTime.getTime()) {
      return new Error('Event cannot start after it ends')
    }

    const duration = data.endDateTime.getTime() - data.startDateTime.getTime()
    if (duration < 3600000) {
      return new Error('Event must have at least one hour of duration')
    }

    if (data.status !== 'available') {
      return new Error("New events' status must be 'available'")
    }

    if (
      data.ageGroup !== 0 &&
      data.ageGroup !== 10 &&
      data.ageGroup !== 12 &&
      data.ageGroup !== 14 &&
      data.ageGroup !== 16 &&
      data.ageGroup !== 18
    ) {
      return new Error('Invalid age group')
    }

    return new Event(data, id)
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

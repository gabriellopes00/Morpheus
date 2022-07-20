import { Entity } from '@/shared/entity'

export interface TicketOptionData {
  eventId: string
  title: string
  description: string
  minimumBuysQuantity: number
  maximumBuysQuantity: number
  price: number
  totalQuantity: number
  remainingQuantity: number
}

export class TicketOption extends Entity<TicketOptionData> {
  constructor(data: TicketOptionData, id: string) {
    super(data, id)
  }

  static create(data: TicketOptionData, id: string): TicketOption | Error {
    if (data.minimumBuysQuantity > data.maximumBuysQuantity) {
      return new Error('Minimum buys quantity cannot be greater than maximum buys quantity')
    }

    if (data.minimumBuysQuantity < 1) {
      return new Error('Minimum buys quantity cannot be less than 1')
    }

    return new TicketOption(data, id)
  }

  public get eventId(): string {
    return this.data.eventId
  }

  public get title(): string {
    return this.data.title
  }

  public get description(): string {
    return this.data.description
  }

  public get minimumBuysQuantity(): number {
    return this.data.minimumBuysQuantity
  }

  public get maximumBuysQuantity(): number {
    return this.data.maximumBuysQuantity
  }

  public get price(): number {
    return this.data.price
  }

  public get totalQuantity(): number {
    return this.data.totalQuantity
  }

  public get remainingQuantity(): number {
    return this.data.remainingQuantity
  }
}

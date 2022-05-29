import { Entity } from '@/shared/entity'

export interface TicketOptionLotData {
  number: number
  ticketOptionId: string
  price: number
  totalQuantity: number
  remainingQuantity: number
}

export class TicketOptionLot extends Entity<TicketOptionLotData> {
  constructor(data: TicketOptionLotData, id: string) {
    super(data, id)
  }

  static create(data: TicketOptionLotData, id: string): TicketOptionLot | Error {
    if (data.number < 1) {
      return new Error('Ticket option lot number cannot be less than 1')
    }

    if (data.price < 0) {
      return new Error('Ticket option lot price cannot be less than 0')
    }

    if (data.totalQuantity < 1) {
      return new Error('Ticket option lot total quantity cannot be less than 1')
    }

    if (data.remainingQuantity < 0) {
      return new Error('Ticket option lot remaining quantity cannot be less than 0')
    }

    if (data.remainingQuantity > data.totalQuantity) {
      return new Error('Ticket option lot remaining quantity cannot be greater than total quantity')
    }

    return new TicketOptionLot(data, id)
  }

  public get number(): number {
    return this.data.number
  }

  public get ticketOptionId(): string {
    return this.data.ticketOptionId
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

  public get isSoldOut(): boolean {
    return this.remainingQuantity === 0
  }
}

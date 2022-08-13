import { Request, Response } from 'express'
import {
  CreateEventCredentials,
  CreateEventUseCase
} from '../../usecases/events/create-event-usecase'
import {
  CreateLocationCredentials,
  CreateLocationUseCase
} from '../../usecases/events/create-location-usecase'
import { FindEvent } from '../../usecases/events/find-event-usecase'
import { FindEventsUseCase } from '../../usecases/events/find-events-usecase'
import { FindNearbyEvents } from '../../usecases/events/find-nearby-events'
import { UpdateEvent } from '../../usecases/events/update-event'
import {
  CreateTicketOption,
  CreateTicketOptionCredentials
} from '../../usecases/tickets/create-ticket-option'

interface CreateParams extends CreateEventCredentials {
  location: CreateLocationCredentials
  ticketOptions: CreateTicketOptionCredentials[]
}

type EventRequest<T = any> = Modify<Request, { body: T }>

export class EventController {
  constructor(
    private readonly createEvent: CreateEventUseCase,
    private readonly createLocation: CreateLocationUseCase,
    private readonly createTicketOption: CreateTicketOption,
    private readonly findEvents: FindEventsUseCase,
    private readonly findEvent: FindEvent,
    private readonly updateEvent: UpdateEvent,
    private readonly findNearbyEvents: FindNearbyEvents
  ) {
    const methods = Object.getOwnPropertyNames(Object.getPrototypeOf(this))
    methods.filter(m => m !== 'constructor').forEach(m => (this[m] = this[m].bind(this)))
  }

  public async create(req: EventRequest<CreateParams>, res: Response): Promise<Response> {
    try {
      const data = req.body
      const organizerAccountId = String(req.headers.account_id)

      const result = await this.createEvent.execute({ ...data, organizerAccountId })
      if (result instanceof Error) return res.status(400).json({ error: result.message })

      const locationResult = await this.createLocation.execute({
        ...data.location,
        eventId: result.id
      })
      if (locationResult instanceof Error) {
        return res.status(400).json({ error: locationResult.message })
      }

      const options = data.ticketOptions.map(option => ({ ...option, eventId: result.id }))
      const optionResult = await this.createTicketOption.execute(options)
      if (optionResult instanceof Error) {
        return res.status(400).json({ error: optionResult.message })
      }

      return res.status(201).json({ event: result })
    } catch (error) {
      return res.status(500).json({ error: error.message })
    }
  }

  public async findAll(req: EventRequest, res: Response): Promise<Response> {
    try {
      const result = await this.findEvents.execute()
      return res.status(200).json({ result })
    } catch (error) {
      return res.status(500).json({ error: error.message })
    }
  }

  public async findById(req: EventRequest, res: Response): Promise<Response> {
    try {
      const result = await this.findEvent.execute({ id: req.params.id })
      return res.status(200).json({ result })
    } catch (error) {
      return res.status(500).json({ error: error.message })
    }
  }

  public async findNearby(req: Request, res: Response): Promise<Response> {
    try {
      const result = await this.findNearbyEvents.execute({
        latitude: Number(req.query.latitude),
        longitude: Number(req.query.longitude)
      })
      return res.status(200).json({ result })
    } catch (error) {
      return res.status(500).json({ error: error.message })
    }
  }

  public async update(req: EventRequest, res: Response): Promise<Response> {
    try {
      const result = await this.updateEvent.execute({ id: req.params.id, ...req.body })
      return res.status(200).json({ result })
    } catch (error) {
      return res.status(500).json({ error: error.message })
    }
  }
}

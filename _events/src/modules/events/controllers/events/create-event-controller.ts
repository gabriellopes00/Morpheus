import { Controller } from '@/core/presentation/controller'
import { HttpRequest, HttpResponse } from '@/core/presentation/http'
import { badRequest, created } from '@/presentation/http'
import { Event } from '../../domain/event'
import {
  CreateEventCredentials,
  CreateEventUseCase
} from '../../usecases/events/create-event-usecase'
import {
  CreateLocationCredentials,
  CreateLocationUseCase
} from '../../usecases/events/create-location-usecase'
import {
  CreateTicketOption,
  CreateTicketOptionCredentials
} from '../../usecases/tickets/create-ticket-option'

interface Params extends CreateEventCredentials {
  location: CreateLocationCredentials
  ticketOptions: CreateTicketOptionCredentials[]
}

export class CreateEventController implements Controller {
  constructor(
    private readonly createEvent: CreateEventUseCase,
    private readonly createLocation: CreateLocationUseCase,
    private readonly createTicketOption: CreateTicketOption
  ) {}

  public async handle(data: HttpRequest<Params>): Promise<HttpResponse<Event>> {
    const result = await this.createEvent.execute(data.params)
    if (result instanceof Error) return badRequest(result)

    const locationResult = await this.createLocation.execute(data.params.location)
    if (locationResult instanceof Error) return badRequest(locationResult)

    const optionResult = await this.createTicketOption.execute(data.params.ticketOptions)
    if (optionResult instanceof Error) return badRequest(optionResult)

    return created(result)
  }
}

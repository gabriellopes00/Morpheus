import { JwtEncrypter } from '@/infra/crypto/jwt-encrypter'
import { UidGenerator } from '@/infra/crypto/uuid-generator'
import { TypeORMDataSource } from '@/infra/db/data-source'
import { CategoryRepository } from '@/infra/db/repositories/categories-repository'
import { PgEventsRepository } from '@/infra/db/repositories/events-repository'
import { PgLocationsRepository } from '@/infra/db/repositories/location-repository'
import { PgTicketOptionsRepository } from '@/infra/db/repositories/ticket-option-reposiotry'
import { AuthValidation } from '@/modules/accounts/auth/validate-auth'
import { EventController } from '@/modules/events/controllers/events/create-event-controller'
import { CreateEventUseCase } from '@/modules/events/usecases/events/create-event-usecase'
import { CreateLocationUseCase } from '@/modules/events/usecases/events/create-location-usecase'
import { FindEvent } from '@/modules/events/usecases/events/find-event-usecase'
import { FindEventsUseCase } from '@/modules/events/usecases/events/find-events-usecase'
import { FindNearbyEvents } from '@/modules/events/usecases/events/find-nearby-events'
import { UpdateEvent } from '@/modules/events/usecases/events/update-event'
import { CreateTicketOption } from '@/modules/events/usecases/tickets/create-ticket-option'
import { Router } from 'express'
import { AuthMiddleware } from '../middlewares/auth-middleware'

const router = Router()

const createEvent = new CreateEventUseCase(
  new PgEventsRepository(TypeORMDataSource.getDataSource()),
  new UidGenerator()
)

const createLocation = new CreateLocationUseCase(
  new PgLocationsRepository(TypeORMDataSource.getDataSource()),
  new UidGenerator()
)

const createTicketOption = new CreateTicketOption(
  new PgTicketOptionsRepository(TypeORMDataSource.getDataSource()),
  new PgEventsRepository(TypeORMDataSource.getDataSource()),
  new UidGenerator()
)

const findEventsUsecase = new FindEventsUseCase(
  new PgEventsRepository(TypeORMDataSource.getDataSource()),
  new CategoryRepository(TypeORMDataSource.getDataSource())
)

const findEvent = new FindEvent(
  new PgEventsRepository(TypeORMDataSource.getDataSource()),
  new PgTicketOptionsRepository(TypeORMDataSource.getDataSource())
)

const updateEvent = new UpdateEvent(new PgEventsRepository(TypeORMDataSource.getDataSource()))

const findNearbyEvents = new FindNearbyEvents(
  new PgEventsRepository(TypeORMDataSource.getDataSource()),
  new PgLocationsRepository(TypeORMDataSource.getDataSource())
)

const controller = new EventController(
  createEvent,
  createLocation,
  createTicketOption,
  findEventsUsecase,
  findEvent,
  updateEvent,
  findNearbyEvents
)
const middleware = new AuthMiddleware(new AuthValidation(new JwtEncrypter()))

router.post('/events', middleware.handle, controller.create)
router.get('/events', middleware.handle, controller.findAll)
router.get('/events/nearby', middleware.handle, controller.findNearby)
router.get('/events/:id', middleware.handle, controller.findById)
router.put('/events/:id', middleware.handle, controller.update)

export { router as eventsRouter }

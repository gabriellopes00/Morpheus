import { JwtEncrypter } from '@/infra/crypto/jwt-encrypter'
import { UidGenerator } from '@/infra/crypto/uuid-generator'
import { TypeORMDataSource } from '@/infra/db/data-source'
import { PgEventsRepository } from '@/infra/db/repositories/events-repository'
import { PgLocationsRepository } from '@/infra/db/repositories/location-repository'
import { PgTicketOptionsRepository } from '@/infra/db/repositories/ticket-option-reposiotry'
import { AuthValidation } from '@/modules/accounts/auth/validate-auth'
import { EventController } from '@/modules/events/controllers/events/create-event-controller'
import { CreateEventUseCase } from '@/modules/events/usecases/events/create-event-usecase'
import { CreateLocationUseCase } from '@/modules/events/usecases/events/create-location-usecase'
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

const controller = new EventController(createEvent, createLocation, createTicketOption)
const middleware = new AuthMiddleware(new AuthValidation(new JwtEncrypter()))

router.post('/events', middleware.handle, controller.create)
router.get('/events')
router.get('/events/:id')
router.get('/events/nearby')
router.put('/events/:id')

export { router as eventsRouter }

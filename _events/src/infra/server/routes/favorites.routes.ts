import { JwtEncrypter } from '@/infra/crypto/jwt-encrypter'
import { TypeORMDataSource } from '@/infra/db/data-source'
import { PgFavoritesRepository } from '@/infra/db/repositories/favorites-repository'
import { AuthValidation } from '@/modules/accounts/auth/validate-auth'
import { FavoriteEventController } from '@/modules/events/controllers/events/favorite-event-controller'
import { Router } from 'express'
import { AuthMiddleware } from '../middlewares/auth-middleware'

const router = Router()

const repository = new PgFavoritesRepository(TypeORMDataSource.getDataSource())
const controller = new FavoriteEventController(repository)
const middleware = new AuthMiddleware(new AuthValidation(new JwtEncrypter()))

router.post('/favorites', middleware.handle, controller.favorite)
router.get('/favorites', middleware.handle, controller.findByAccount)
router.get('/favorites/:id/count', middleware.handle, controller.count)
router.delete('/favorites/:id', middleware.handle, controller.delete)

export { router as favoritesRouter }

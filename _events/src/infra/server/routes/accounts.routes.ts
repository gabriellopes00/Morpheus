import { Argon2Hasher } from '@/infra/crypto/argon2-hasher'
import { JwtEncrypter } from '@/infra/crypto/jwt-encrypter'
import { TypeORMDataSource } from '@/infra/db/data-source'
import { PgAccountRepository } from '@/infra/db/repositories/account-repository'
import { AuthValidation } from '@/modules/accounts/auth/validate-auth'
import { AccountController } from '@/modules/accounts/controllers/account-controller'
import { CreateAccount } from '@/modules/accounts/usecases/create-account'
import { FindAccountById } from '@/modules/accounts/usecases/find-account-by-id'
import { UpdateAccount } from '@/modules/accounts/usecases/update-account'
import { Router } from 'express'
import { AuthMiddleware } from '../middlewares/auth-middleware'

const router = Router()

const accountRepository = new PgAccountRepository(TypeORMDataSource.getDataSource())
const argon2Hasher = new Argon2Hasher()
const createAccount = new CreateAccount(accountRepository, argon2Hasher)
const updateAccount = new UpdateAccount(accountRepository) // upload image
const findAccount = new FindAccountById(accountRepository)
const controller = new AccountController(createAccount, findAccount, updateAccount)
const middleware = new AuthMiddleware(new AuthValidation(new JwtEncrypter()))

router.post('/accounts', controller.create)
router.get('/accounts/me', middleware.handle, controller.me)
router.put('/accounts/:id', middleware.handle, controller.update)

export { router as accountRoutes }

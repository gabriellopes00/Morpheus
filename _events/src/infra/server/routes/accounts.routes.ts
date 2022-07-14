import { Argon2Hasher } from '@/infra/crypto/argon2-hasher'
import { TypeORMDataSource } from '@/infra/db/data-source'
import { PgAccountRepository } from '@/infra/db/repositories/account-repository'
import { AccountController } from '@/modules/accounts/controllers/account-controller'
import { CreateAccount } from '@/modules/accounts/usecases/create-account'
import { FindAccountById } from '@/modules/accounts/usecases/find-account-by-id'
import { UpdateAccount } from '@/modules/accounts/usecases/update-account'
import { Router } from 'express'

const router = Router()

const accountRepository = new PgAccountRepository(TypeORMDataSource.getDataSource())
const argon2Hasher = new Argon2Hasher()
const createAccount = new CreateAccount(accountRepository, argon2Hasher)
const updateAccount = new UpdateAccount(accountRepository) // upload image
const findAccount = new FindAccountById(accountRepository)
const controller = new AccountController(createAccount, findAccount, updateAccount)

router.post('/accounts', controller.create)
router.get('/accounts/me', controller.me)
router.get('/accounts/:id', controller.findById)
router.put('/accounts/:id', controller.update)

export { router as accountRoutes }

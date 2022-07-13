import { Argon2Hasher } from '@/infra/crypto/argon2-hasher'
import { JwtEncrypter } from '@/infra/crypto/jwt-encrypter'
import { LoginAccount } from '@/modules/accounts/auth/login-user'
import { AuthController } from '@/modules/accounts/controllers/auth-controller'
import { Router } from 'express'

const router = Router()

const loginAccount = new LoginAccount(null, new JwtEncrypter(), new Argon2Hasher())
const controller = new AuthController(loginAccount)

router.post('/auth', controller.login)

export { router as authRoutes }

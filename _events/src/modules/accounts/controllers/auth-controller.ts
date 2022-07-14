import logger from '@/config/logger'
import { LoginAccount, LoginAccountParams } from '@/modules/accounts/auth/login-user'
import { Request, Response } from 'express'

export interface AuthRequest<T = any> extends Request {
  body: T
}

export class AuthController {
  constructor(private readonly loginAccount: LoginAccount) {
    const methods = Object.getOwnPropertyNames(Object.getPrototypeOf(this))
    methods.filter(m => m !== 'constructor').forEach(m => (this[m] = this[m].bind(this)))
  }

  public async login(req: AuthRequest<LoginAccountParams>, res: Response): Promise<Response> {
    try {
      const { email, password } = req.body

      const loginResult = await this.loginAccount.execute({ email, password })
      if (loginResult instanceof Error) return res.status(400).json({ error: loginResult.message })

      return res.status(200).json({
        auth: {
          accessToken: loginResult.accessToken,
          account: {
            id: loginResult.account.id,
            name: loginResult.account.name,
            email: loginResult.account.email,
            document: loginResult.account.document,
            gender: loginResult.account.gender,
            birthDate: loginResult.account.birthDate,
            avatarUrl: loginResult.account.avatarUrl
          }
        }
      })
    } catch (error) {
      logger.error(error, 'internal server error')
      return res.status(500).json({ error: 'Error interno do servidor. Tente novamente...' })
    }
  }
}

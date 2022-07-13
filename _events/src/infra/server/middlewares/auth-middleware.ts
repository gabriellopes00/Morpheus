import { ExpiredTokenError } from '@/core/infra/encrypter'
import { AuthValidation } from '@/modules/accounts/auth/validate-auth'
import { NextFunction, Request, Response } from 'express'

export class AuthMiddleware {
  constructor(private readonly authValidation: AuthValidation) {
    this.handle = this.handle.bind(this)
  }

  public async handle(req: Request, res: Response, next: NextFunction) {
    const token = req.headers?.authorization?.split(' ')[1]
    if (!token) return res.status(401).json({ error: 'Token de autenticação necessário' })

    const authResult = await this.authValidation.execute({ token })
    if (authResult instanceof ExpiredTokenError) {
      return res.status(401).json({ error: authResult.message, code: 'token.expired' })
    } else if (authResult instanceof Error) {
      return res.status(401).json({ error: authResult.message })
    }

    Object.assign(req.headers, { account_id: authResult.accountId })

    next()
  }
}

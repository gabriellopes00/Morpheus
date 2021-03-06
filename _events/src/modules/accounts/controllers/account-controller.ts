import logger from '@/config/logger'
import { Request, Response } from 'express'
import { CreateAccount, CreateAccountParams } from '../usecases/create-account'
import { FindAccountById } from '../usecases/find-account-by-id'
import { UpdateAccount, UpdateAccountParams } from '../usecases/update-account'

export interface AccountRequest<T = any> extends Request {
  body: T
}

export class AccountController {
  constructor(
    private readonly createAccount: CreateAccount,
    private readonly findAccountById: FindAccountById,
    private readonly updateAccount: UpdateAccount
  ) {
    const methods = Object.getOwnPropertyNames(Object.getPrototypeOf(this))
    methods.filter(m => m !== 'constructor').forEach(m => (this[m] = this[m].bind(this)))
  }

  async create(req: AccountRequest<CreateAccountParams>, res: Response) {
    try {
      const { name, document, email, birthDate, gender, avatarUrl, password } = req.body

      const result = await this.createAccount.execute({
        name,
        document,
        email,
        birthDate,
        gender,
        avatarUrl,
        password
      })
      if (result instanceof Error) return res.status(400).json({ error: result.message })

      return res.status(201).json({
        account: {
          id: result.id,
          name: result.name,
          email: result.email,
          document: result.document,
          gender: result.gender,
          birthDate: result.birthDate,
          avatarUrl: result.avatarUrl
        }
      })
    } catch (error) {
      logger.error(error, 'internal server error')
      return res.status(500).json({ error: 'Error interno do servidor. Tente novamente...' })
    }
  }

  async update(req: AccountRequest<UpdateAccountParams>, res: Response) {
    try {
      const accountId = String(req.headers.account_id)
      const id = req.params.id

      if (accountId !== id) return res.status(403).json({ error: 'Atualização não autorizada' })

      const { name, avatarUrl, gender } = req.body

      const result = await this.updateAccount.execute({
        id: accountId,
        name,
        avatarUrl,
        gender
      })
      if (result instanceof Error) {
        return res.status(400).json({ error: result.message })
      }

      delete result.password

      return res.status(200).json({
        account: {
          id: result.id,
          name: result.name,
          email: result.email,
          document: result.document,
          gender: result.gender,
          birthDate: result.birthDate,
          avatarUrl: result.avatarUrl
        }
      })
    } catch (error) {
      logger.error(error, 'internal server error')
      return res.status(500).json({ error: 'Error interno do servidor. Tente novamente...' })
    }
  }

  async me(req: Request, res: Response) {
    try {
      const id = String(req.headers.account_id)
      const result = await this.findAccountById.execute({ id })

      if (result instanceof Error) return res.status(404).json({ error: result.message })
      else if (result === null) return res.status(404).json({ error: 'Usuário não encontrado' })

      delete result.password

      return res.status(200).json({
        account: {
          id: result.id,
          name: result.name,
          email: result.email,
          document: result.document,
          gender: result.gender,
          birthDate: result.birthDate,
          avatarUrl: result.avatarUrl
        }
      })
    } catch (error) {
      logger.error(error, 'internal server error')
      return res.status(500).json({ error: 'Error interno do servidor. Tente novamente...' })
    }
  }
}

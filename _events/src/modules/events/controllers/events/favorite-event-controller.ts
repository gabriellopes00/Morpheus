import { PgFavoritesRepository } from '@/infra/db/repositories/favorites-repository'
import { Request, Response } from 'express'

export class FavoriteEventController {
  constructor(private readonly repository: PgFavoritesRepository) {
    const methods = Object.getOwnPropertyNames(Object.getPrototypeOf(this))
    methods.filter(m => m !== 'constructor').forEach(m => (this[m] = this[m].bind(this)))
  }

  public async favorite(req: Request, res: Response): Promise<Response> {
    try {
      const accountId = String(req.headers.account_id)
      const { eventId } = req.body
      await this.repository.save(accountId, eventId)
      return res.sendStatus(204)
    } catch (error) {
      return res.status(500).json({ error: 'Error interno do servidor. Tente novamente...' })
    }
  }

  public async count(req: Request, res: Response): Promise<Response> {
    try {
      const eventId = req.params.id
      const total = await this.repository.count(eventId)
      return res.status(200).json({ total })
    } catch (error) {
      return res.status(500).json({ error: 'Error interno do servidor. Tente novamente...' })
    }
  }

  public async findByAccount(req: Request, res: Response): Promise<Response> {
    try {
      const accountId = String(req.headers.account_id)
      const events = await this.repository.findByAccount(accountId)
      return res.status(200).json({ events })
    } catch (error) {
      return res.status(500).json({ error: 'Error interno do servidor. Tente novamente...' })
    }
  }

  public async delete(req: Request, res: Response): Promise<Response> {
    try {
      const accountId = String(req.headers.account_id)
      const { eventId } = req.body
      await this.repository.delete(accountId, eventId)
      return res.sendStatus(204)
    } catch (error) {
      return res.status(500).json({ error: 'Error interno do servidor. Tente novamente...' })
    }
  }
}

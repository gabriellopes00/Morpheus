import logger from '@/config/logger'
import { env } from 'process'
import { ConnectionNotFoundError, DataSource } from 'typeorm'
import { AccountEntity } from './entities/account-entity'
import { EventEntity } from './entities/event-entity'
import { FavoriteEntity } from './entities/favorite-entity'
import { CreateAccountMigration } from './migrations/1653147077419-create-account-migration'
import { CreateCategoriesMigration } from './migrations/1657896118944-create-categories-migration'
import { CreateSubjectsMigration } from './migrations/1657896596925-create-subjects-migrations'
import { CreateEventsMigration } from './migrations/1657896600958-create-events-migrations'
import { CreateLocationsMigration } from './migrations/1657896603908-create-locations-migrations'
import { CreateTicketOptionsMigration } from './migrations/1657896606712-create-ticket-options-migrations'
import { CreateFavoritesMigration } from './migrations/1658315610357-create-favorites'

const { DB_HOST, DB_USER, DB_NAME, DB_PASS, DB_PORT } = env

export class TypeORMDataSource {
  private static dataSource: DataSource

  private constructor() {}

  static getDataSource(): DataSource {
    if (TypeORMDataSource.dataSource === undefined) {
      this.setDataSource()
    }

    return TypeORMDataSource.dataSource
  }

  static async connect(): Promise<void> {
    this.setDataSource()
    await this.dataSource.initialize()
    logger.info('Database connected successfully')
  }

  private static setDataSource() {
    this.dataSource = new DataSource({
      type: 'postgres',
      host: DB_HOST,
      port: Number(DB_PORT),
      ssl: { rejectUnauthorized: false },
      username: DB_USER,
      password: DB_PASS,
      database: DB_NAME,
      entities: [AccountEntity, FavoriteEntity, EventEntity],
      migrationsTableName: '_migrations',
      migrations: [
        CreateAccountMigration,
        CreateCategoriesMigration,
        CreateSubjectsMigration,
        CreateEventsMigration,
        CreateLocationsMigration,
        CreateTicketOptionsMigration,
        CreateFavoritesMigration
      ]
    })
  }

  static async disconnect(): Promise<void> {
    if (this.dataSource === undefined) throw new ConnectionNotFoundError('default')
    await this.dataSource.destroy()
    this.dataSource = undefined
    logger.info('Database disconnected successfully')
  }
}

import logger from '@/config/logger'
import { env } from 'process'
import { ConnectionNotFoundError, DataSource } from 'typeorm'
import { AccountEntity } from './entities/account-entity'

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
      username: DB_USER,
      password: DB_PASS,
      database: DB_NAME,
      entities: [AccountEntity],
      migrationsTableName: '_migrations'
    })
  }

  static async disconnect(): Promise<void> {
    if (this.dataSource === undefined) throw new ConnectionNotFoundError('default')
    await this.dataSource.destroy()
    this.dataSource = undefined
    logger.info('Database disconnected successfully')
  }
}

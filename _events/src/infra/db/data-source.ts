import 'dotenv/config'
import { env } from 'process'
import 'reflect-metadata'
import { DataSource } from 'typeorm'
import { AccountEntity } from './entities/account-entity'
import { CreateAccountMigration } from './migrations/create-account-migration'

const { DB_HOST, DB_USER, DB_PASS, DB_NAME, DB_PORT } = env

export const PostgresDataSource = new DataSource({
  type: 'postgres',
  host: DB_HOST,
  port: Number(DB_PORT),
  username: DB_USER,
  password: DB_PASS,
  database: DB_NAME,
  ssl: { rejectUnauthorized: false },
  entities: [AccountEntity],
  migrations: [CreateAccountMigration],
  migrationsTableName: '_migrations'
})

PostgresDataSource.initialize().then(async d => await d.runMigrations())

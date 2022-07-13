import 'dotenv/config'
import 'module-alias/register'
import 'reflect-metadata'

import { TypeORMDataSource } from '@/infra/db/data-source'
import { env } from 'process'
import logger from '@/config/logger'

async function bootstrap() {
  await TypeORMDataSource.connect()
  await TypeORMDataSource.getDataSource().runMigrations()

  const app = (await import('../infra/server/express')).app

  const server = app.listen(env.PORT, () => {
    logger.info('Server started successfully')
  })

  // eslint-disable-next-line
  const exitSignals: NodeJS.Signals[] = ['SIGINT', 'SIGTERM', 'SIGQUIT']
  for (const signal of exitSignals) {
    process.on(signal, async () => {
      try {
        await TypeORMDataSource.disconnect()
        server.close(err => err && logger.error(`Server stopped with error: ${err}`))
        logger.info('Server stopped successfully')
        process.exit(0)
      } catch (error) {
        logger.error(`App exited with error: ${error}`)
        process.exit(1)
      }
    })
  }
}

process.on('unhandledRejection', (reason, promise) => {
  logger.error(`App exiting due an unhandled promise: ${promise} and reason: ${reason}`)
  throw reason
})

process.on('uncaughtException', error => {
  logger.error(`App exiting due to an uncaught exception: ${error}`)
  process.exit(0)
})

bootstrap()

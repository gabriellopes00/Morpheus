module.exports = {
  type: process.env.DB_DRIVER,
  url: process.env.DB_URL,
  timezone: process.env.DB_TIME_ZONE,
  synchronize: false,
  logging: false,
  migrations: [__dirname + '/src/framework/database/migrations/*.ts'],
  entities: [__dirname + '/src/framework/database/models/*.ts'],
  cli: { migrationsDir: __dirname + '/src/framework/database/migrations/' }
}

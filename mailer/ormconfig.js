module.exports = {
  type: 'postgres',
  url: process.env.DB_URL,
  // ssl: { rejectUnauthorized: false },
  synchronize: false,
  logging: false,

  migrations: [__dirname + '/dist/framework/database/migrations/*.js'],
  entities: [__dirname + '/dist/framework/database/entities/*.js'],
  cli: { migrationsDir: __dirname + '/src/framework/database/migrations/' }
}

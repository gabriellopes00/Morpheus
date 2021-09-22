module.exports = {
  type: 'postgres',
  host: process.env.DB_HOST,
  port: process.env.DB_PORT,
  username: process.env.DB_USER,
  password: process.env.DB_PASS,
  database: process.env.DB_NAME,
  ssl: { rejectUnauthorized: false },
  synchronize: false,
  logging: false,

  migrations: [__dirname + '/dist/framework/database/migrations/*.js'],
  entities: [__dirname + '/dist/framework/database/entities/*.js'],
  cli: { migrationsDir: __dirname + '/src/framework/database/migrations/' }
}

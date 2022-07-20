import express from 'express'
import { accountRoutes } from './routes/accounts.routes'
import { authRoutes } from './routes/auth.routes'
import { eventsRouter } from './routes/events.routes'
import { favoritesRouter } from './routes/favorites.routes'

const app = express()

app.use(express.json())

app.use(authRoutes)
app.use(accountRoutes)
app.use(favoritesRouter)
app.use(eventsRouter)

export { app }

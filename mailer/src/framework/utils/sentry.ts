import * as Sentry from '@sentry/node'

const { SENTRY_DSN } = process.env

Sentry.init({ dsn: SENTRY_DSN, tracesSampleRate: 1.0 })

export { Sentry }

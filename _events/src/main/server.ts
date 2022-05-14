import 'module-alias/register'
import 'dotenv/config'
import logger from '../config/logger'

logger.info(process.env.PORT)

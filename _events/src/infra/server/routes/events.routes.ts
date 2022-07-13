import { Router } from 'express'

const router = Router()

router.post('/events')
router.get('/events')
router.get('/events/:id')
router.get('/events/nearby')
router.put('/events/:id')

export { router as eventsRouter }

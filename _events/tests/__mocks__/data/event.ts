import { EventVisibility } from '@/modules/events/domain/event'
import { CreateEventCredentials } from '@/modules/events/usecases/events/create-event-usecase'
import { faker } from '@faker-js/faker'

export const event: CreateEventCredentials = {
  name: faker.name.jobTitle(),
  description: faker.lorem.lines(),
  ageGroup: faker.helpers.arrayElement([0, 10, 12, 16, 18]),
  coverUrl: faker.image.imageUrl(),
  startDateTime: new Date('2022-05-29T13:21:36.639Z'),
  organizerAccountId: crypto.randomUUID(),
  endDateTime: new Date('2022-05-29T15:21:36.639Z'),
  subjectId: crypto.randomUUID(),
  categoryId: crypto.randomUUID(),
  visibility: faker.helpers.arrayElement<EventVisibility>(['public', 'private'])
}

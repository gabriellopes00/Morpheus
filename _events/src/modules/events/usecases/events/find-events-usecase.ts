import { FindRepository } from '@/shared/repositories'
import { Event } from '../../domain/event'

export interface FindEventsResult {
  [key: string]: Event[]
}

export interface Category {
  id: string
  name: string
}

export class FindEventsUseCase {
  constructor(
    private readonly repository: FindRepository<Event>,
    private readonly categoryRepository: FindRepository<Category>
  ) {}

  public async execute(): Promise<FindEventsResult> {
    const result: FindEventsResult = {}
    const categories = await this.categoryRepository.findAll()

    for (const category of categories) {
      const events = await this.repository.findAllBy('categoryId', category.id)
      result[category.name] = events
    }

    return result
  }
}

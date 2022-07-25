import { Category } from '@/modules/events/usecases/events/find-events-usecase'
import { DataSource } from 'typeorm'

export class CategoryRepository {
  constructor(private readonly dataSource: DataSource) {}

  public async findAll(): Promise<Category[]> {
    const query = this.dataSource.createQueryBuilder()

    const data = await query.select('*').from('categories', 'categories').execute()
    return data.map(e => ({ id: e.id, name: e.name }))
  }
}

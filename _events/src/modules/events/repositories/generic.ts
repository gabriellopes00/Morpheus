export interface SaveRepository<T = never> {
  save(event: T): Promise<void>
  saveAll(events: T[]): Promise<void>
}

export interface FindRepository<T = never> {
  find<I = any, S = string>(term: S): Promise<I | T>
  exists(term: T | { id: string }): Promise<boolean>
  findAll(): Promise<T[]>
}

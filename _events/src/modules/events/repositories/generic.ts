export interface SaveRepository<T = never> {
  save<I = any>(event: I | T): Promise<void>
  saveAll<I = any>(events: I[] | T[]): Promise<void>
}

export interface FindRepository<T = never> {
  find<I = any, S = string>(term: S): Promise<I | T>
  exists<S = string>(term: S): Promise<boolean>
  findAll<I = any>(): Promise<I[] | T[]>
}

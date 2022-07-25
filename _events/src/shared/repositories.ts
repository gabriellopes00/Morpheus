export interface SaveRepository<T = never> {
  save?(event: T): Promise<void>
  saveAll?(events: T[]): Promise<void>
}

export interface FindRepository<T = never> {
  findBy?(key: keyof T, value: any): Promise<T>
  findAllBy?(key: keyof T, value: any): Promise<T[]>
  exists?({ id: string }): Promise<boolean>
  findAll?(): Promise<T[]>
}

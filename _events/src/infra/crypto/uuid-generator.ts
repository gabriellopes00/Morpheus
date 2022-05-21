import { UUIDGenerator } from '@/core/infra/uuid-generator'
import { v4 as uuid } from 'uuid'

export class UidGenerator implements UUIDGenerator {
  generate(): string {
    return uuid()
  }
}

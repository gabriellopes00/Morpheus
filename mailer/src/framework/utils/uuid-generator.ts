import { UUIDGenerator } from '@/ports/uuid-generator'
import { v4 as uuid } from 'uuid'

export class V4UUIDGenerator implements UUIDGenerator {
  public generate() {
    return uuid()
  }
}

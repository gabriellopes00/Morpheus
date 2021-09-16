import { V4UUIDGenerator } from '@/framework/utils/uuid-generator'
import { UUIDGenerator } from '@/ports/uuid-generator'

export function createUUIDGenerator(): UUIDGenerator {
  return new V4UUIDGenerator()
}

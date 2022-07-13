import * as Yup from 'yup'

export class YupValidator {
  constructor(private readonly schema: Yup.AnyObjectSchema) {}

  async validate<T = any>(input: T) {
    try {
      await this.schema.validate(input, { stripUnknown: true })
      return null
    } catch (error) {
      return error
    }
  }
}

export interface Validator {
  validate<T = any>(input: T): Promise<Error>
}

export interface ValidationError extends Error {}

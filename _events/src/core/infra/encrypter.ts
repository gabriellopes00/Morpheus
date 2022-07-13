export class ExpiredTokenError extends Error {
  constructor() {
    super('Token expirado')
  }
}
/**
 * Encrypter is responsible by generating(encrypting) and decrypting access tokens.
 * Accounts payloads will be encrypted in a single token, and it will be decrypted
 * to validate the user authenticity.
 */
export interface Encrypter {
  encrypt<T = any>(payload: T, expirationTime: string | 0): Promise<string>
  decrypt<T = any>(token: string): Promise<T | ExpiredTokenError>
}

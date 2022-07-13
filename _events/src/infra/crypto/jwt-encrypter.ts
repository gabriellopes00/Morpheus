import { Encrypter, ExpiredTokenError } from '@/core/infra/encrypter'
import { sign, verify, TokenExpiredError } from 'jsonwebtoken'
import { env } from 'process'

export class JwtEncrypter implements Encrypter {
  public async encrypt(payload: any, expirationTime: string | 0): Promise<string> {
    return sign({ ...payload }, env.ENCRYPTION_KEY, {
      algorithm: 'HS256',
      expiresIn: expirationTime
    })
  }

  public async decrypt<T = any>(token: string): Promise<T | Error> {
    try {
      return verify(token, env.ENCRYPTION_KEY, { algorithms: ['HS256'] }) as T
    } catch (error) {
      if (error instanceof TokenExpiredError) return new ExpiredTokenError()
      return null
    }
  }
}

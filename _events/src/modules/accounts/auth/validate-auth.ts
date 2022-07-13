import { Encrypter, ExpiredTokenError } from '@/core/infra/encrypter'

export interface AuthValidationResult {
  accountId: number
}

export interface AuthValidationParams {
  token: string
}

export interface AccessTokenPayload {
  id: number
}

export class AuthValidation {
  constructor(private readonly encrypter: Encrypter) {}

  public async execute(params: AuthValidationParams): Promise<AuthValidationResult | Error> {
    const tokenData = await this.encrypter.decrypt<AccessTokenPayload>(params.token)
    if (tokenData instanceof ExpiredTokenError) return tokenData

    return { accountId: tokenData.id }
  }
}

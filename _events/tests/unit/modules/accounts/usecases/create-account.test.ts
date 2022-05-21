import { UidGenerator } from '@/infra/crypto/uuid-generator'
import { CreateAccount } from '@/modules/accounts/usecases/create-account'
import { MockAccountRepository } from '@t/__mocks__/repositories/mock-account-repository'

describe('Create Account', () => {
  const mockRepository = new MockAccountRepository()
  const usecase = new CreateAccount(new UidGenerator(), mockRepository)

  beforeAll(() => {
    jest.useFakeTimers('modern')
    jest.setSystemTime(new Date(2022, 3, 1))
  })

  afterAll(() => jest.useRealTimers())

  it('Should create an account successfully', async () => {
    const referencedId = '292b6a92-da4d-4661-b5c8-f3cb5a1527b8'
    await usecase.execute(referencedId)
    const account = mockRepository.rows.find(a => a.referencedId === referencedId)
    expect(account.id).toMatch(/^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$/)
    expect(account.referencedId).toEqual(referencedId)
    expect(account.createdAt).toEqual(new Date())
  })
})

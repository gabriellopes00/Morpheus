import { Account } from '@/modules/accounts/domain/account'
import { DeleteAccount } from '@/modules/accounts/usecases/delete-account'
import { MockAccountRepository } from '@t/__mocks__/repositories/mock-account-repository'

describe('Delete Account', () => {
  const mockRepository = new MockAccountRepository()
  const usecase = new DeleteAccount(mockRepository)
  const referencedId = '292b6a92-da4d-4661-b5c8-f3cb5a1527b8'

  beforeAll(() => {
    jest.useFakeTimers('modern')
    jest.setSystemTime(new Date(2022, 3, 1))
  })

  afterAll(() => jest.useRealTimers())

  async function _insertMockAccount() {
    mockRepository.rows.push(new Account({ referencedId }, 'id'))
  }

  it('Should create an account successfully', async () => {
    await _insertMockAccount()
    await usecase.execute(referencedId)
    const index = mockRepository.rows.findIndex(a => a.referencedId === referencedId)
    expect(index).toEqual(-1)
  })
})

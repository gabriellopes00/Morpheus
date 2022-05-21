export interface DeleteAccountRepository {
  delete(referencedId: string): Promise<void>
}

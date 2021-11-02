package usecases

type DeleteAccount interface {
	Delete(accountId string) error
}

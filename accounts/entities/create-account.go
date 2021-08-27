package entities

type CreateAccount interface {
	Create(data Account) (*Account, error)
}

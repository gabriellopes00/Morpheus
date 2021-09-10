package entities

type Account struct {
	Id string `json:"account_id"`
}

func NewAccount(id string) *Account {
	return &Account{
		Id: id,
	}
}

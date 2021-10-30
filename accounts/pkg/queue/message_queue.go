package queue

const (
	KeyAccountCreated = "account_created"
	KeyAccountDeleted = "account_deleted"

	ExchangeAccounts = "accounts_ex"
)

type MessageQueue interface {
	SendMessage(exchange, routingKey string, payload []byte) error
}

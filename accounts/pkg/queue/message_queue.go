package queue

const (
	KeyAccountCreated = "account_created"
	KeyAccountDeleted = "account_deleted"
	KeyAccountUpdated = "account_updated"

	ExchangeAccounts = "accounts_ex"
)

type MessageQueue interface {
	PublishMessage(exchange, key string, payload []byte) error
}

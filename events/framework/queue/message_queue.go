package queue

const (
	KeyEventCreated   = "event_created"
	KeyAccountCreated = "account_created"
	KeyAccountDeleted = "account_deleted"

	ExchangeEvents   = "events_ex"
	ExchangeAccounts = "accounts_ex"

	QueueAccount = "accounts"
)

type MessageQueue interface {
	Consume(queue string, channel chan<- []byte)
	PublishMessage(exchange, routingKey string, message []byte) error
}

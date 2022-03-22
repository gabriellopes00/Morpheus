package queue

const (
	KeyEventCreated  = "event_created"
	KeyEventUpdated  = "event_updated"
	KeyEventSoldOut  = "event_sold_out"
	KeyEventCanceled = "event_canceled"

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

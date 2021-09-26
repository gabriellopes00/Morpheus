package queue

const (
	QueueAccountCreated = "account_created"
	QueueAccountUpdated = "account_updated"
	QueueAccountDeleted = "account_deleted"
)

type MessageQueue interface {
	SendMessage(queue string, payload []byte) error
}

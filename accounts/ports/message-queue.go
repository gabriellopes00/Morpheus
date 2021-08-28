package ports

type MessageQueue interface {
	SendMessage(payload []byte) error
}

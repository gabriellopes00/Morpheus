package interfaces

type MessageQueue interface {
	SendMessage(payload []byte) error
}

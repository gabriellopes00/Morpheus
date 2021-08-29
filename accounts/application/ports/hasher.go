package ports

type Hasher interface {
	GenHash(payload string) (string, error)
}

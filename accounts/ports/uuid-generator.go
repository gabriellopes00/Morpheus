package ports

type UUIDGenerator interface {
	Generate() (string, error)
}

package storage

import "io"

type Provider interface {
	UploadFile(file io.Reader, filename string) (string, error)
}

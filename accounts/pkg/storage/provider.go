package storage

import "io"

type Provider interface {
	UploadFile(file io.Reader, filename, filetype string) (string, error)
}

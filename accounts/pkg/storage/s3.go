package storage

import (
	"accounts/config/env"
	"io"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type S3StorageProvider struct {
	session *session.Session
}

func NewS3StorageProvider(awsSession *session.Session) *S3StorageProvider {
	return &S3StorageProvider{
		session: awsSession,
	}
}

func (s3 *S3StorageProvider) UploadFile(file io.Reader, filename, filetype string) (string, error) {
	uploader := s3manager.NewUploader(s3.session)
	up, err := uploader.Upload(&s3manager.UploadInput{
		Bucket:             aws.String(env.AWS_S3_BUCKET_NAME),
		Key:                aws.String(filename),
		Body:               file,
		ContentType:        aws.String(filetype),
		ContentDisposition: aws.String("inline"),
		ACL:                aws.String("public-read"),
	})
	if err != nil {
		return "", err
	}

	return up.Location, nil
}

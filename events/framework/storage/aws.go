package storage

import (
	"events/config/env"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

func CreateAWSSession() (*session.Session, error) {
	config := aws.Config{
		Region: aws.String(env.AWS_S3_REGION),
		Credentials: credentials.NewStaticCredentials(
			env.AWS_S3_ACCESS_KEY_ID,
			env.AWS_S3_SECRET_ACCESS_KEY,
			"",
		),
	}

	return session.NewSession(&config)
}

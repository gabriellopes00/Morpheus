package env

import (
	"accounts/pkg/logger"
	"errors"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

var (
	DB_USER      = ""
	DB_PASS      = ""
	DB_PORT      = 0
	DB_HOST      = ""
	DB_NAME      = ""
	DB_DRIVER    = ""
	DB_SSL_MODE  = ""
	DB_TIME_ZONE = ""

	REDIS_HOST = ""
	REDIS_PASS = ""
	REDIS_PORT = 0

	PORT = 0

	RABBITMQ_HOST  = ""
	RABBITMQ_USER  = ""
	RABBITMQ_PASS  = ""
	RABBITMQ_PORT  = 0
	RABBITMQ_VHOST = ""

	KEYCLOACK_HOST           = ""
	KEYCLOACK_PORT           = 0
	KEYCLOACK_REALM          = ""
	KEYCLOACK_CLIENT_ID      = ""
	KEYCLOACK_CLIENT_SECRET  = ""
	KEYCLOACK_PUBLIC_RSA_KEY = ""

	AWS_S3_REGION            = ""
	AWS_S3_ACCESS_KEY_ID     = ""
	AWS_S3_SECRET_ACCESS_KEY = ""
	AWS_S3_BUCKET_NAME       = ""
)

var ErrLoadingEnvVars = errors.New("error on loading environment variables")

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln(ErrLoadingEnvVars)
	}

	PORT, err = strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		logger.Logger.Error(ErrLoadingEnvVars.Error(),
			zap.String("error_message", err.Error()),
			zap.String("variable", "PORT"),
		)
	}

	DB_USER = os.Getenv("DB_USER")
	DB_PASS = os.Getenv("DB_PASS")
	DB_HOST = os.Getenv("DB_HOST")
	DB_NAME = os.Getenv("DB_NAME")
	DB_DRIVER = os.Getenv("DB_DRIVER")
	DB_SSL_MODE = os.Getenv("DB_SSL_MODE")
	DB_TIME_ZONE = os.Getenv("DB_TIME_ZONE")
	DB_PORT, err = strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		logger.Logger.Error(ErrLoadingEnvVars.Error(),
			zap.String("error_message", err.Error()),
			zap.String("variable", "DB_PORT"),
		)
	}

	REDIS_HOST = os.Getenv("REDIS_HOST")
	REDIS_PASS = os.Getenv("REDIS_PASS")
	REDIS_PORT, err = strconv.Atoi(os.Getenv("REDIS_PORT"))
	if err != nil {
		logger.Logger.Error(ErrLoadingEnvVars.Error(),
			zap.String("error_message", err.Error()),
			zap.String("variable", "REDIS_PORT"),
		)
	}

	RABBITMQ_HOST = os.Getenv("RABBITMQ_HOST")
	RABBITMQ_USER = os.Getenv("RABBITMQ_USER")
	RABBITMQ_PASS = os.Getenv("RABBITMQ_PASS")
	RABBITMQ_VHOST = os.Getenv("RABBITMQ_VHOST")
	RABBITMQ_PORT, err = strconv.Atoi(os.Getenv("RABBITMQ_PORT"))
	if err != nil {
		logger.Logger.Error(ErrLoadingEnvVars.Error(),
			zap.String("error_message", err.Error()),
			zap.String("variable", "RABBITMQ_PORT"),
		)
	}

	KEYCLOACK_CLIENT_ID = os.Getenv("KEYCLOACK_CLIENT_ID")
	KEYCLOACK_CLIENT_SECRET = os.Getenv("KEYCLOACK_CLIENT_SECRET")
	KEYCLOACK_HOST = os.Getenv("KEYCLOACK_HOST")
	KEYCLOACK_REALM = os.Getenv("KEYCLOACK_REALM")
	KEYCLOACK_PUBLIC_RSA_KEY = os.Getenv("KEYCLOACK_PUBLIC_RSA_KEY")
	KEYCLOACK_PORT, err = strconv.Atoi(os.Getenv("KEYCLOACK_PORT"))
	if err != nil {
		logger.Logger.Error(ErrLoadingEnvVars.Error(),
			zap.String("error_message", err.Error()),
			zap.String("variable", "KEYCLOACK_PORT"),
		)
	}

	AWS_S3_REGION = os.Getenv("AWS_S3_REGION")
	AWS_S3_ACCESS_KEY_ID = os.Getenv("AWS_S3_ACCESS_KEY_ID")
	AWS_S3_SECRET_ACCESS_KEY = os.Getenv("AWS_S3_SECRET_ACCESS_KEY")
	AWS_S3_BUCKET_NAME = os.Getenv("AWS_S3_BUCKET_NAME")

}

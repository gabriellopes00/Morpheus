package env

import (
	"errors"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
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

	AUTH_TOKEN_KEY = ""

	PORT = 0

	RABBITMQ_HOST  = ""
	RABBITMQ_USER  = ""
	RABBITMQ_PASS  = ""
	RABBITMQ_PORT  = 0
	RABBITMQ_VHOST = ""
)

var ErrLoadingEnvVars = errors.New("error on loading environment variables")

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln(ErrLoadingEnvVars)
	}

	PORT, err = strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		log.Fatalln(err)
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
		log.Fatalln(err)
	}

	AUTH_TOKEN_KEY = os.Getenv("AUTH_TOKEN_KEY")

	RABBITMQ_HOST = os.Getenv("RABBITMQ_HOST")
	RABBITMQ_USER = os.Getenv("RABBITMQ_USER")
	RABBITMQ_PASS = os.Getenv("RABBITMQ_PASS")
	RABBITMQ_VHOST = os.Getenv("RABBITMQ_VHOST")
	RABBITMQ_PORT, err = strconv.Atoi(os.Getenv("RABBITMQ_PORT"))
	if err != nil {
		log.Fatalln(err)
	}

}

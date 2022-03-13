package handlers

import (
	"accounts/application"
	"accounts/config/env"
	"fmt"
	"log"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/labstack/echo/v4"
)

var AccessKeyID string
var SecretAccessKey string
var MyRegion string

func ConnectAws() *session.Session {
	sess, err := session.NewSession(
		&aws.Config{
			Region: aws.String(env.AWS_S3_REGION),
			Credentials: credentials.NewStaticCredentials(
				env.AWS_S3_ACCESS_KEY_ID,
				env.AWS_S3_SECRET_ACCESS_KEY,
				"",
			),
		})
	if err != nil {
		log.Fatalln(err)
	}
	return sess
}

type AvatarUploadHandler struct {
	Usecase application.UpdateAccount
}

func NewAvatarUploadHandler(usecase application.UpdateAccount) *AvatarUploadHandler {
	return &AvatarUploadHandler{
		Usecase: usecase,
	}
}

func (h *AvatarUploadHandler) Handle(c echo.Context) error {
	accountId := c.Request().Header.Get("account_id")
	paramId := c.Param("id")

	if accountId != paramId {
		return c.JSON(
			http.StatusForbidden,
			map[string]string{"error": "forbidden access"})
	}

	file, err := c.FormFile("avatar")
	if err != nil {
		return c.JSON(http.StatusBadRequest, "ivalid file upload")
	}

	src, err := file.Open()
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, nil)
	}
	defer src.Close()

	buff := make([]byte, 512)
	_, err = src.Read(buff)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, nil)
	}

	// allow only image files
	switch http.DetectContentType(buff) {
	case "image/jpeg", "image/jpg", "image/png":
	default:
		return c.JSON(http.StatusBadRequest,
			map[string]string{"error": "invalid file type. Accounts avatar must be either jpeg, jpg or png files"})
	}

	sess := ConnectAws()
	uploader := s3manager.NewUploader(sess)
	up, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(env.AWS_S3_BUCKET_NAME),
		Key:    aws.String("avatar-" + accountId),
		Body:   src,
	})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}

	_, err = h.Usecase.Update(
		accountId,
		&application.UpdateAccountDTO{AvatarUrl: up.Location},
	)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}

	return c.JSON(
		http.StatusOK,
		map[string]interface{}{
			"message": "avatar uploaded successfully " + up.Location,
		},
	)
}

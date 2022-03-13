package handlers

import (
	"accounts/application"
	"accounts/pkg/storage"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

type AvatarUploadHandler struct {
	usecase         application.UpdateAccount
	storageProvider storage.Provider
}

func NewAvatarUploadHandler(usecase application.UpdateAccount, storageProvider storage.Provider) *AvatarUploadHandler {
	return &AvatarUploadHandler{
		usecase:         usecase,
		storageProvider: storageProvider,
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
		return c.NoContent(http.StatusBadRequest)
	}
	defer src.Close()

	err = h.validateFileType(file.Header.Get("Content-Type"))
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, err)
	}

	fileName := "avatar-" + accountId + "." + strings.Split(file.Header.Get("Content-Type"), "/")[1]
	fileUrl, err := h.storageProvider.UploadFile(src, fileName)
	if err != nil {
		fmt.Println(err)
		return c.NoContent(http.StatusInternalServerError)
	}

	_, err = h.usecase.Update(
		accountId,
		&application.UpdateAccountDTO{AvatarUrl: fileUrl},
	)
	if err != nil {
		fmt.Println(err)
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.NoContent(http.StatusNoContent)
}

func (*AvatarUploadHandler) validateFileType(fileType string) error {
	switch fileType {
	case "image/jpeg", "image/jpg", "image/png":
		return nil
	default:
		return errors.New("invalid file type. Allowed file types are: \"image/jpeg\", \"image/jpg\", \"image/png\"")
	}
}

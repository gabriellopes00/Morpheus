package handlers

import (
	"accounts/application"
	"accounts/pkg/logger"
	"accounts/pkg/storage"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
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
		return c.JSON(http.StatusBadRequest, "error while uploading file")
	}

	src, err := file.Open()
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	defer src.Close()

	filetype := file.Header.Get("Content-Type")
	err = h.validateFileType(filetype)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	fileName := "avatar-" + accountId
	fileUrl, err := h.storageProvider.UploadFile(src, fileName, filetype)
	if err != nil {
		logger.Logger.Error("error while uploading file in provider", zap.String("error_message", err.Error()))
		return c.NoContent(http.StatusInternalServerError)
	}

	_, err = h.usecase.Update(
		accountId,
		&application.UpdateAccountDTO{AvatarUrl: fileUrl},
	)
	if err != nil {
		logger.Logger.Error("error while updating account", zap.String("error_message", err.Error()))
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.String(http.StatusOK, fileUrl)
}

func (*AvatarUploadHandler) validateFileType(fileType string) error {
	switch fileType {
	case "image/jpeg", "image/jpg", "image/png":
		return nil
	default:
		return errors.New("invalid file type. Allowed file types are: \"image/jpeg\", \"image/jpg\", \"image/png\"")
	}
}

package handlers

import (
	"errors"
	"events/application"
	"events/framework/logger"
	"events/framework/storage"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type CoverUploadHandler struct {
	findEvents      *application.FindEvents
	updateEvent     *application.UpdateEvent
	storageProvider storage.Provider
}

func NewCoverUploadHandler(findEvents *application.FindEvents, updateEvent *application.UpdateEvent, storageProvider storage.Provider) *CoverUploadHandler {
	return &CoverUploadHandler{
		storageProvider: storageProvider,
		updateEvent:     updateEvent,
		findEvents:      findEvents,
	}
}

func (h *CoverUploadHandler) Handle(c echo.Context) error {
	accountId := c.Request().Header.Get("account_id")
	id := c.Param("id")

	event, err := h.findEvents.FindEventById(id, false)
	if err != nil {
		return c.JSON(http.StatusConflict, map[string]string{"error": "No events found with given id"})
	}

	if event.OrganizerAccountId != accountId {
		return c.NoContent(http.StatusForbidden)
	}

	file, err := c.FormFile("cover")
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

	fileName := "event-cover-" + id
	fileUrl, err := h.storageProvider.UploadFile(src, fileName, filetype)
	if err != nil {
		logger.Logger.Error("error while uploading file in provider", zap.String("error_message", err.Error()))
		return c.NoContent(http.StatusInternalServerError)
	}

	_, err = h.updateEvent.UpdateData(id, &application.UpdateEventDTO{CoverUrl: fileUrl})
	if err != nil {
		logger.Logger.Error("error while updating account", zap.String("error_message", err.Error()))
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.NoContent(http.StatusNoContent)
}

func (*CoverUploadHandler) validateFileType(fileType string) error {
	switch fileType {
	case "image/jpeg", "image/jpg", "image/png":
		return nil
	default:
		return errors.New("invalid file type. Allowed file types are: \"image/jpeg\", \"image/jpg\", \"image/png\"")
	}
}

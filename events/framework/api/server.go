package api

import (
	"events/application"
	"events/framework/api/handlers"
	"events/framework/api/middlewares"
	"events/framework/auth"
	"events/framework/db/repositories"
	"events/framework/encrypter"
	"events/framework/geocode"
	"events/framework/queue"
	"events/framework/storage"
	"net/http"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/streadway/amqp"
	"gorm.io/gorm"
)

func SetupServer(router *echo.Echo, database *gorm.DB, amqpConn *amqp.Channel, awsSession *session.Session) {

	// init adapters
	eventsRepo := repositories.NewPgEventsRepository(database)
	rabbitMQ := queue.NewRabbitMQ(amqpConn)
	jwtEncrypter := encrypter.NewJwtEncrypter()
	keycloack := auth.NewKeycloackauthProvider(jwtEncrypter)
	geocodeProvider := geocode.NewGeocodeProvider()
	s3StorageProvider := storage.NewS3StorageProvider(awsSession)

	// init usecases
	createEvent := application.NewCreateEvent(eventsRepo)
	findEvents := application.NewFindEvents(eventsRepo, *geocodeProvider)
	updateEvents := application.NewUpdateEvent(eventsRepo)

	// init handlers
	createEventHandler := handlers.NewCreateEventHandler(createEvent, rabbitMQ)
	findAccountEventsHandler := handlers.NewFindAccountEventsHandler(findEvents)
	findEventsHandler := handlers.NewFindEventsHandler(findEvents)
	findNearbyEvents := handlers.NewFindNearbyEventsHandler(findEvents)
	findAllEventsHandler := handlers.NewFindAllEventsHandler(findEvents)
	updateEventsHandler := handlers.NewUpdateEventHandler(updateEvents, findEvents, rabbitMQ)
	cancelEventsHandler := handlers.NewCancelEventHandler(updateEvents, findEvents, rabbitMQ)
	coverUploadHandler := handlers.NewCoverUploadHandler(findEvents, updateEvents, s3StorageProvider)

	// init middlewares
	authMiddleware := middlewares.NewAuthMiddleware(keycloack)

	// setup middlewares
	router.Use(middleware.CORS())
	router.Use(middleware.RequestID())
	router.Use(middleware.Secure())
	router.Use(middleware.BodyLimit("2M"))
	router.Use(middleware.Logger())
	router.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(10)))

	events := router.Group("/events", authMiddleware.Auth)
	events.POST("", createEventHandler.Create)
	events.GET("/:id", findEventsHandler.Handle)
	events.GET("/nearby", findNearbyEvents.Handle)
	events.GET("/", findAllEventsHandler.Handle)
	events.PUT("/:id", updateEventsHandler.Handle)
	events.PUT("/:id/cancel", cancelEventsHandler.Handle)
	events.PUT("/:id/cover", coverUploadHandler.Handle)

	ticketOpts := events.Group("/:event_id/ticket_options")
	ticketOpts.POST("", func(c echo.Context) error { return c.NoContent(http.StatusOK) })
	ticketOpts.DELETE("/:ticket_id", func(c echo.Context) error { return c.NoContent(http.StatusOK) })
	ticketOpts.PUT("/:ticket_id", func(c echo.Context) error { return c.NoContent(http.StatusOK) })
	ticketOpts.GET("", func(c echo.Context) error { return c.NoContent(http.StatusOK) })
	ticketOpts.GET("/:ticket_id", func(c echo.Context) error { return c.NoContent(http.StatusOK) })

	lots := ticketOpts.Group("/:ticket_option_id/lots")
	lots.POST("", func(c echo.Context) error { return c.NoContent(http.StatusOK) })
	lots.DELETE("/:lot_id", func(c echo.Context) error { return c.NoContent(http.StatusOK) })
	lots.PUT("/:lot_id", func(c echo.Context) error { return c.NoContent(http.StatusOK) })
	lots.GET("", func(c echo.Context) error { return c.NoContent(http.StatusOK) })
	lots.GET("/:lot_id", func(c echo.Context) error { return c.NoContent(http.StatusOK) })

	accounts := router.Group("/accounts")
	accounts.GET("/:account_id/events", findAccountEventsHandler.Handle, authMiddleware.Auth)
}

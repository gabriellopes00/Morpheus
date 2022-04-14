package api

import (
	"events/application"
	"events/framework/api/handlers"
	"events/framework/api/middlewares"
	"events/framework/auth"
	"events/framework/db/repositories"
	"events/framework/encrypter"
	"events/framework/queue"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/streadway/amqp"
	"gorm.io/gorm"
)

func SetupServer(router *echo.Echo, database *gorm.DB, amqpConn *amqp.Channel) {

	// init adapters
	eventsRepo := repositories.NewPgEventsRepository(database)
	rabbitMQ := queue.NewRabbitMQ(amqpConn)
	jwtEncrypter := encrypter.NewJwtEncrypter()
	keycloack := auth.NewKeycloackauthProvider(jwtEncrypter)

	// init usecases
	createEvent := application.NewCreateEvent(eventsRepo)
	findEvents := application.NewFindEvents(eventsRepo)
	updateEvents := application.NewUpdateEvent(eventsRepo)

	// init handlers
	createEventHandler := handlers.NewCreateEventHandler(createEvent, rabbitMQ)
	findAccountEventsHandler := handlers.NewFindAccountEventsHandler(findEvents)
	findEventsHandler := handlers.NewFindEventsHandler(findEvents)
	findAllEventsHandler := handlers.NewFindAllEventsHandler(findEvents)
	updateEventsHandler := handlers.NewUpdateEventHandler(updateEvents, rabbitMQ)
	cancelEventsHandler := handlers.NewCancelEventHandler(updateEvents, rabbitMQ)

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
	events.GET("/", findAllEventsHandler.Handle)
	events.PUT("/:id", updateEventsHandler.Handle)
	events.PUT("/:id/cancel", cancelEventsHandler.Handle)

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

package api

import (
	"database/sql"
	"events/application"
	"events/framework/api/handlers"
	"events/framework/api/middlewares"
	"events/framework/db/repositories"
	"events/framework/encrypter"
	"events/framework/queue"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/streadway/amqp"
)

func SetupServer(router *echo.Echo, database *sql.DB, amqpConn *amqp.Channel) {

	// init adapters
	eventsRepo := repositories.NewPgEventsRepository(database)
	rabbitMQ := queue.NewRabbitMQ(amqpConn)
	jwtEncrypter := encrypter.NewJwtEncrypter()

	// init usecases
	createEvent := application.NewCreateEventUsecase(eventsRepo)
	getEvents := application.NewGetEvents(eventsRepo)

	// init handlers
	createEventHandler := handlers.NewCreateEventHandler(createEvent, rabbitMQ)
	getEventsHandler := handlers.NewGetEventsHandler(getEvents)

	// init middlewares
	authMiddleware := middlewares.NewAuthMiddleware(jwtEncrypter)

	// setup middlewares
	router.Use(middleware.CORS())
	router.Use(middleware.RequestID())
	router.Use(middleware.Secure())
	router.Use(middleware.BodyLimit("2M"))
	router.Use(middleware.Logger())
	router.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(10)))

	// setup routes
	router.GET("/ping", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"status": "Ok"})
	})
	router.GET("/error", func(c echo.Context) error {
		return c.JSON(http.StatusInternalServerError, map[string]string{"status": "Err"})
	})

	events := router.Group("/events")
	events.POST("", createEventHandler.Create, authMiddleware.Auth)

	accounts := router.Group("/accounts")
	accounts.GET("/:account_id/events", getEventsHandler.Handle, authMiddleware.Auth)
}

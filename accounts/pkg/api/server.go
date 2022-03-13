package api

import (
	"accounts/application"
	"accounts/pkg/api/handlers"
	"accounts/pkg/api/middlewares"
	"accounts/pkg/auth"
	"accounts/pkg/cache"
	"accounts/pkg/db"
	"accounts/pkg/encrypter"
	"accounts/pkg/queue"
	"accounts/pkg/storage"
	"database/sql"
	"net/http"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/go-redis/redis/v8"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/streadway/amqp"
)

func SetupServer(router *echo.Echo, database *sql.DB, rabbitmq *amqp.Channel, redis *redis.Client, awsSession *session.Session) {

	// init adapters
	accountRepo := db.NewPgAccountRepository(database)
	rabbitMQ := queue.NewRabbitMQ(rabbitmq)
	jwtEncrypter := encrypter.NewEncrypter()
	keycloack := auth.NewKeycloackauthProvider(jwtEncrypter)
	redisCache := cache.NewRedisCacheRepository(redis)
	s3Storage := storage.NewS3StorageProvider(awsSession)

	// init usecases
	createAccount := application.NewCreateAccount(accountRepo, keycloack)
	deleteAccount := application.NewDeleteAccount(accountRepo, redisCache)
	findAccount := application.NewFindAccount(accountRepo, redisCache)
	updateAccount := application.NewUpdateAccount(accountRepo, redisCache)

	// init handlers
	createAccountHandler := handlers.NewCreateAccountHandler(*createAccount, rabbitMQ, keycloack)
	avatarUploadHandler := handlers.NewAvatarUploadHandler(*updateAccount, s3Storage)
	authHandler := handlers.NewAuthHandler(keycloack)
	refreshAuthHandler := handlers.NewRefreshAuthHandler(keycloack)
	deleteAccountHandler := handlers.NewDeleteAccountHandler(*deleteAccount, rabbitMQ)
	getAccountHandler := handlers.NewFindAccountHandler(*findAccount)
	updateAccountHandler := handlers.NewUpdateAccountHandler(*updateAccount, rabbitMQ)

	// init middlewares
	authMiddleware := middlewares.NewAuthMiddleware(keycloack)

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

	r := router.Group("/accounts")

	r.POST("", createAccountHandler.Handle)
	r.POST("/auth", authHandler.Handle)
	r.POST("/auth/refresh", refreshAuthHandler.Handle)
	r.PUT("/:id/avatar/upload", avatarUploadHandler.Handle, authMiddleware.Auth)
	r.GET("/:id", getAccountHandler.Handle, authMiddleware.Auth)
	r.DELETE("/:id", deleteAccountHandler.Handle, authMiddleware.Auth)
	r.PUT("/:id", updateAccountHandler.Handle, authMiddleware.Auth)
}

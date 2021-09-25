package middlewares

import (
	"accounts/framework/encrypter"
	"errors"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

var (
	ErrInternalServer = errors.New("internal server error")
	ErrInvalidToken   = errors.New("invalid authentication token")
	ErrMissingToken   = errors.New("missing authentication token")
)

type authMiddleware struct {
	Encrypter encrypter.Encrypter
}

func NewAuthMiddleware(encrypter encrypter.Encrypter) *authMiddleware {
	return &authMiddleware{
		Encrypter: encrypter,
	}
}

func (m *authMiddleware) Auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		bearerToken := c.Request().Header.Get("Authorization")
		authorization := strings.Split(bearerToken, " ")
		if len(authorization) != 2 {
			return c.JSON(
				http.StatusUnauthorized,
				map[string]string{"error": ErrMissingToken.Error()})
		}

		token := authorization[1]

		accountId, err := m.Encrypter.DecryptAuthToken(token)
		if err != nil {
			if errors.Is(err, encrypter.ErrInvalidToken) {
				return c.JSON(
					http.StatusUnauthorized,
					map[string]string{"error": ErrInvalidToken.Error()})
			}

			return c.JSON(
				http.StatusInternalServerError,
				map[string]string{"error": ErrInternalServer.Error()})
		}

		c.Request().Header.Set("account_id", accountId)

		return next(c)
	}
}

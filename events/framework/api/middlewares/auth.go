package middlewares

import (
	"errors"
	"events/framework/encrypter"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

var (
	ErrMissingToken = errors.New("missing authentication token")
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

		accountId, err := m.Encrypter.Decrypt(token)
		if err != nil {
			return c.JSON(
				http.StatusUnauthorized,
				map[string]string{"error": err.Error()})
		}

		c.Request().Header.Set("account_id", accountId)

		return next(c)
	}
}

package middlewares

import (
	"accounts/framework/encrypter"
	"accounts/interfaces"
	"errors"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

var (
	ErrInternalServer    = errors.New("internal server error")
	ErrInvalidToken      = errors.New("invalid authentication token")
	ErrMissingToken      = errors.New("missing authentication token")
	ErrForbiddenDeletion = errors.New("forbidden account deletion")
)

type authMiddleware struct {
	Encrypter interfaces.Encrypter
}

func NewAuthMiddleware(encrypter interfaces.Encrypter) *authMiddleware {
	return &authMiddleware{
		Encrypter: encrypter,
	}
}

func (m *authMiddleware) Auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		accountId := c.Param("id")

		bearerToken := c.Request().Header.Get("Authorization")
		authorization := strings.Split(bearerToken, " ")
		if len(authorization) != 2 {
			return c.JSON(
				http.StatusUnauthorized,
				map[string]string{"error": ErrMissingToken.Error()})
		}

		token := authorization[1]

		account, err := m.Encrypter.Decrypt(token)
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

		if accountId != account.Id {
			return c.JSON(
				http.StatusForbidden,
				map[string]string{"error": ErrForbiddenDeletion.Error()})
		}

		c.Request().Header.Set("account_id", account.Id)

		return next(c)
	}
}

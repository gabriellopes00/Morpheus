package middlewares

import (
	"accounts/framework/encrypter"
	"accounts/interfaces"
	"errors"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
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

		var token string

		bearerToken := c.Request().Header.Get("Authorization")
		authorization := strings.Split(bearerToken, " ")
		if len(authorization) != 2 {
			return c.JSON(http.StatusUnauthorized, nil)
		}

		token = authorization[1]

		account, err := m.Encrypter.Decrypt(token)
		if err != nil {
			if errors.Is(err, encrypter.ErrInvalidToken) {
				return c.JSON(http.StatusUnauthorized, err)
			}

			return c.JSON(http.StatusInternalServerError, err)
		}

		c.Request().Header.Set("account_id", account.Id)

		return next(c)
	}
}

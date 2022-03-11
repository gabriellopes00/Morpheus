package middlewares

import (
	"accounts/pkg/auth"
	"accounts/pkg/encrypter"
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
	AuthProvider auth.AuthProvider
}

func NewAuthMiddleware(AuthProvider auth.AuthProvider) *authMiddleware {
	return &authMiddleware{
		AuthProvider: AuthProvider,
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

		accountInfo, err := m.AuthProvider.AuthUser(token)
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

		c.Request().Header.Set("account_id", accountInfo.Id)

		return next(c)
	}
}

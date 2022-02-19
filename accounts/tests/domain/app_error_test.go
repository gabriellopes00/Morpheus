package domain_test

import (
	domain_error "accounts/domain/errors"
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsAppError(t *testing.T) {
	t.Parallel()

	t.Run("Valid App Error", func(t *testing.T) {
		appError := domain_error.NewAppError("App Error", "lorem ipsum dolor sit amet")

		isValid := domain_error.IsAppError(*appError)
		require.True(t, isValid)

		isValid = domain_error.IsAppError(appError)
		require.True(t, isValid)
	})

	t.Run("Invalid App Error", func(t *testing.T) {
		appError := errors.New("any error")

		isValid := domain_error.IsAppError(appError)
		require.False(t, isValid)
	})

	t.Run("App Error Message", func(t *testing.T) {
		appError := domain_error.NewAppError("App Error", "lorem ipsum dolor sit amet")

		require.Equal(t, "App Error: lorem ipsum dolor sit amet", appError.Error())
	})

}

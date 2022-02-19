package domain_test

import (
	"accounts/domain/entities"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestNewAccount(t *testing.T) {
	t.Parallel()

	t.Run("Valid account input", func(t *testing.T) {
		_, err := entities.NewAccount(
			"John Doe",
			"johndoe@mai.com",
			"lorepipsum",
			"https://johndoe.jpg",
			time.Now().Local().String(),
			"asdf",
		)

		require.Nil(t, err)
	})
}

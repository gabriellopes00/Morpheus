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
			"https://johndoe.jpg",
			time.Now().Local().String(),
			"111.222.333-44",
		)

		require.Nil(t, err)
	})
}

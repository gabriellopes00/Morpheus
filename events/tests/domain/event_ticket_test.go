package domain_test

import (
	"events/domain/entities"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTicketOption(t *testing.T) {
	t.Parallel()

	_, err := entities.NewTicketOption("vip", "vip ticket", "2022-03-24T18:29:37.001Z", "2022-04-24T18:29:37.001Z", "", 6, 1, nil)
	assert.Nil(t, err)

	_, err = entities.NewTicketOption("vip", "", "2022-03-24T18:29:37.001Z", "2022-04-24T18:29:37.001Z", "", 6, 1, nil)
	assert.Nil(t, err)

	_, err = entities.NewTicketOption("", "vip ticket", "2022-03-24T18:29:37.001Z", "2022-04-24T18:29:37.001Z", "", 6, 1, nil)
	assert.Error(t, err)

	_, err = entities.NewTicketOption("vip", "vip ticket", "2022-03-24T18:29:37.001Z", "2022-04-24T18:29:37.001Z", "", 6, 0, nil)
	assert.Error(t, err)

	_, err = entities.NewTicketOption("vip", "vip ticket", "2022-03-24T18:29:37.001Z", "2022-04-24T18:29:37.001Z", "", 1, 3, nil)
	assert.Error(t, err)

	_, err = entities.NewTicketOption("vip", "vip ticket", "2022-04-24T18:29:37.001Z", "2022-03-24T18:29:37.001Z", "", 1, 3, nil)
	assert.Error(t, err)

	_, err = entities.NewTicketOption("vip", "vip ticket", "2022-03-24T18:29:37.001Z", "invalid", "", 1, 3, nil)
	assert.Error(t, err)

	_, err = entities.NewTicketOption("vip", "vip ticket", "invalid", "2022-03-24T18:29:37.001Z", "", 1, 3, nil)
	assert.Error(t, err)

}

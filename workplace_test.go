package doppler

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	update_workplace_name = "Test"
	update_billing_email  = "example@web.com"
)

func TestGetWorkplace(t *testing.T) {
	dp, err := NewFromEnv()
	assert.Nil(t, err)

	workplace, err := dp.GetWorkplace()
	assert.Nil(t, err)
	assert.NotNil(t, workplace)
}

func TestUpdateWorkplace(t *testing.T) {
	dp, err := NewFromEnv()
	assert.Nil(t, err)

	workplace, err := dp.UpdateWorkplace(WorkplaceParams{Name: update_workplace_name, BillingEmail: update_billing_email})
	assert.Nil(t, err)
	assert.NotNil(t, workplace)

	assert.Equal(t, update_billing_email, workplace.Workplace.BillingEmail)
	assert.Equal(t, update_workplace_name, workplace.Workplace.Name)
}

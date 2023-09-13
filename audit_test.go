package doppler

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_doppler_GetWorkplaceUsers(t *testing.T) {
	dp, err := New("")
	assert.Nil(t, err)

	users, err := dp.GetWorkplaceUsers(true, 1)
	assert.Nil(t, err)
	assert.NotNil(t, users)
}

func Test_doppler_GetWorkplaceUser(t *testing.T) {
	dp, err := New("")
	assert.Nil(t, err)

	user, err := dp.GetWorkplaceUser("4adac93e-4ae9-4c41-a23e-a4fc689d26cc", true)
	assert.Nil(t, err)
	assert.NotNil(t, user)
}

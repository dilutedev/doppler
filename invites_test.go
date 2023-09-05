package doppler

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListInvites(t *testing.T) {
	dp, err := NewFromEnv()
	assert.Nil(t, err)
	assert.NotNil(t, dp)

	// test without page and limit params
	invites, err := dp.ListInvites(nil, nil)
	assert.Nil(t, err)
	assert.NotNil(t, invites)
	assert.IsType(t, &Invites{}, invites)

	// test with page and limit params
	invites, err = dp.ListInvites(&test_page, &test_limit)
	assert.Nil(t, err)
	assert.NotNil(t, invites)
	assert.IsType(t, &Invites{}, invites)
}

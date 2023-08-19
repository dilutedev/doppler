package doppler

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	doppler, err := NewFromEnv()
	assert.Nil(t, err)

	assert.NotNil(t, doppler)
}

package doppler

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	test_page   int = 1
	test_limit  int = 10
	test_log_id     = "GBIoiNPme2C1vsg6GQ5n3Msz"
)

func TestRetrieveLogs(t *testing.T) {
	dp, err := NewFromEnv()
	assert.Nil(t, err)
	assert.NotNil(t, dp)

	// test with limit param
	logs, err := dp.RetrieveLogs(test_page, &test_limit)
	assert.Nil(t, err)
	assert.NotNil(t, logs)
	assert.Equal(t, test_page, logs.Page)
	assert.LessOrEqual(t, len(logs.Logs), test_limit)

	// test without limit param
	logs, err = dp.RetrieveLogs(test_page, nil)
	assert.Nil(t, err)
	assert.NotNil(t, logs)
	assert.Equal(t, test_page, logs.Page)
}

func Test_RetrieveLog(t *testing.T) {
	dp, err := NewFromEnv()
	assert.Nil(t, err)
	assert.NotNil(t, dp)

	log, err := dp.RetrieveLog(test_log_id)
	assert.Nil(t, err)
	assert.NotNil(t, log)

	assert.Equal(t, log.Log.ID, test_log_id)
}

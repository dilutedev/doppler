package doppler

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	configLogID = "dT41rFsPw4yhC9LIcnpVYo0z"
)

func TestListConfigLogs(t *testing.T) {
	dp, err := NewFromEnv()
	assert.Nil(t, err)
	assert.NotNil(t, dp)

	// test with limit param
	logs, err := dp.ListConfigLogs(projectName, config, test_page, &test_limit)
	assert.Nil(t, err)
	assert.NotNil(t, logs)
	assert.IsType(t, &ConfigLogs{}, logs)
	assert.Equal(t, test_page, logs.Page)
	assert.LessOrEqual(t, len(logs.Logs), test_limit)

	// test without limit param
	logs, err = dp.ListConfigLogs(projectName, config, test_page, nil)
	assert.Nil(t, err)
	assert.IsType(t, &ConfigLogs{}, logs)
	assert.NotNil(t, logs)
	assert.Equal(t, test_page, logs.Page)
}

func TestRetrieveConfigLog(t *testing.T) {
	dp, err := NewFromEnv()
	assert.Nil(t, err)
	assert.NotNil(t, dp)

	iConfigLog, err := dp.RetrieveConfigLog(projectName, config, configLogID)
	assert.Nil(t, err)
	assert.NotNil(t, iConfigLog)
	// assert.IsType(t, &IConfigLog{}, iConfigLog)
}

func TestRollbackConfigLog(t *testing.T) {
	dp, err := NewFromEnv()
	assert.Nil(t, err)
	assert.NotNil(t, dp)

	iConfigLog, err := dp.RollbackConfigLog(projectName, config, configLogID)
	assert.Nil(t, err)
	assert.NotNil(t, iConfigLog)
	assert.IsType(t, &IConfigLog{}, iConfigLog)
}

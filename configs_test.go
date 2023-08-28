package doppler

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	config = environmentSlug + "_test"
)

func TestListConfigs(t *testing.T) {
	dp, err := NewFromEnv()
	assert.Nil(t, err)
	assert.NotNil(t, dp)

	// test with limit param
	configs, err := dp.ListConfigs(projectName, environmentSlug, test_page, &test_limit)
	assert.Nil(t, err)
	assert.NotNil(t, configs)
	assert.IsType(t, &Configs{}, configs)
	assert.Equal(t, test_page, configs.Page)
	assert.LessOrEqual(t, len(configs.Configs), test_limit)

	// test without limit param
	configs, err = dp.ListConfigs(projectName, environmentSlug, test_page, nil)
	assert.Nil(t, err)
	assert.IsType(t, &Configs{}, configs)
	assert.NotNil(t, configs)
	assert.Equal(t, test_page, configs.Page)

}

func TestCreateConfig(t *testing.T) {
	dp, err := NewFromEnv()
	assert.Nil(t, err)
	assert.NotNil(t, dp)

	iConfig, err := dp.CreateConfig(CreateConfigParams{Project: projectName, Environment: environmentSlug, Name: config})
	assert.Nil(t, err)
	assert.NotNil(t, iConfig)
	assert.IsType(t, &IConfig{}, iConfig)
	// creating another config that will be deleted
	iConfig, err = dp.CreateConfig(CreateConfigParams{Project: projectName, Environment: environmentSlug, Name: environmentSlug + "_test2"})

}

func TestRetrieveConfig(t *testing.T) {
	dp, err := NewFromEnv()
	assert.Nil(t, err)
	assert.NotNil(t, dp)

	iConfig, err := dp.RetrieveConfig(projectName, config)
	assert.Nil(t, err)
	assert.NotNil(t, iConfig)
	assert.IsType(t, &IConfig{}, iConfig)
}

func TestUpdateConfig(t *testing.T) {
	dp, err := NewFromEnv()
	assert.Nil(t, err)
	assert.NotNil(t, dp)

	iConfig, err := dp.UpdateConfig(ModifyConfigParams{Project: projectName, Config: environmentSlug + "_test2", Name: environmentSlug + "_itest2"})
	assert.Nil(t, err)
	assert.NotNil(t, iConfig)
	assert.IsType(t, &IConfig{}, iConfig)
}
func TestDeleteConfig(t *testing.T) {
	dp, err := NewFromEnv()
	assert.Nil(t, err)
	assert.NotNil(t, dp)

	success, err := dp.DeleteConfig(DeletConfigParams{Project: projectName, Config: environmentSlug + "_itest2"})
	assert.Nil(t, err)
	assert.NotNil(t, success)
	assert.IsType(t, &Success{}, success)
}
func TestCloneConfig(t *testing.T) {
	dp, err := NewFromEnv()
	assert.Nil(t, err)
	assert.NotNil(t, dp)

	iConfig, err := dp.CloneConfig(ModifyConfigParams{Project: projectName, Config: config, Name: config + "clone"})
	assert.Nil(t, err)
	assert.NotNil(t, iConfig)
	assert.IsType(t, &IConfig{}, iConfig)
}
func TestLockConfig(t *testing.T) {
	dp, err := NewFromEnv()
	assert.Nil(t, err)
	assert.NotNil(t, dp)

	iConfig, err := dp.LockConfig(DeletConfigParams{Project: projectName, Config: config})
	assert.Nil(t, err)
	assert.NotNil(t, iConfig)
	assert.IsType(t, &IConfig{}, iConfig)
}
func TestUnlockConfig(t *testing.T) {
	dp, err := NewFromEnv()
	assert.Nil(t, err)
	assert.NotNil(t, dp)

	iConfig, err := dp.UnlockConfig(DeletConfigParams{Project: projectName, Config: config})
	assert.Nil(t, err)
	assert.NotNil(t, iConfig)
	assert.IsType(t, &IConfig{}, iConfig)
}

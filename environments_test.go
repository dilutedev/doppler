package doppler

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	projectName     = "test"
	environmentSlug = "dev"
)

func TestListEnvironments(t *testing.T) {
	dp, err := NewFromEnv()
	assert.Nil(t, err)
	assert.NotNil(t, dp)

	environments, err := dp.ListEnvironments(projectName)
	assert.Nil(t, err)
	assert.NotNil(t, environments)

}

func TestRetrieveEnvironment(t *testing.T) {
	dp, err := NewFromEnv()
	assert.Nil(t, err)
	assert.NotNil(t, dp)

	environment, err := dp.RetrieveEnvironment(projectName, environmentSlug)
	assert.Nil(t, err)
	assert.NotNil(t, environment)
}

func TestCreateEnvironment(t *testing.T) {
	dp, err := NewFromEnv()
	assert.Nil(t, err)
	assert.NotNil(t, dp)

	environment, err := dp.CreateEnvironment(projectName, EnvironmentBodyParams{
		Name: "Continuous Integration",
		Slug: "ci",
	})
	assert.Nil(t, err)
	assert.NotNil(t, environment)
}

func TestDeleteEnvironment(t *testing.T) {
	dp, err := NewFromEnv()
	assert.Nil(t, err)
	assert.NotNil(t, dp)

	testSlug := "ci"
	data, err := dp.DeleteEnvironment(projectName, testSlug)
	assert.Nil(t, err)
	assert.NotNil(t, data)
}

func TestRenameEnvironment(t *testing.T) {
	dp, err := NewFromEnv()
	assert.Nil(t, err)
	assert.NotNil(t, dp)

	environments, err := dp.RenameEnvironment(projectName, environmentSlug, EnvironmentBodyParams{
		Name: "Development env",
		Slug: "devn",
	})
	assert.Nil(t, err)
	assert.NotNil(t, environments)
}

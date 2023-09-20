package doppler

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	testToken  = "test_token"
	testExpiry = "1718924458" // Thursday, June 20, 2024 11:00:58 PM
	tokenSlug  = ""
)

func TestListServiceTokens(t *testing.T) {
	dp, err := NewFromEnv()
	assert.Nil(t, err)
	assert.NotNil(t, dp)

	serviceTokens, err := dp.ListServiceTokens(projectName, config)
	assert.Nil(t, err)
	assert.NotNil(t, serviceTokens)
	assert.IsType(t, &ServiceTokens{}, serviceTokens)
}

func TestCreateServiceToken(t *testing.T) {
	dp, err := NewFromEnv()
	assert.Nil(t, err)
	assert.NotNil(t, dp)

	serviceToken, err := dp.CreateServiceToken(CreateTokenParams{
		Name:     testToken,
		Project:  projectName,
		Config:   config,
		ExpireAt: testExpiry,
		Access:   "read",
	})
	assert.Nil(t, err)
	assert.NotNil(t, serviceToken)
	assert.IsType(t, &ServiceTokenModel{}, serviceToken)
	tokenSlug = serviceToken.Token.Slug
}

func TestDeleteServiceToken(t *testing.T) {
	dp, err := NewFromEnv()
	assert.Nil(t, err)
	assert.NotNil(t, dp)

	response, err := dp.DeleteServiceToken(DeleteTokenParams{
		Project: projectName,
		Config:  config,
		Slug:    tokenSlug,
	})
	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.IsType(t, &Success{}, response)
}

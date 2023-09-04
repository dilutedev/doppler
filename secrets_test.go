package doppler

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	secret = "TEST_SECRET"
)

func TestListSecrets(t *testing.T) {
	dp, err := NewFromEnv()
	assert.Nil(t, err)
	assert.NotNil(t, dp)

	// test including dynamic secrets and managed secrets
	secrets, err := dp.ListSecrets(ListSecretsParams{
		Project:               projectName,
		Config:                config,
		IncludeDynamicSecrets: true,
		DynamicSecretsTTLSec:  60,
		Secrets:               "TEST_SECRET",
		IncludeManagedSecrets: true,
	})
	assert.Nil(t, err)
	assert.NotNil(t, secrets)
	assert.IsType(t, &Secrets{}, secrets)

	// test without dynamic secrets and managed secrets
	secrets, err = dp.ListSecrets(ListSecretsParams{
		Project:               projectName,
		Config:                config,
		IncludeDynamicSecrets: false,
		Secrets:               "TEST_SECRET",
		IncludeManagedSecrets: false,
	})
	assert.Nil(t, err)
	assert.NotNil(t, secrets)
	assert.IsType(t, &Secrets{}, secrets)
}
func TestListSecretNames(t *testing.T) {
	dp, err := NewFromEnv()
	assert.Nil(t, err)
	assert.NotNil(t, dp)

	// test including dynamic secrets and managed secrets
	secretNames, err := dp.ListSecretNames(ListSecretNamesParams{
		Project:               projectName,
		Config:                config,
		IncludeDynamicSecrets: true,
		IncludeManagedSecrets: true,
	})
	assert.Nil(t, err)
	assert.NotNil(t, secretNames)
	assert.IsType(t, &SecretNames{}, secretNames)

	// test without dynamic secrets and managed secrets
	secretNames, err = dp.ListSecretNames(ListSecretNamesParams{
		Project:               projectName,
		Config:                config,
		IncludeDynamicSecrets: false,
		IncludeManagedSecrets: false,
	})
	assert.Nil(t, err)
	assert.NotNil(t, secretNames)
	assert.IsType(t, &SecretNames{}, secretNames)
}
func TestRetrieveSecret(t *testing.T) {
	dp, err := NewFromEnv()
	assert.Nil(t, err)
	assert.NotNil(t, dp)

	secret, err := dp.RetrieveSecret(projectName, config, secret)
	assert.Nil(t, err)
	assert.NotNil(t, secret)
	assert.IsType(t, &Secret{}, secret)
}
func TestDeleteSecret(t *testing.T) {
	dp, err := NewFromEnv()
	assert.Nil(t, err)
	assert.NotNil(t, dp)

	response, err := dp.DeleteSecret(projectName, config, secret)
	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, "", response)
}
func TestUpdateSecret(t *testing.T) {
	dp, err := NewFromEnv()
	assert.Nil(t, err)
	assert.NotNil(t, dp)

	secrets, err := dp.UpdateSecret(UpdateSecretParams{
		Project: projectName,
		Config:  config,
		Secrets: map[string]string{
			"TEST_SECRET": "test123",
		},
	},
	)
	assert.Nil(t, err)
	assert.NotNil(t, secrets)
	assert.IsType(t, &Secrets{}, secrets)
}
func TestDownloadSecret(t *testing.T) {
	dp, err := NewFromEnv()
	
	assert.Nil(t, err)
	assert.NotNil(t, dp)

	// test including dynamic secrets
	secrets, err := dp.DownloadSecret(DownloadSecretParams{
		Project:               projectName,
		Config:                config,
		IncludeDynamicSecrets: true,
		DynamicSecretsTTLSec:  60,
		Format:                "json",
		NameTransformer:       "lower-kebab",
	})
	assert.Nil(t, err)
	assert.NotNil(t, secrets)
	assert.Contains(t, secrets, projectName)

	// test without dynamic secrets
	secrets, err = dp.DownloadSecret(DownloadSecretParams{
		Project:               projectName,
		Config:                config,
		IncludeDynamicSecrets: false,
		Format:                "json",
		NameTransformer:       "lower-kebab",
	})
	assert.Nil(t, err)
	assert.NotNil(t, secrets)
	assert.Contains(t, secrets, projectName)
}
func TestUpdateNote(t *testing.T) {
	dp, err := NewFromEnv()
	assert.Nil(t, err)
	assert.NotNil(t, dp)

	response, err := dp.UpdateNote(SetNoteParams{
		Project: projectName,
		Config:  config,
		Secret:  secret,
		Note:    "test note",
	})
	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.IsType(t, &NoteResponse{}, response)
}

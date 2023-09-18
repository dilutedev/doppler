package doppler

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	serviceAccountSlug    = "78bf254d-889d-4c43-b68a-034fa67230ef"
	newServiceAccountSlug = ""
)

func TestListServiceAccounts(t *testing.T) {
	dp, err := NewFromEnv()
	assert.Nil(t, err)
	assert.NotNil(t, dp)

	serviceAccounts, err := dp.ListServiceAccounts(&test_page, &test_limit)
	assert.Nil(t, err)
	assert.NotNil(t, serviceAccounts)
	assert.IsType(t, &ServiceAccounts{}, serviceAccounts)
}

func TestCreateServiceAccount(t *testing.T) {
	dp, err := NewFromEnv()
	assert.Nil(t, err)
	assert.NotNil(t, dp)

	serviceAccount, err := dp.CreateServiceAccount(ServiceAccountBodyParams{
		Name: "testSlug",
		WorkplaceRole: WorkplaceRoleObject{
			Identifier: "admin",
		},
	})
	assert.Nil(t, err)
	assert.NotNil(t, serviceAccount)
	assert.IsType(t, &ServiceAccountModel{}, serviceAccount)
	newServiceAccountSlug = serviceAccount.ServiceAccount.Slug
}
func TestRetrieveServiceAccount(t *testing.T) {
	dp, err := NewFromEnv()
	assert.Nil(t, err)
	assert.NotNil(t, dp)

	serviceAccount, err := dp.RetrieveServiceAccount(serviceAccountSlug)
	assert.Nil(t, err)
	assert.NotNil(t, serviceAccount)
	assert.IsType(t, &ServiceAccountModel{}, serviceAccount)
}

func TestUpdateServiceAccount(t *testing.T) {
	dp, err := NewFromEnv()
	assert.Nil(t, err)
	assert.NotNil(t, dp)

	serviceAccount, err := dp.UpdateServiceAccount(newServiceAccountSlug, ServiceAccountBodyParams{
		Name: "testSlugUpdate",
		WorkplaceRole: WorkplaceRoleObject{
			Identifier: "collaborator",
		},
	})
	assert.Nil(t, err)
	assert.NotNil(t, serviceAccount)
	assert.IsType(t, &ServiceAccountModel{}, serviceAccount)
}

func TestDeleteServiceAccount(t *testing.T) {
	dp, err := NewFromEnv()
	assert.Nil(t, err)
	assert.NotNil(t, dp)

	serviceAccount, err := dp.DeleteServiceAccount(newServiceAccountSlug)
	assert.Nil(t, err)
	assert.NotNil(t, serviceAccount)
}

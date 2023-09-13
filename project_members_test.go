package doppler

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_doppler_ListProjectMembers(t *testing.T) {
	dp, err := NewFromEnv()
	assert.Nil(t, err)

	pm, err := dp.ListProjectMembers("dopplersdk", 1, 20)
	assert.Nil(t, err)
	assert.NotNil(t, pm)
}

func Test_doppler_RetrieveProjectMember(t *testing.T) {
	dp, err := NewFromEnv()
	assert.Nil(t, err)

	pm, err := dp.RetrieveProjectMember(MemberInvite, "a5aa8014-26a4-4a90-bcfc-758b2ce3c409", "dopplersdk")
	assert.Nil(t, err)
	assert.NotNil(t, pm)
}

func Test_doppler_AddProjectMember(t *testing.T) {
	dp, err := NewFromEnv()
	assert.Nil(t, err)

	pm, err := dp.AddProjectMember("dopplersdk", AddProjectMemberParam{
		Type:         "invite",
		Slug:         "a5aa8014-26a4-4a90-bcfc-758b2ce3c409",
		Role:         "admin",
		Environments: []string{},
	})
	assert.Nil(t, err)
	assert.NotNil(t, pm)
}

func Test_doppler_UpdateProjectMember(t *testing.T) {
	dp, err := NewFromEnv()
	assert.Nil(t, err)

	pm, err := dp.UpdateProjectMember("dopplersdk", "a5aa8014-26a4-4a90-bcfc-758b2ce3c409", MemberInvite, UpdateProjectMemberParams{
		Role:         "admin",
		Environments: []string{},
	})
	assert.Nil(t, err)
	assert.NotNil(t, pm)
}

func Test_doppler_RemoveProjectMember(t *testing.T) {
	dp, err := NewFromEnv()
	assert.Nil(t, err)

	err = dp.RemoveProjectMember(MemberWorkplaceUser, "", "")
	assert.Nil(t, err)
}

package doppler

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListstatuss(t *testing.T) {
	dp, err := NewFromEnv()
	assert.Nil(t, err)
	assert.NotNil(t, dp)

	// test with limit
	projects, err := dp.ListProjects(test_page, &test_limit)
	assert.Nil(t, err)
	assert.NotNil(t, projects)
	assert.Equal(t, projects.Page, test_page)
	assert.LessOrEqual(t, len(projects.Projects), test_limit)

	// test without limit
	projects, err = dp.ListProjects(test_page, nil)
	assert.Nil(t, err)
	assert.NotNil(t, projects)
	assert.Equal(t, projects.Page, test_page)
}

func TestCreateProject(t *testing.T) {
	dp, err := NewFromEnv()
	assert.Nil(t, err)
	assert.NotNil(t, dp)

	project, err := dp.CreateProject(CreateProjectParams{
		Name:        "Test",
		Description: "test description",
	})
	assert.Nil(t, err)
	assert.NotNil(t, project)
}

func TestRetrieveProject(t *testing.T) {
	dp, err := NewFromEnv()
	assert.Nil(t, err)
	assert.NotNil(t, dp)

	name := "test"
	project, err := dp.RetrieveProject(name)

	assert.Nil(t, err)
	assert.NotNil(t, project)
	assert.Equal(t, name, project.Project.ID)

}

func TestUpdateProject(t *testing.T) {
	dp, err := NewFromEnv()
	assert.Nil(t, err)
	assert.NotNil(t, dp)

	project, err := dp.UpdateProject(UpdateProjectParams{
		Name:        "new-name",
		Description: "new-description",
		ProjectID:   "test",
	})
	assert.Nil(t, err)
	assert.NotNil(t, project)
}

func TestDeleteProject(t *testing.T) {
	dp, err := NewFromEnv()
	assert.Nil(t, err)
	assert.NotNil(t, dp)

	status, err := dp.DeleteProject("test")
	assert.Nil(t, err)
	assert.NotNil(t, status)
}

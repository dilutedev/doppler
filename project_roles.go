package doppler

import (
	"encoding/json"
	"net/http"
)

type ProjectRole struct {
	Name         string   `json:"name,omitempty"`
	Permissions  []string `json:"permissions,omitempty"`
	Identifier   string   `json:"identifier,omitempty"`
	CreatedAt    string   `json:"created_at,omitempty"`
	IsCustomRole bool     `json:"is_custom_role,omitempty"`
}

type ProjectRoles struct {
	ProjectRoles []ProjectRole `json:"roles,omitempty"`
	Success      bool        `json:"success,omitempty"`
}

type ProjectPermissions struct {
	Permissions []string `json:"permissions,omitempty"`
	Success     bool     `json:"success,omitempty"`
}

type RetrieveProjectResponse struct {
	Role    ProjectRole `json:"role"`
	Success bool        `json:"success"`
}

func (dp *doppler) ListProjectRoles() (*ProjectRoles, error) {
	request, err := http.NewRequest(
		http.MethodGet,
		"/v3/projects/roles",
		nil,
	)
	if err != nil {
		return nil, err
	}

	body, err := dp.makeApiRequest(request)
	if err != nil {
		return nil, err
	}

	data := &ProjectRoles{}
	err = json.Unmarshal(body, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (dp *doppler) ListProjectPermissions() (*ProjectPermissions, error) {
	request, err := http.NewRequest(
		http.MethodGet,
		"/v3/projects/permissions",
		nil,
	)
	if err != nil {
		return nil, err
	}

	body, err := dp.makeApiRequest(request)
	if err != nil {
		return nil, err
	}

	data := &ProjectPermissions{}
	err = json.Unmarshal(body, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (dp *doppler) RetrieveProjectRole(role string) (*RetrieveProjectResponse, error) {
	request, err := http.NewRequest(
		http.MethodGet,
		"/v3/workplace/roles/role/"+role,
		nil,
	)
	if err != nil {
		return nil, err
	}

	body, err := dp.makeApiRequest(request)
	if err != nil {
		return nil, err
	}

	data := &RetrieveProjectResponse{}
	err = json.Unmarshal(body, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (dp *doppler) CreateProjectRole(role string) (*RetrieveProjectResponse, error) {
	// TODO: Add payload (missing in docs)
	request, err := http.NewRequest(
		http.MethodPost,
		"/v3/projects/roles",
		nil,
	)
	if err != nil {
		return nil, err
	}

	body, err := dp.makeApiRequest(request)
	if err != nil {
		return nil, err
	}

	data := &RetrieveProjectResponse{}
	err = json.Unmarshal(body, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (dp *doppler) UpdateProjectRole(role string) (*RetrieveProjectResponse, error) {
	// TODO: Add payload (missing in docs)
	request, err := http.NewRequest(
		http.MethodPatch,
		"/v3/projects/roles/role/"+role,
		nil,
	)
	if err != nil {
		return nil, err
	}

	body, err := dp.makeApiRequest(request)
	if err != nil {
		return nil, err
	}

	data := &RetrieveProjectResponse{}
	err = json.Unmarshal(body, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (dp *doppler) DeleteProjectRole(role string) (string, error) { // response is empty in docs
	// TODO: Add payload (missing in docs)
	request, err := http.NewRequest(
		http.MethodDelete,
		"/v3/projects/roles/role/"+role,
		nil,
	)
	if err != nil {
		return "", err
	}

	body, err := dp.makeApiRequest(request)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

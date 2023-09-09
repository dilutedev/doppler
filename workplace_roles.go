package doppler

import (
	"encoding/json"
	"net/http"
)

type WorkplaceRole struct {
	Name         string   `json:"name,omitempty"`
	Permissions  []string `json:"permissions,omitempty"`
	Identifier   string   `json:"identifier,omitempty"`
	CreatedAt    string   `json:"created_at,omitempty"`
	IsCustomRole bool     `json:"is_custom_role,omitempty"`
	IsInlineRole bool     `json:"is_inline_role,omitempty"`
}

type WorkplaceRoles struct {
	WorkplaceRoles []WorkplaceRole `json:"roles,omitempty"`
	Success        string          `json:"success,omitempty"`
}

type WorkplacePermissions struct {
	Permissions []string `json:"permissions,omitempty"`
	Success     bool     `json:"success,omitempty"`
}

type RetrieveWorkplaceResponse struct {
	Role    ProjectRole `json:"role"`
	Success bool        `json:"success"`
}

func (dp *doppler) ListWorkplaceRoles() (*WorkplaceRoles, error) {
	request, err := http.NewRequest(
		http.MethodGet,
		"/v3/workplace/roles",
		nil,
	)
	if err != nil {
		return nil, err
	}

	body, err := dp.makeApiRequest(request)
	if err != nil {
		return nil, err
	}

	data := &WorkplaceRoles{}
	err = json.Unmarshal(body, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (dp *doppler) ListWorkplacePermissions() (*WorkplacePermissions, error) {
	request, err := http.NewRequest(
		http.MethodGet,
		"/v3/workplace/permissions",
		nil,
	)
	if err != nil {
		return nil, err
	}

	body, err := dp.makeApiRequest(request)
	if err != nil {
		return nil, err
	}

	data := &WorkplacePermissions{}
	err = json.Unmarshal(body, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (dp *doppler) RetrieveWorkplaceRole(role string) (*RetrieveWorkplaceResponse, error) {
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

	data := &RetrieveWorkplaceResponse{}
	err = json.Unmarshal(body, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (dp *doppler) CreateWorkplaceRole(role string) (*WorkplaceRole, error) {
	// TODO: Add payload (missing in docs)
	request, err := http.NewRequest(
		http.MethodPost,
		"/v3/workplace/roles",
		nil,
	)
	if err != nil {
		return nil, err
	}

	body, err := dp.makeApiRequest(request)
	if err != nil {
		return nil, err
	}

	data := &WorkplaceRole{}
	err = json.Unmarshal(body, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (dp *doppler) UpdateWorkplaceRole(role string) (*WorkplaceRole, error) {
	// TODO: Add payload (missing in docs)
	request, err := http.NewRequest(
		http.MethodPatch,
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

	data := &WorkplaceRole{}
	err = json.Unmarshal(body, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (dp *doppler) DeleteWorkplaceRole(role string) (string, error) {
	request, err := http.NewRequest(
		http.MethodDelete,
		"/v3/workplace/roles/role/"+role,
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

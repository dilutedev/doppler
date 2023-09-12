package doppler

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
)

type ServiceAccount struct {
	Name          string        `json:"name,omitempty"`
	Slug          string        `json:"slug,omitempty"`
	CreatedAt     string        `json:"created_at,omitempty"`
	WorkplaceRole WorkplaceRole `json:"workplace_role,omitempty"`
}

type ServiceAccounts struct {
	ServiceAccounts []ServiceAccount `json:"roles,omitempty"`
	Success         string           `json:"success,omitempty"`
}

type ServiceAccountModel struct {
	ServiceAccount ServiceAccount `json:"service_account,omitempty"`
	Success        string         `json:"success,omitempty"`
}

type ServiceAccountBodyParams struct {
	Name          string              `json:"name,omitempty"`
	WorkplaceRole WorkplaceRoleObject `json:"workplace_role,omitempty"`
}

type WorkplaceRoleObject struct {
	Identifier  string   `json:"identifier,omitempty"`
	Permissions []string `json:"permissions,omitempty"`
}

func (dp *doppler) ListServiceAccounts(page, limit *int) (*ServiceAccounts, error) {
	defaultLimit := 20
	defaultPage := 1
	if page == nil || *page <= 0 {
		page = &defaultPage
	}
	if limit == nil || *limit <= 0 {
		limit = &defaultLimit
	}
	request, err := http.NewRequest(
		http.MethodGet,
		"/v3/workplace/service_accounts?page="+strconv.Itoa(*page)+"&per_page="+strconv.Itoa(*limit),
		nil,
	)
	if err != nil {
		return nil, err
	}

	body, err := dp.makeApiRequest(request)
	if err != nil {
		return nil, err
	}

	data := &ServiceAccounts{}
	err = json.Unmarshal(body, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (dp *doppler) RetrieveServiceAccount(slug string) (*ServiceAccountModel, error) {
	request, err := http.NewRequest(
		http.MethodGet,
		"/v3/workplace/service_accounts/service_account/"+slug,
		nil,
	)
	if err != nil {
		return nil, err
	}

	body, err := dp.makeApiRequest(request)
	if err != nil {
		return nil, err
	}

	data := &ServiceAccountModel{}
	err = json.Unmarshal(body, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (dp *doppler) CreateServiceAccount(params ServiceAccountBodyParams) (*ServiceAccountModel, error) {
	if params.WorkplaceRole.Identifier != "" || params.WorkplaceRole.Permissions != nil {
		return nil, errors.New("You may provide an identifier OR permissions, but not both")
	}
	payload, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest(
		http.MethodPost,
		"/v3/workplace/service_accounts",
		bytes.NewReader(payload),
	)
	if err != nil {
		return nil, err
	}

	body, err := dp.makeApiRequest(request)
	if err != nil {
		return nil, err
	}

	data := &ServiceAccountModel{}
	err = json.Unmarshal(body, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (dp *doppler) UpdateServiceAccount(slug string, params ServiceAccountBodyParams) (*ServiceAccountModel, error) {
	if params.WorkplaceRole.Identifier != "" || params.WorkplaceRole.Permissions != nil {
		return nil, errors.New("You may provide an identifier OR permissions, but not both")
	}
	payload, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest(
		http.MethodPatch,
		"/v3/workplace/service_accounts/service_account/"+slug,
		bytes.NewReader(payload),
	)
	if err != nil {
		return nil, err
	}

	body, err := dp.makeApiRequest(request)
	if err != nil {
		return nil, err
	}

	data := &ServiceAccountModel{}
	err = json.Unmarshal(body, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (dp *doppler) DeleteServiceAccount(slug string) (string, error) {

	request, err := http.NewRequest(
		http.MethodDelete,
		"/v3/workplace/service_accounts/service_account/"+slug,
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

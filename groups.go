package doppler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strconv"
)

type Group struct {
	Name               string             `json:"name,omitempty"`
	Slug               string             `json:"slug,omitempty"`
	CreatedAt          string             `json:"created_at,omitempty"`
	DefaultProjectRole DefaultProjectRole `json:"default_project_role,omitempty"`
}

type Groups struct {
	Groups  []Group `json:"groups,omitempty"`
	Success string  `json:"success,omitempty"`
}

type GroupBodyParams struct {
	Name               string `json:"name,omitempty"`
	DefaultProjectRole string `json:"default_project_role,omitempty"` // Identifier of the project role
}

type GroupData struct {
	Group    Group     `json:"group,omitempty"`
	Projects []Project `json:"projects,omitempty"`
	Members  []Member  `json:"members,omitempty"`
}

type DefaultProjectRole struct {
	Identifier string `json:"identifier,omitempty"`
}

type MemberBodyParams struct {
	Type string `json:"type,omitempty"`
	Slug string `json:"slug,omitempty"` // The member's slug
}

func (dp *doppler) ListGroups(page, limit *int) (*Groups, error) {
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
		"/v3/workplace/groups?page="+strconv.Itoa(*page)+"&per_page="+strconv.Itoa(*limit),
		nil,
	)
	if err != nil {
		return nil, err
	}

	body, err := dp.makeApiRequest(request)
	if err != nil {
		return nil, err
	}

	data := &Groups{}
	err = json.Unmarshal(body, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (dp *doppler) RetrieveGroup(slug string) (*GroupData, error) {
	request, err := http.NewRequest(
		http.MethodGet,
		"/v3/workplace/groups/group/"+slug,
		nil,
	)
	if err != nil {
		return nil, err
	}

	body, err := dp.makeApiRequest(request)
	if err != nil {
		return nil, err
	}

	data := &GroupData{}
	err = json.Unmarshal(body, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (dp *doppler) CreateGroup(params GroupBodyParams) (*GroupData, error) {
	payload, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest(
		http.MethodPost,
		"/v3/workplace/groups",
		bytes.NewReader(payload),
	)
	if err != nil {
		return nil, err
	}

	body, err := dp.makeApiRequest(request)
	if err != nil {
		return nil, err
	}

	data := &GroupData{}
	err = json.Unmarshal(body, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (dp *doppler) UpdateGroup(slug string, params GroupBodyParams) (*GroupData, error) {
	payload, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest(
		http.MethodPatch,
		"/v3/workplace/groups/group/"+slug,
		bytes.NewReader(payload),
	)
	if err != nil {
		return nil, err
	}

	body, err := dp.makeApiRequest(request)
	if err != nil {
		return nil, err
	}

	data := &GroupData{}
	err = json.Unmarshal(body, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (dp *doppler) DeleteGroup(slug string) (string, error) {
	request, err := http.NewRequest(
		http.MethodDelete,
		"/v3/workplace/groups/group/"+slug,
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

// TODO: docs incomplete
func (dp *doppler) AddMember(slug string, params MemberBodyParams) (string, error) {
	params.Type = "workplace_user"
	payload, err := json.Marshal(params)
	if err != nil {
		return "", err
	}
	request, err := http.NewRequest(
		http.MethodPost,
		"/v3/workplace/groups/group/"+slug+"/members",
		bytes.NewReader(payload),
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

func (dp *doppler) DeleteMember(slug, memberSlug, memberType string) (string, error) {
	request, err := http.NewRequest(
		http.MethodDelete,
		"/v3/workplace/groups/group/"+slug+"/members/"+memberType+"/"+memberSlug,
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

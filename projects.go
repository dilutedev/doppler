package doppler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

type Projects struct {
	Projects []Project `json:"projects"`
	Page     int       `json:"page"`
	Success  bool      `json:"success"`
}

type Project struct {
	ID          string      `json:"id"`
	Slug        string      `json:"slug"`
	Name        string      `json:"name"`
	Description interface{} `json:"description"`
	CreatedAt   string      `json:"created_at"`
}

type IProject struct {
	Project Project `json:"project"`
	Success bool    `json:"success"`
}

type Success struct {
	Success bool `json:"success"`
}

type CreateProjectParams struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
}

type UpdateProjectParams struct {
	ProjectID   string `json:"project"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
}

// list projects
func (dp *doppler) ListProjects(page int, limit *int) (*Projects, error) {
	default_per_page := 20

	if limit == nil {
		limit = &default_per_page
	}

	request, err := http.NewRequest(
		http.MethodGet,
		"/v3/projects?page="+strconv.Itoa(page)+"&per_page="+strconv.Itoa(*limit),
		nil,
	)
	if err != nil {
		return nil, err
	}

	body, err := dp.makeApiRequest(request)
	if err != nil {
		return nil, err
	}

	data := &Projects{}
	if err := json.Unmarshal(body, data); err != nil {
		return nil, err
	}
	return data, nil
}

// create new project
func (dp *doppler) CreateProject(params CreateProjectParams) (*IProject, error) {
	payload, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	req, _ := http.NewRequest(
		http.MethodPost,
		"/v3/projects",
		bytes.NewReader(payload),
	)

	body, err := dp.makeApiRequest(req)
	if err != nil {
		return nil, err
	}

	data := &IProject{}
	err = json.Unmarshal(body, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// retrieve project
func (dp *doppler) RetrieveProject(id string) (*IProject, error) {
	var data IProject
	req, _ := http.NewRequest(
		http.MethodGet,
		"/v3/projects/project?project="+id,
		nil,
	)

	body, err := dp.makeApiRequest(req)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

// update project
func (dp *doppler) UpdateProject(params UpdateProjectParams) (*IProject, error) {

	payload, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	req, _ := http.NewRequest(
		http.MethodPost,
		"/v3/projects/project",
		bytes.NewReader(payload),
	)

	body, err := dp.makeApiRequest(req)
	if err != nil {
		return nil, err
	}

	data := &IProject{}
	err = json.Unmarshal(body, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// delete project
func (dp *doppler) DeleteProject(project_id string) (*Success, error) {
	req, _ := http.NewRequest(
		http.MethodDelete,
		"/v3/projects/project",
		strings.NewReader(`{"project":"`+project_id+`"}`),
	)

	body, err := dp.makeApiRequest(req)
	if err != nil {
		return nil, err
	}

	data := &Success{}
	err = json.Unmarshal(body, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

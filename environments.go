package doppler

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type Environment struct {
	ID             string `json:"id"`               // An identifier for the object
	Name           string `json:"name"`             // Name of the environment
	Project        string `json:"project"`          // Identifier of the project the environment belongs to
	CreatedAt      string `json:"created_at"`       // Date and time of the object's creation
	InitialFetchAt string `json:"initial_fetch_at"` // Date and time of the first secrets fetch from a config in the environment
}

type Environments struct {
	Environments []Environment `json:"environments"`
	Page         int           `json:"page"`
	Success      bool          `json:"success"`
}

type IEnvironment struct {
	Environment Environment `json:"environment"`
	Success     bool        `json:"success,omitempty"`
}

type EnvironmentBodyParams struct {
	Name string `json:"name"`
	Slug string `json:"slug"`
}

/*
List environments

  - projectID: The project's name
*/
func (dp *doppler) ListEnvironments(projectID string) (*Environments, error) {

	request, err := http.NewRequest(
		http.MethodGet,
		"/v3/environments?project="+projectID,
		nil,
	)
	if err != nil {
		return nil, err
	}

	body, err := dp.makeApiRequest(request)
	if err != nil {
		return nil, err
	}

	data := &Environments{}
	if err := json.Unmarshal(body, data); err != nil {
		return nil, err
	}
	return data, nil
}

/*
Retrieve environment

  - projectID: The project's name
  - environmentID: The environment's slug
*/
func (dp *doppler) RetrieveEnvironment(projectID string, environmentID string) (*IEnvironment, error) {
	request, err := http.NewRequest(
		http.MethodGet,
		"/v3/environments/environment/?project="+projectID+"&environment="+environmentID,
		nil,
	)
	if err != nil {
		return nil, err
	}

	body, err := dp.makeApiRequest(request)
	if err != nil {
		return nil, err
	}

	data := &IEnvironment{}
	if err := json.Unmarshal(body, data); err != nil {
		return nil, err
	}
	return data, nil
}

/*
Create Environment

  - projectID: The project's name
  - params: Request Body parameters
*/
func (dp *doppler) CreateEnvironment(projectID string, params EnvironmentBodyParams) (*IEnvironment, error) {
	payload, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest(
		http.MethodPost,
		"/v3/environments?project="+projectID,
		bytes.NewReader(payload),
	)
	if err != nil {
		return nil, err
	}
	body, err := dp.makeApiRequest(request)
	if err != nil {
		return nil, err
	}
	data := &IEnvironment{}
	err = json.Unmarshal(body, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

/*
Delete Environment

  - projectID: The project's name
  - environmentID: The environment's slug
*/
func (dp *doppler) DeleteEnvironment(projectID string, environmentID string) (*Success, error) {
	request, err := http.NewRequest(
		http.MethodDelete,
		"/v3/environments/environment?project="+projectID+"&environment="+environmentID,
		nil,
	)
	if err != nil {
		return nil, err
	}

	body, err := dp.makeApiRequest(request)
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

/*
Rename Environment

  - projectID: The project's name
  - environmentID: The environment's slug
  - params: Request Body parameters
*/
func (dp *doppler) RenameEnvironment(projectID string, environmentID string, params EnvironmentBodyParams) (*IEnvironment, error) {
	payload, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest(
		http.MethodPost,
		"/v3/environments/environment?project="+projectID+"&environment="+environmentID,
		bytes.NewReader(payload),
	)
	if err != nil {
		return nil, err
	}

	body, err := dp.makeApiRequest(request)
	if err != nil {
		return nil, err
	}

	data := &IEnvironment{}
	err = json.Unmarshal(body, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

package doppler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strconv"
)

type Config struct {
	Name           string `json:"name"`             // Name of the config.
	Project        string `json:"project"`          // Identifier of the project that the config belongs to.
	Environment    string `json:"environment"`      // Identifier of the environment that the config belongs to.
	CreatedAt      string `json:"created_at"`       // Date and time of the object's creation.
	InitialFetchAt string `json:"initial_fetch_at"` // Date and time of the first secrets fetch.
	LastFetchAt    string `json:"last_fetch_at"`    // Date and time of the last secrets fetch.
	Root           bool   `json:"root"`             // Whether the config is the root of the environment.
	Locked         bool   `json:"locked"`           // Whether the config can be renamed and/or deleted.
	Slug           string `json:"slug"`
}

type Configs struct {
	Configs []Config `json:"configs"`
	Page    int      `json:"page"`
	Success bool     `json:"success"`
}

type CreateConfigParams struct {
	Project     string `json:"project"`     // Unique identifier for the project object
	Environment string `json:"environment"` // Identifier for the environment object
	Name        string `json:"name"`        // Name of the new branch config
}

// ModifyConfigParams is used for update and clone params
type ModifyConfigParams struct {
	Project string `json:"project"` // Unique identifier for the project (project name)
	Config  string `json:"config"`  // Name of the config object
	Name    string `json:"name"`    // The new name of config
}

type DeletConfigParams struct {
	Project string `json:"project"` // Unique identifier for the project (project name)
	Config  string `json:"config"`  // Name of the config object
}

type IConfig struct {
	Config  Config `json:"config"`
	Success bool   `json:"success"`
}

/*
Fetch all configs

  - projectName: Unique identifier for the project
  - environmentSlug: The environment's slug
  - page: page number
  - perPage: Items per page (default: 20)
*/
func (dp *doppler) ListConfigs(projectName, environmentSlug string, page int, perPage *int) (*Configs, error) {
	defaultPerPage := 20
	if perPage == nil {
		perPage = &defaultPerPage
	}
	request, err := http.NewRequest(
		http.MethodGet,
		"/v3/configs?project="+projectName+"&environment="+environmentSlug+"&page="+strconv.Itoa(page)+"&per_page="+strconv.Itoa(*perPage),
		nil,
	)
	if err != nil {
		return nil, err
	}

	data := &Configs{}
	body, err := dp.makeApiRequest(request)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

/*
Create a new branch config.
*/
func (dp *doppler) CreateConfig(params CreateConfigParams) (*IConfig, error) {
	payload, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}
	request, _ := http.NewRequest(
		http.MethodPost,
		"/v3/configs",
		bytes.NewReader(payload),
	)

	body, err := dp.makeApiRequest(request)
	if err != nil {
		return nil, err
	}
	data := &IConfig{}
	err = json.Unmarshal(body, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

/*
Fetch a config's details

  - project: Unique identifier for the project
  - config: Name of the config object
*/
func (dp *doppler) RetrieveConfig(project, config string) (*IConfig, error) {
	request, err := http.NewRequest(
		http.MethodGet,
		"/v3/configs/config?project="+project+"&config="+config,
		nil,
	)
	if err != nil {
		return nil, err
	}

	body, err := dp.makeApiRequest(request)
	if err != nil {
		return nil, err
	}
	data := &IConfig{}
	err = json.Unmarshal(body, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

/*
Modify an existing config.
*/
func (dp *doppler) UpdateConfig(params ModifyConfigParams) (*IConfig, error) {
	payload, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest(
		http.MethodPost,
		"/v3/configs/config",
		bytes.NewReader(payload),
	)
	if err != nil {
		return nil, err
	}

	body, err := dp.makeApiRequest(request)
	if err != nil {
		return nil, err
	}

	data := &IConfig{}
	err = json.Unmarshal(body, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

/*
Permanently delete the config.
*/
func (dp *doppler) DeleteConfig(params DeletConfigParams) (*Success, error) {
	payload, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest(
		http.MethodDelete,
		"/v3/configs/config",
		bytes.NewReader(payload),
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
Create a new branch config by cloning another.
This duplicates a branch config and all its secrets.
*/
func (dp *doppler) CloneConfig(params ModifyConfigParams) (*IConfig, error) {
	payload, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest(
		http.MethodPost,
		"/v3/configs/config/clone",
		bytes.NewReader(payload),
	)

	if err != nil {
		return nil, err
	}
	body, err := dp.makeApiRequest(request)
	if err != nil {
		return nil, err
	}

	data := &IConfig{}
	err = json.Unmarshal(body, data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

/*
Prevent the config from being renamed or deleted.
*/
func (dp *doppler) LockConfig(params DeletConfigParams) (*IConfig, error) {
	payload, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest(
		http.MethodPost,
		"/v3/configs/config/lock",
		bytes.NewReader(payload),
	)

	if err != nil {
		return nil, err
	}
	body, err := dp.makeApiRequest(request)
	if err != nil {
		return nil, err
	}

	data := &IConfig{}
	err = json.Unmarshal(body, data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

/*
Allow the config to be renamed and/or deleted.
*/
func (dp *doppler) UnlockConfig(params DeletConfigParams) (*IConfig, error) {
	payload, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest(
		http.MethodPost,
		"/v3/configs/config/unlock",
		bytes.NewReader(payload),
	)

	if err != nil {
		return nil, err
	}
	body, err := dp.makeApiRequest(request)
	if err != nil {
		return nil, err
	}

	data := &IConfig{}
	err = json.Unmarshal(body, data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

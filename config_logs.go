package doppler

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type ConfigLog struct {
	ID          string          `json:"id"`   // Unique identifier for the object
	Text        string          `json:"text"` // Text describing the event
	HTML        string          `json:"html"` // HTML describing the event
	Diff        []ConfigLogDiff `json:"diff,omitempty"`
	Rollback    bool            `json:"rollback"` // Is this config log a rollback of a previous log
	User        User            `json:"user"`
	Project     string          `json:"project"`     // Unique identifier for the object
	Environment string          `json:"environment"` // Unique identifier for the enironment object
	Config      string          `json:"config"`      // The config's name
	CreatedAt   string          `json:"created_at"`  // Date and time of the object's creation
}

type ConfigLogDiff struct {
	Name    string `json:"name"`
	Added   string `json:"added"`
	Removed string `json:"removed"`
}

type ConfigLogs struct {
	Logs    []ConfigLog `json:"logs"`
	Page    int         `json:"page"`
	Success bool        `json:"success"`
}

type IConfigLog struct {
	Log     ConfigLog `json:"log"`
	Success bool      `json:"success"`
}

type RollbackConfigLogParams struct {
	Project string `json:"project"` // Unique identifier for the project (project name)
	Config  string `json:"config"`  // Name of the config object
	Log     string `json:"log"`     // Unique identifier for the log object
}

/*
List config logs

  - projectName: Unique identifier for the project
  - config: Name of the config object
  - page: page number
  - perPage: Items per page (default: 20)
*/
func (dp *doppler) ListConfigLogs(projectName, config string, page int, perPage *int) (*ConfigLogs, error) {
	defaultPerPage := 20
	if perPage == nil {
		perPage = &defaultPerPage
	}
	request, err := http.NewRequest(
		http.MethodGet,
		"/v3/configs/config/logs?project="+projectName+"&config="+config+"&page="+strconv.Itoa(page)+"&per_page="+strconv.Itoa(*perPage),
		nil,
	)
	if err != nil {
		return nil, err
	}

	data := &ConfigLogs{}
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
Fetch a config log's details

  - project: Unique identifier for the project
  - config: Name of the config object
  - logID: Unique identifier for the log object
*/
func (dp *doppler) RetrieveConfigLog(project, config, logID string) (*IConfigLog, error) {
	request, err := http.NewRequest(
		http.MethodGet,
		"/v3/configs/config/logs/log?project="+project+"&config="+config+"&log="+logID,
		nil,
	)
	if err != nil {
		return nil, err
	}

	body, err := dp.makeApiRequest(request)
	if err != nil {
		return nil, err
	}
	data := &IConfigLog{}
	err = json.Unmarshal(body, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

/*
- project: Unique identifier for the project
- config: Name of the config object
- logID: Unique identifier for the log object
*/
func (dp *doppler) RollbackConfigLog(project, config, logID string) (*IConfigLog, error) {
	request, err := http.NewRequest(
		http.MethodPost,
		"/v3/configs/config/logs/log/rollback?project="+project+"&config="+config+"&log="+logID,
		nil,
	)

	body, err := dp.makeApiRequest(request)
	if err != nil {
		return nil, err
	}
	data := &IConfigLog{}
	err = json.Unmarshal(body, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

package doppler

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type ActivityLogs struct {
	Page    int   `json:"page"`
	Logs    []Log `json:"logs"`
	Success bool  `json:"success"`
}

type ActivityLog struct {
	Log     Log  `json:"log"`
	Success bool `json:"success"`
}

type Log struct {
	ID                 string  `json:"id"`
	Text               string  `json:"text"`
	HTML               string  `json:"html"`
	CreatedAt          string  `json:"created_at"`
	EnclaveConfig      *string `json:"enclave_config"`
	EnclaveEnvironment *string `json:"enclave_environment"`
	EnclaveProject     *string `json:"enclave_project"`
	User               User    `json:"user"`
	Diff               *Diff   `json:"diff"`
}

type Diff struct {
	Name    *string  `json:"name"`
	Added   []string `json:"added"`
	Removed []string `json:"removed"`
	Updated []string `json:"updated"`
}

type User struct {
	Email           string `json:"email"`
	Name            string `json:"name"`
	Username        string `json:"username"`
	ProfileImageURL string `json:"profile_image_url"`
}

// list activity logs
func (dp *doppler) RetrieveLogs(page int, limit *int) (*ActivityLogs, error) {
	var (
		default_per_page int = 20
		data             ActivityLogs
	)

	if limit == nil {
		limit = &default_per_page
	}

	request, err := http.NewRequest(
		http.MethodGet,
		"/v3/logs?page="+strconv.Itoa(page)+"&per_page="+strconv.Itoa(*limit),
		nil,
	)

	if err != nil {
		return nil, err
	}

	body, err := dp.makeApiRequest(request)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(body, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

// get activity log
func (dp *doppler) RetrieveLog(log_id string) (*ActivityLog, error) {
	var (
		request, err = http.NewRequest(
			http.MethodGet,
			"/v3/logs/log?log="+log_id,
			nil,
		)
		data ActivityLog
	)

	if err != nil {
		return nil, err
	}

	body, err := dp.makeApiRequest(request)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(body, &data); err != nil {
		return nil, err
	}

	return &data, nil
}

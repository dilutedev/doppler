package doppler

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type Sync struct {
	Slug         string `json:"slug,omitempty"`
	Integration  string `json:"integration,omitempty"`
	Project      string `json:"project,omitempty"`
	Config       string `json:"config,omitempty"`
	Enabled      bool   `json:"enabled,omitempty"`
	LastSyncedAt string `json:"last_synced_at,omitempty"`
}

type SyncData struct {
	Sync    Sync `json:"sync,omitempty"`
	Success bool `json:"success,omitempty"`
}

type SyncQueryParams struct {
	Project          string // The project slug
	Config           string // The config slug
	Sync             string // The sync slug: use with RetrieveSync and DeleteSync
	DeleteFromTarget bool   // use with DeleteSync function only
}

type SyncBodyParams struct {
	Integration  string // The integration slug which the sync will use
	ImportOption string // prefer_doppler or prefer_integration, defaults to none
}

// Create a new secrets sync.
func (dp *doppler) CreateSync(queryParams SyncQueryParams, bodyParams SyncBodyParams) (*SyncData, error) {
	if bodyParams.ImportOption == "" {
		bodyParams.ImportOption = "none"
	}
	request, err := http.NewRequest(
		http.MethodPost,
		"/v3/configs/config/syncs?project="+queryParams.Project+"&config="+queryParams.Config,
		nil,
	)
	if err != nil {
		return nil, err
	}

	body, err := dp.makeApiRequest(request)
	if err != nil {
		return nil, err
	}

	data := &SyncData{}
	err = json.Unmarshal(body, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// Retrieve an existing secrets sync.
func (dp *doppler) RetrieveSync(params SyncQueryParams) (*SyncData, error) {
	request, err := http.NewRequest(
		http.MethodGet,
		"/v3/configs/config/syncs/sync?project="+params.Project+"&config="+params.Config+"&sync="+params.Sync,
		nil,
	)

	if err != nil {
		return nil, err
	}

	body, err := dp.makeApiRequest(request)
	if err != nil {
		return nil, err
	}

	data := &SyncData{}
	err = json.Unmarshal(body, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// Delete an existing sync.
func (dp *doppler) DeleteSync(params SyncQueryParams) (string, error) {
	request, err := http.NewRequest(
		http.MethodDelete,
		"/v3/configs/config/syncs/sync?project="+params.Project+"&config="+params.Config+"&sync="+params.Sync+"&delete_from_target="+strconv.FormatBool(params.DeleteFromTarget),
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

package doppler

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type ServiceToken struct {
	Name         string `json:"name,omitempty"`
	Slug         string `json:"slug,omitempty"`
	Access       string `json:"access,omitempty"`
	CreatedAt    string `json:"created_at,omitempty"`
	Key          string `json:"key,omitempty"`
	Project      string `json:"project,omitempty"`
	Environment  string `json:"environment,omitempty"`
	Config       string `json:"config,omitempty"`
	ExpiresAt    string `json:"expires_at,omitempty"`
	LastSeenAt   string `json:"last_seen_at,omitempty"`
	TokenPreview string `json:"token_preview,omitempty"`
}

type ServiceTokens struct {
	Tokens  []ServiceToken `json:"tokens,omitempty"`
	Success bool           `json:"success,omitempty"`
}

type ServiceTokenModel struct {
	Token   ServiceToken `json:"token,omitempty"`
	Success bool         `json:"success,omitempty"`
}

type CreateTokenParams struct {
	Project  string `json:"project,omitempty"`   // Unique identifier for the project object.
	Config   string `json:"config,omitempty"`    // Name of the config object.
	Name     string `json:"name,omitempty"`      // Name of the service token.
	ExpireAt string `json:"expire_at,omitempty"` // Unix timestamp of when token should expire.
	Access   string `json:"access,omitempty"`    // Token's capabilities. "read/write" or "read". Default: read
}

type DeleteTokenParams struct {
	Project string `json:"project,omitempty"` // Unique identifier for the project object.
	Config  string `json:"config,omitempty"`  // Name of the config object.
	Slug    string `json:"slug,omitempty"`    // The slug of the service token.
}

func (dp *doppler) ListServiceTokens(project, config string) (*ServiceTokens, error) {
	request, err := http.NewRequest(
		http.MethodGet,
		"/v3/configs/config/tokens?project="+project+"&config="+config,
		nil,
	)
	if err != nil {
		return nil, err
	}

	body, err := dp.makeApiRequest(request)
	if err != nil {
		return nil, err
	}

	data := &ServiceTokens{}
	err = json.Unmarshal(body, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (dp *doppler) CreateServiceToken(params CreateTokenParams) (*ServiceTokenModel, error) {
	if params.Access != "read" && params.Access != "read/write" {
		params.Access = "read"
	}
	payload, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest(
		http.MethodPost,
		"/v3/configs/config/tokens",
		bytes.NewReader(payload),
	)
	if err != nil {
		return nil, err
	}

	body, err := dp.makeApiRequest(request)
	if err != nil {
		return nil, err
	}

	data := &ServiceTokenModel{}
	err = json.Unmarshal(body, data)
	if err != nil {
		return nil, err
	}

	return data, nil

}

func (dp *doppler) DeleteServiceToken(params DeleteTokenParams) (*Success, error) {
	payload, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest(
		http.MethodDelete,
		"/v3/configs/config/tokens/token",
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

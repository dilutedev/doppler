package doppler

import (
	"encoding/json"
	"net/http"
	"strings"
)

type Integrations struct {
	Integrations []Integration `json:"integrations,omitempty"`
	Success      bool          `json:"success,omitempty"`
}

type Integration struct {
	Slug    string `json:"slug,omitempty"`
	Name    string `json:"name,omitempty"`
	Type    string `json:"type,omitempty"`
	Kind    string `json:"kind,omitempty"`
	Enabled bool   `json:"enabled,omitempty"`
	// Sync
}

type IntegrationData struct {
	Integration Integration `json:"integration,omitempty"`
	Success     bool        `json:"success,omitempty"`
}

type UpdateIntegrationParams struct {
	Name        string `json:"name,omitempty"`        // The new name of the integration
	Data        string `json:"data,omitempty"`        // The new authentication data for the integration
	Integration string `json:"integration,omitempty"` // The slug of the integration to update
}

func (dp *doppler) ListIntegrations() (*Integrations, error) {
	request, err := http.NewRequest(
		http.MethodGet,
		"/v3/integrations",
		nil,
	)
	if err != nil {
		return nil, err
	}

	body, err := dp.makeApiRequest(request)
	if err != nil {
		return nil, err
	}

	data := &Integrations{}
	err = json.Unmarshal(body, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// TODO: docs incomplete
func (dp *doppler) CreateIntegration(integrationType, name string) (*IntegrationData, error) {
	request, err := http.NewRequest(
		http.MethodPost,
		"/v3/integrations",
		strings.NewReader("{\"name\":\""+name+"\",\"type\":\""+integrationType+"\"}"),
	)
	if err != nil {
		return nil, err
	}

	body, err := dp.makeApiRequest(request)
	if err != nil {
		return nil, err
	}

	data := &IntegrationData{}
	err = json.Unmarshal(body, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (dp *doppler) RetrieveIntegration(integrationSlug string) (*IntegrationData, error) {
	request, err := http.NewRequest(
		http.MethodGet,
		"/v3/integrations/integration?integration="+integrationSlug,
		nil,
	)
	if err != nil {
		return nil, err
	}

	body, err := dp.makeApiRequest(request)
	if err != nil {
		return nil, err
	}

	data := &IntegrationData{}
	err = json.Unmarshal(body, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// TODO
func (dp *doppler) UpdateIntegration(params UpdateIntegrationParams) (string, error) { // no response according to docs
	request, err := http.NewRequest(
		http.MethodPut,
		"/v3/integrations/integration?integration="+params.Integration,
		nil,
	)
	if err != nil {
		return "", err
	}

	body, err := dp.makeApiRequest(request)
	if err != nil {
		return "", err
	}

	// data := &Integration{}
	// err = json.Unmarshal(body, data)
	// if err != nil {
	// 	return nil, err
	// }

	return string(body), nil
}

// TODO
func (dp *doppler) DeleteIntegration(integration string) (string, error) { // not sure, can't test
	request, err := http.NewRequest(
		http.MethodDelete,
		"/v3/integrations/integration?integration="+integration,
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

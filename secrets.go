package doppler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

type Secrets struct {
	Secrets map[string]Secret `json:"secrets"`
	Success bool              `json:"success"`
}

type SecretNames struct {
	Names   []string `json:"names"`
	Success bool     `json:"success"`
}

type Secret struct {
	Name    string      `json:"name"`
	Value   SecretValue `json:"value"`
	Success bool        `json:"success"`
}

type SecretValue struct {
	Raw                string `json:"raw"`
	Computed           string `json:"computed"`
	Note               string `json:"note"`
	RawVisibility      string `json:"rawVisibility,omitempty"`
	ComputedVisibility string `json:"computedVisibility,omitempty"`
}

type ListSecretsParams struct {
	Project               string // Unique identifier for the project object.
	Config                string // Name of the config object.
	IncludeDynamicSecrets bool   // Whether or not to issue leases and include dynamic secret values for the config
	DynamicSecretsTTLSec  int    // The number of seconds until dynamic leases expire. Must be used with IncludeDynamicSecrets. Defaults to 1800 (30 minutes).
	Secrets               string // A comma-separated list of secrets to include in the response, may only contain uppercase letters, numbers, and underscores.
	IncludeManagedSecrets bool   // Whether to include Doppler's auto-generated (managed) secrets. defaults to false
}

type ListSecretNamesParams struct {
	Project               string // Unique identifier for the project object.
	Config                string // Name of the config object.
	IncludeDynamicSecrets bool   // Whether or not to issue leases and include dynamic secret values for the config
	IncludeManagedSecrets bool   // Whether to include Doppler's auto-generated (managed) secrets. defaults to false
}

type UpdateSecretParams struct {
	Project string            `json:"project"`
	Config  string            `json:"config"`
	Secrets map[string]string `json:"secrets"` // map of secret to new value
}

type SetNoteParams struct {
	Project string `json:"project"` // Unique identifier for the project object(project name)
	Config  string `json:"config"`  // Name of the config object
	Secret  string `json:"secret"`
	Note    string `json:"note"`
}

type NoteResponse struct {
	Secret string `json:"secret"`
	Note   string `json:"note"`
}

type DownloadSecretParams struct {
	Project               string
	Config                string
	Format                string // Acceptable values: json, env, yaml, docker, env-no-quotes, dotnet-json. defaults json
	NameTransformer       string // Acceptable values: camel, upper-camel, lower-snake, tf-var, dotnet, dotnet-env, lower-kebab. defaults to upper snake case
	IncludeDynamicSecrets bool
	DynamicSecretsTTLSec  int
}

func (dp *doppler) ListSecrets(params ListSecretsParams) (*Secrets, error) {
	var url strings.Builder
	url.WriteString("/v3/configs/config/secrets?project=")
	url.WriteString(params.Project)
	url.WriteString("&config=")
	url.WriteString(params.Config)
	if params.IncludeDynamicSecrets {
		url.WriteString("&include_dynamic_secrets=true&dynamic_secrets_ttl_sec=")
		if params.DynamicSecretsTTLSec == 0 {
			params.DynamicSecretsTTLSec = 1800
		}
		url.WriteString(strconv.Itoa(params.DynamicSecretsTTLSec))
	} else {
		url.WriteString("&include_dynamic_secrets=false")
	}
	url.WriteString("&secrets=")
	url.WriteString(params.Secrets)
	url.WriteString("&include_managed_secrets=")
	url.WriteString(strconv.FormatBool(params.IncludeManagedSecrets))
	request, err := http.NewRequest(
		http.MethodGet,
		url.String(),
		nil,
	)
	if err != nil {
		return nil, err
	}

	body, err := dp.makeApiRequest(request)
	if err != nil {
		return nil, err
	}

	data := &Secrets{}
	err = json.Unmarshal(body, data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// List Secret names
func (dp *doppler) ListSecretNames(params ListSecretNamesParams) (*SecretNames, error) {
	var url strings.Builder
	url.WriteString("/v3/configs/config/secrets/names?project=")
	url.WriteString(params.Project)
	url.WriteString("&config=")
	url.WriteString(params.Config)
	url.WriteString("&include_dynamic_secrets=")
	url.WriteString(strconv.FormatBool(params.IncludeDynamicSecrets))
	url.WriteString("&include_managed_secrets=")
	url.WriteString(strconv.FormatBool(params.IncludeManagedSecrets))
	request, err := http.NewRequest(
		http.MethodGet,
		url.String(),
		nil,
	)
	if err != nil {
		return nil, err
	}

	body, err := dp.makeApiRequest(request)
	if err != nil {
		return nil, err
	}

	data := &SecretNames{}
	err = json.Unmarshal(body, data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

/*
Retrieves a secret

  - project: Unique identifier for the project object
  - config: Name of the config object
  - name: Name of the secret
*/
func (dp *doppler) RetrieveSecret(project, config, name string) (*Secret, error) {
	request, err := http.NewRequest(
		http.MethodGet,
		"/v3/configs/config/secret?project="+project+"&config="+config+"&name="+name,
		nil,
	)
	if err != nil {
		return nil, err
	}

	body, err := dp.makeApiRequest(request)
	if err != nil {
		return nil, err
	}

	data := &Secret{}
	err = json.Unmarshal(body, data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (dp *doppler) DeleteSecret(project, config, name string) (string, error) {
	request, err := http.NewRequest(
		http.MethodDelete,
		"/v3/configs/config/secret?project="+project+"&config="+config+"&name="+name,
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

func (dp *doppler) UpdateSecret(params UpdateSecretParams) (*Secrets, error) {
	payload, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest(
		http.MethodGet,
		"/v3/configs/config/secrets",
		bytes.NewReader(payload),
	)
	if err != nil {
		return nil, err
	}

	body, err := dp.makeApiRequest(request)
	if err != nil {
		return nil, err
	}

	data := &Secrets{}
	err = json.Unmarshal(body, data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (dp *doppler) DownloadSecret(params DownloadSecretParams) (string, error) {
	var url strings.Builder
	url.WriteString("/v3/configs/config/secrets/download?project=")
	url.WriteString(params.Project)
	url.WriteString("&config=")
	url.WriteString(params.Config)
	url.WriteString("&format=")
	url.WriteString(params.Format)
	url.WriteString("&name_transformer=")
	url.WriteString(params.NameTransformer)
	if params.IncludeDynamicSecrets {
		url.WriteString("&include_dynamic_secrets=true&dynamic_secrets_ttl_sec=")
		if params.DynamicSecretsTTLSec == 0 {
			params.DynamicSecretsTTLSec = 1800
		}
		url.WriteString(strconv.Itoa(params.DynamicSecretsTTLSec))
	} else {
		url.WriteString("&include_dynamic_secrets=false")
	}
	request, err := http.NewRequest(
		http.MethodGet,
		url.String(),
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

// Set a note on a secret
func (dp *doppler) UpdateNote(params SetNoteParams) (*NoteResponse, error) {
	payload, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest(
		http.MethodPost,
		"/v3/configs/config/secrets/note",
		bytes.NewReader(payload),
	)
	if err != nil {
		return nil, err
	}

	body, err := dp.makeApiRequest(request)
	if err != nil {
		return nil, err
	}

	data := &NoteResponse{}
	err = json.Unmarshal(body, data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

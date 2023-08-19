package doppler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"regexp"
)

var (
	TokenFormats = map[string]*regexp.Regexp{
		"service":        regexp.MustCompile(`dp\.st\.(?:[a-z0-9\-_]{2,35}\.)?[a-zA-Z0-9]{40,44}`),
		"cli":            regexp.MustCompile(`dp\.ct\.[a-zA-Z0-9]{40,44}`),
		"personal":       regexp.MustCompile(`dp\.pt\.[a-zA-Z0-9]{40,44}`),
		"service_access": regexp.MustCompile(`dp\.sa\.[a-zA-Z0-9]{40,44}`),
		"scim":           regexp.MustCompile(`dp\.scim\.[a-zA-Z0-9]{40,44}`),
		"audit":          regexp.MustCompile(`dp\.audit\.[a-zA-Z0-9]{40,44}`),
	}
)

type token struct {
	Key  string
	Type string
}

type RevokeTokenParam struct {
	Token string `json:"token,omitempty"`
}

// validate auth token
func (t *token) validate() error {
	for _type, pattern := range TokenFormats {
		if pattern.MatchString(t.Key) {
			t.Type = _type
			return nil
		}
	}
	return ErrInvalidToken
}

// Revoke auth tokens
func (dp *doppler) RevokeTokens(params []RevokeTokenParam) error {
	payload, err := json.Marshal(params)
	if err != nil {
		return err
	}
	req, err := http.NewRequest(
		http.MethodPost,
		"/v3/auth/revoke",
		bytes.NewReader(payload),
	)

	if err != nil {
		return err
	}

	_, err = dp.makeApiRequest(req)
	return err
}

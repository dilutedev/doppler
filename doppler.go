package doppler

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"os"
)

type doppler struct {
	client   *http.Client
	token    token
	base_url string
}

// failed response
type FailedRequest struct {
	Messages []string `json:"messages"`
	Success  bool     `json:"success"`
}

// create new doppler config
func New(key string) (*doppler, error) {
	apiToken := token{
		Key: key,
	}

	if err := apiToken.validate(); err != nil {
		return nil, err
	}

	return &doppler{
		client:   &http.Client{},
		token:    apiToken,
		base_url: "api.doppler.com",
	}, nil
}

func NewFromEnv() (*doppler, error) {
	apiToken := token{
		Key: os.Getenv("DOPPLER_KEY"),
	}

	if err := apiToken.validate(); err != nil {
		return nil, err
	}

	return &doppler{
		client:   &http.Client{},
		token:    apiToken,
		base_url: "api.doppler.com",
	}, nil
}

func (dp *doppler) makeApiRequest(request *http.Request) ([]byte, error) {
	var (
		errData     *FailedRequest
		request_url = &url.URL{
			Host:       dp.base_url,
			Scheme:     "https",
			Path:       request.URL.Path,
			RawQuery:   request.URL.RawQuery,
			ForceQuery: request.URL.ForceQuery,
		}
	)

	request.URL = request_url
	request.Header = http.Header{
		"accept":        {"application/json"},
		"content-type":  {"application/json"},
		"authorization": {"Bearer " + dp.token.Key},
	}

	res, err := dp.client.Do(request)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 && res.StatusCode != 204 {
		if err := json.Unmarshal(body, &errData); err != nil {
			return nil, err
		}

		return nil, &DopplerError{
			Status:   res.StatusCode,
			Messages: errData.Messages,
		}
	}

	return body, nil
}

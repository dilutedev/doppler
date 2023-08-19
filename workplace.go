package doppler

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type Workplace struct {
	Workplace WorkplaceClass `json:"workplace"`
	Success   bool           `json:"success"`
}

type WorkplaceClass struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	BillingEmail string `json:"billing_email"`
}

type WorkplaceParams struct {
	Name         string `json:"name,omitempty"`
	BillingEmail string `json:"billing_email,omitempty"`
}

// get workplace info
func (dp *doppler) GetWorkplace() (*Workplace, error) {
	var (
		request, err = http.NewRequest(
			http.MethodGet,
			"/v3/workplace",
			nil,
		)
		workplace = Workplace{}
	)

	if err != nil {
		return nil, err
	}

	body, err := dp.makeApiRequest(request)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(body, &workplace); err != nil {
		return nil, err
	}

	return &workplace, nil
}

// update work place info
func (dp *doppler) UpdateWorkplace(params WorkplaceParams) (*Workplace, error) {
	payload, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest(
		http.MethodPost,
		"/v3/workplace",
		bytes.NewReader(payload),
	)
	workplace := Workplace{}

	if err != nil {
		return nil, err
	}

	body, err := dp.makeApiRequest(request)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(body, &workplace); err != nil {
		return nil, err
	}

	return &workplace, nil
}

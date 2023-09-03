package doppler

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type IssueLeaseArgs struct {
	Project       string `json:"project"`
	Config        string `json:"config"`
	DynamicSecret string `json:"dynamic_secret"`
	TtlSec        int32  `json:"ttl_sec"`
}

/*
Issue a lease for a dynamic secret
*/
func (dp *doppler) IssueLease(args IssueLeaseArgs) (data *interface{}, err error) {

	payload, err := json.Marshal(args)
	if err != nil {
		return
	}

	req, _ := http.NewRequest(
		http.MethodPost,
		"/v3/configs/config/dynamic_secrets/dynamic_secret/leases",
		bytes.NewReader(payload))

	body, err := dp.makeApiRequest(req)
	if err != nil {
		return
	}

	data = new(interface{})
	err = json.Unmarshal(body, &data)
	if err != nil {
		return
	}

	return
}

type RevokeLeaseArgs struct {
	Project       string `json:"project"`
	Config        string `json:"config"`
	DynamicSecret string `json:"dynamic_secret"`
	Slug          string `json:"slug"`
}
type RevokeLeaseData struct {
	Success bool `json:"success"`
}

/*
Revoke a lease for a dynamic secret
*/
func (dp *doppler) RevokeLease(args RevokeLeaseArgs) (data *RevokeLeaseData, err error) {
	payload, err := json.Marshal(args)
	if err != nil {
		return
	}

	req, _ := http.NewRequest(
		http.MethodDelete,
		"/v3/configs/config/dynamic_secrets/dynamic_secret/leases/lease",
		bytes.NewReader(payload))

	body, err := dp.makeApiRequest(req)
	if err != nil {
		return
	}

	data = new(RevokeLeaseData)
	err = json.Unmarshal(body, &data)
	if err != nil {
		return
	}

	return
}

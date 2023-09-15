package doppler

import (
	"encoding/json"
	"net/http"
	"strings"
)

type IPs struct {
	IPs     []string `json:"ips"`
	Success bool     `json:"success"`
}

type IP struct {
	IP string `json:"ip"`
}

/*
List trusted IPs
*/
func (dp *doppler) ListTrustedIPs(project, config string) (*IPs, error) {
	request, err := http.NewRequest(
		http.MethodGet,
		"/v3/configs/config/trusted_ips?project="+project+"&config="+config,
		nil,
	)
	if err != nil {
		return nil, err
	}
	body, err := dp.makeApiRequest(request)
	if err != nil {
		return nil, err
	}

	data := &IPs{}
	err = json.Unmarshal(body, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

/*
Add a trusted IP
*/
func (dp *doppler) AddIP(project, config, ip string) (*IP, error) {
	payload := strings.NewReader("{\"ip\":\"" + ip + "\"}")
	request, err := http.NewRequest(
		http.MethodPost,
		"/v3/configs/config/trusted_ips?project="+project+"&config="+config,
		payload,
	)
	if err != nil {
		return nil, err
	}
	body, err := dp.makeApiRequest(request)
	if err != nil {
		return nil, err
	}

	data := &IP{}
	err = json.Unmarshal(body, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

/*
Delete a trusted IP
*/
func (dp *doppler) DeleteIP(project, config, ip string) (string, error) {
	payload := strings.NewReader("{\"ip\":\"" + ip + "\"}")
	request, err := http.NewRequest(
		http.MethodDelete,
		"/v3/configs/config/trusted_ips?project="+project+"&config="+config,
		payload,
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

package doppler

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
)

type WorkplaceUsers struct {
	WorkplaceUsers []WorkplaceUser `json:"workplace_users"`
	Page           int64           `json:"page"`
	Success        bool            `json:"success"`
}

type WorkplaceUser struct {
	ID        string `json:"id"`
	Access    string `json:"access"`
	CreatedAt string `json:"created_at"`
	User      WUser  `json:"user"`
}

type WUser struct {
	Email                string `json:"email"`
	Name                 string `json:"name"`
	Username             string `json:"username"`
	ProfileImageURL      string `json:"profile_image_url"`
	MfaEnabled           bool   `json:"mfa_enabled"`
	ThirdpartySsoEnabled bool   `json:"thirdparty_sso_enabled"`
	SamlSsoEnabled       bool   `json:"saml_sso_enabled"`
}

type WorkplaceUserResp struct {
	WorkplaceUser WorkplaceUser `json:"workplace_user"`
	Success       bool          `json:"success"`
}

/*
Get all users of a workplace
@param settings: bool
If true, the api will return more information if users have e.g. SAML enabled and/or Multi Factor Auth enabled
*/
func (dp *doppler) GetWorkplaceUsers(settings bool, page int32) (*WorkplaceUsers, error) {
	// support only audit tokens
	if dp.token.Type != "audit" {
		return nil, errors.New("audit token required")
	}

	req, err := http.NewRequest(
		http.MethodGet,
		fmt.Sprintf("/v3/workplace/users?settings=%v&page=%d", settings, page),
		nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	body, err := dp.makeApiRequest(req)
	if err != nil {
		return nil, err
	}

	data := new(WorkplaceUsers)
	err = json.Unmarshal(body, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

/*
Get a specific user in a workplace
@param settings: bool
If true, the api will return more information if user has e.g. SAML enabled and/or Multi Factor Auth enabled
*/
func (dp *doppler) GetWorkplaceUser(user_id string, settings bool) (*WorkplaceUser, error) {
	// support only audit tokens
	if dp.token.Type != "audit" {
		return nil, errors.New("audit token required")
	}

	req, err := http.NewRequest(
		http.MethodGet,
		fmt.Sprintf("/v3/workplace/users/%s?settings=%v", user_id, settings),
		nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	body, err := dp.makeApiRequest(req)
	if err != nil {
		return nil, err
	}

	data := new(WorkplaceUserResp)
	err = json.Unmarshal(body, data)
	if err != nil {
		return nil, err
	}

	return &data.WorkplaceUser, nil
}

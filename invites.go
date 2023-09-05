package doppler

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type Invites struct {
	Invites []Invite `json:"invites"`
	Success bool     `json:"success"`
}

type Invite struct {
	Slug          string        `json:"slug"`
	Email         string        `json:"email"`
	CreatedAt     string        `json:"created_at"`
	WorkplaceRole WorkplaceRole `json:"workplace_role"`
}

func (dp *doppler) ListInvites(page, limit *int) (*Invites, error) {
	defaultLimit := 20
	defaultPage := 1
	if page == nil || *page <= 0 {
		page = &defaultPage
	}
	if limit == nil || *limit <= 0 {
		limit = &defaultLimit
	}
	request, err := http.NewRequest(
		http.MethodGet,
		"/v3/workplace/invites?page="+strconv.Itoa(*page)+"&per_page="+strconv.Itoa(*limit),
		nil,
	)
	if err != nil {
		return nil, err
	}

	body, err := dp.makeApiRequest(request)
	if err != nil {
		return nil, err
	}

	data := &Invites{}
	err = json.Unmarshal(body, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

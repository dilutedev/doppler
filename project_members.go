package doppler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type MemberType int

const (
	MemberWorkplaceUser MemberType = iota
	MemberGroup
	MemberInvite
	MemberServiceAccount
)

type ProjectMembersResp struct {
	Members []Member `json:"members"`
	Success bool     `json:"success"`
}

type Member struct {
	Type                  string      `json:"type"`
	Slug                  string      `json:"slug"`
	Role                  Role        `json:"role"`
	AccessAllEnvironments bool        `json:"access_all_environments"`
	Environments          interface{} `json:"environments"`
}

type Role struct {
	Identifier string `json:"identifier"`
}

type ProjectMemberResp struct {
	Member  Member `json:"member"`
	Success bool   `json:"success"`
}

/*
List All Project Members
*/
func (dp *doppler) ListProjectMembers(project_id string, page, per_page int32) (*ProjectMembersResp, error) {

	req, err := http.NewRequest(
		http.MethodGet,
		fmt.Sprintf("/v3/projects/project/members?project=%s&page=%d&per_page=%d", project_id, page, per_page),
		nil)
	if err != nil {
		return nil, err
	}

	body, err := dp.makeApiRequest(req)
	if err != nil {
		return nil, err
	}

	data := new(ProjectMembersResp)
	err = json.Unmarshal(body, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

/*Retrieve Project Member*/
func (dp *doppler) RetrieveProjectMember(membertype MemberType, slug, project_id string) (*Member, error) {
	req, err := http.NewRequest(
		http.MethodGet,
		fmt.Sprintf("/v3/projects/project/members/member/%s/%s?project=%s", membertype.String(), slug, project_id),
		nil)
	if err != nil {
		return nil, err
	}

	body, err := dp.makeApiRequest(req)
	if err != nil {
		return nil, err
	}

	data := new(ProjectMemberResp)
	err = json.Unmarshal(body, data)
	if err != nil {
		return nil, err
	}

	return &data.Member, nil
}

type AddProjectMemberParam struct {
	Type         string   `json:"type"`
	Slug         string   `json:"slug"`
	Role         string   `json:"role"`
	Environments []string `json:"environments"`
}

/*Add Project Member*/
func (dp *doppler) AddProjectMember(project_id string, args AddProjectMemberParam) (*Member, error) {
	payload, err := json.Marshal(args)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(
		http.MethodPost,
		fmt.Sprintf("/v3/projects/project/members?project=%s", project_id),
		bytes.NewReader(payload))
	if err != nil {
		return nil, err
	}

	body, err := dp.makeApiRequest(req)
	if err != nil {
		return nil, err
	}

	data := new(ProjectMemberResp)
	err = json.Unmarshal(body, data)
	if err != nil {
		return nil, err
	}

	return &data.Member, nil
}

type UpdateProjectMemberParams struct {
	Role         string   `json:"role"`
	Environments []string `json:"environments"`
}

/*Update Project Member*/
func (dp *doppler) UpdateProjectMember(project_id, member_slug string, member_type MemberType, args UpdateProjectMemberParams) (*Member, error) {
	payload, err := json.Marshal(args)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(
		http.MethodPatch,
		fmt.Sprintf("/v3/projects/project/members/member/%s/%s?project=%s", member_type.String(), member_slug, project_id),
		bytes.NewReader(payload))
	if err != nil {
		return nil, err
	}

	body, err := dp.makeApiRequest(req)
	if err != nil {
		return nil, err
	}

	data := new(ProjectMemberResp)
	err = json.Unmarshal(body, data)
	if err != nil {
		return nil, err
	}

	return &data.Member, nil
}

/*Remove Project Member*/
func (dp *doppler) RemoveProjectMember(member_type MemberType, member_slug, project_id string) error {
	req, err := http.NewRequest(
		http.MethodDelete,
		fmt.Sprintf("/v3/projects/project/members/member/%s/%s?project=%s", member_type.String(), member_slug, project_id),
		nil)
	if err != nil {
		return err
	}

	_, err = dp.makeApiRequest(req)
	if err != nil {
		return err
	}

	return nil
}

func (m MemberType) String() string {
	return [...]string{"workplace_user", "group", "invite", "service_account"}[m]
}

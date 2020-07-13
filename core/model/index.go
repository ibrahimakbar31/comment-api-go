package model

//Pagination struct
type Pagination struct {
	Page    int64 `json:"page" groups:"orgComments,membersOrganization"`
	PerPage int64 `json:"per_page" groups:"orgComments,membersOrganization"`
}

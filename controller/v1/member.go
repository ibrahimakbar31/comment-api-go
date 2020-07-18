package v1

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/ibrahimakbar31/comment-api-go/core/model"
	"github.com/ibrahimakbar31/comment-api-go/middleware"
)

//MembersGetResponse struct
type MembersGetResponse struct {
	Message string `json:"message" groups:"membersOrganization"`
	model.MembersPagination
}

//GetMembersByOrganization function
func GetMembersByOrganization(c *gin.Context, app *middleware.App) (interface{}, string, error) {
	var err error
	var membersPagination model.MembersPagination
	var pagination model.Pagination
	var response MembersGetResponse
	group := "membersOrganization"
	response.Message = "success"
	organization, ok := c.MustGet("organization").(model.Organization)
	if !ok {
		return response, group, errors.New("TOKEN_DATA_INVALID")
	}

	pagination, ok = c.MustGet("pagination").(model.Pagination)
	if !ok {
		return response, group, errors.New("PAGINATION_DATA_INVALID")
	}
	membersPagination.Pagination = pagination
	membersPagination.Members, err = app.DB1.GetMembersByOrganizationID(organization.ID, membersPagination.Pagination)
	if err != nil {
		if err.Error() == "record not found" {
			return response, group, nil
		}
	}
	response.MembersPagination = membersPagination
	return response, group, err
}

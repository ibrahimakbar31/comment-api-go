package v1

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/ibrahimakbar31/comment-api-go/core/model"
	"github.com/ibrahimakbar31/comment-api-go/middleware"
)

//GetMembersByOrganization function
func GetMembersByOrganization(app *middleware.App, c *gin.Context) (model.MembersPagination, error) {
	var err error
	var membersPagination model.MembersPagination
	var pagination model.Pagination

	organization, ok := c.MustGet("organization").(model.Organization)
	if !ok {
		return membersPagination, errors.New("TOKEN_DATA_INVALID")
	}

	pagination, ok = c.MustGet("pagination").(model.Pagination)
	if !ok {
		return membersPagination, errors.New("PAGINATION_DATA_INVALID")
	}
	membersPagination.Pagination = pagination
	membersPagination.Members, err = app.DB1.GetMembersByOrganizationID(organization.ID, membersPagination.Pagination)
	if err != nil {
		return membersPagination, err
	}

	return membersPagination, err
}

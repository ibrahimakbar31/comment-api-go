package v1

import (
	"errors"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/ibrahimakbar31/comment-api-go/core/model"
	"github.com/ibrahimakbar31/comment-api-go/middleware"
	"github.com/jinzhu/copier"
	uuid "github.com/satori/go.uuid"
)

//CreateComment function
func CreateComment(commentInput model.CommentCreate, app *middleware.App, c *gin.Context) (model.Comment, error) {
	var err error
	var comment model.Comment
	memberToken, ok := c.MustGet("memberToken").(middleware.MemberToken)
	if !ok {
		return comment, errors.New("TOKEN_DATA_INVALID")
	}
	organization, ok := c.MustGet("organization").(model.Organization)
	if !ok {
		return comment, errors.New("TOKEN_DATA_INVALID")
	}
	_, err = govalidator.ValidateStruct(commentInput)
	if err != nil {
		return comment, err
	}
	copier.Copy(&comment, commentInput)
	comment.MemberID = uuid.NullUUID{
		UUID:  memberToken.Member.ID,
		Valid: true,
	}
	comment.Member = memberToken.Member
	comment.OrganizationID = uuid.NullUUID{
		UUID:  organization.ID,
		Valid: true,
	}
	comment.Organization = organization

	comment, err = app.DB1.CreateComment(comment)
	if err != nil {
		return comment, err
	}

	return comment, err
}

//GetCommentsByOrganization function
func GetCommentsByOrganization(app *middleware.App, c *gin.Context) (model.CommentsPagination, error) {
	var err error
	var commentsPagination model.CommentsPagination
	var pagination model.Pagination
	organization, ok := c.MustGet("organization").(model.Organization)
	if !ok {
		return commentsPagination, errors.New("TOKEN_DATA_INVALID")
	}
	pagination, ok = c.MustGet("pagination").(model.Pagination)
	if !ok {
		return commentsPagination, errors.New("PAGINATION_DATA_INVALID")
	}
	commentsPagination.Pagination = pagination
	commentsPagination.Comments, err = app.DB1.GetCommentsByOrganizationID(organization.ID, commentsPagination.Pagination)
	if err != nil {
		return commentsPagination, err
	}

	return commentsPagination, err
}

//DeleteCommentsByOrganization function
func DeleteCommentsByOrganization(app *middleware.App, c *gin.Context) (int64, error) {
	var err error
	var deleteCount int64
	organization, ok := c.MustGet("organization").(model.Organization)
	if !ok {
		return deleteCount, errors.New("TOKEN_DATA_INVALID")
	}
	deleteCount, err = app.DB1.DeleteCommentsByOrganizationID(organization.ID)
	if err != nil {
		return deleteCount, err
	}
	return deleteCount, err
}

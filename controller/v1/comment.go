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

//CommentCreateResponse struct
type CommentCreateResponse struct {
	Message string        `json:"message" groups:"comment"`
	Comment model.Comment `json:"comment" groups:"comment"`
}

//CommentsGetResponse struct
type CommentsGetResponse struct {
	Message string `json:"message" groups:"orgComments"`
	model.CommentsPagination
}

//CommentsDeleteResponse struct
type CommentsDeleteResponse struct {
	Message     string `json:"message" groups:"deleteComments"`
	DeleteCount int64  `json:"delete_count" groups:"deleteComments"`
}

//CreateComment function
func CreateComment(c *gin.Context, app *middleware.App) (interface{}, string, error) {
	var err error
	var commentInput model.CommentCreate
	response := CommentCreateResponse{
		Message: "success",
	}
	group := "comment"
	memberToken, ok := c.MustGet("memberToken").(middleware.MemberToken)
	if !ok {
		return response, group, errors.New("TOKEN_DATA_INVALID")
	}
	organization, ok := c.MustGet("organization").(model.Organization)
	if !ok {
		return response, group, errors.New("TOKEN_DATA_INVALID")
	}
	c.ShouldBindJSON(&commentInput)
	_, err = govalidator.ValidateStruct(commentInput)
	if err != nil {
		errs := err.(govalidator.Errors).Errors()
		return response, group, errs[0]
	}
	copier.Copy(&response.Comment, commentInput)
	response.Comment.MemberID = uuid.NullUUID{
		UUID:  memberToken.Member.ID,
		Valid: true,
	}
	response.Comment.Member = memberToken.Member
	response.Comment.OrganizationID = uuid.NullUUID{
		UUID:  organization.ID,
		Valid: true,
	}
	response.Comment.Organization = organization
	err = app.DB1.CreateComment(&response.Comment)
	if err != nil {
		return response, group, err
	}
	return response, group, err
}

//GetCommentsByOrganization function
func GetCommentsByOrganization(c *gin.Context, app *middleware.App) (interface{}, string, error) {
	var err error
	var commentsPagination model.CommentsPagination
	response := CommentsGetResponse{
		Message: "success",
	}
	group := "orgComments"
	organization, ok := c.MustGet("organization").(model.Organization)
	if !ok {
		return response, group, errors.New("TOKEN_DATA_INVALID")
	}
	response.CommentsPagination.Pagination, ok = c.MustGet("pagination").(model.Pagination)
	if !ok {
		return response, group, errors.New("PAGINATION_DATA_INVALID")
	}
	response.CommentsPagination.Comments, err = app.DB1.GetCommentsByOrganizationID(organization.ID, commentsPagination.Pagination)
	if err != nil {
		if err.Error() == "record not found" {
			return response, group, nil
		}
	}
	return response, group, err
}

//DeleteCommentsByOrganization function
func DeleteCommentsByOrganization(c *gin.Context, app *middleware.App) (interface{}, string, error) {
	var err error
	response := CommentsDeleteResponse{
		Message: "success",
	}
	group := "deleteComments"
	organization, ok := c.MustGet("organization").(model.Organization)
	if !ok {
		return response, group, errors.New("TOKEN_DATA_INVALID")
	}
	response.DeleteCount, err = app.DB1.DeleteCommentsByOrganizationID(organization.ID)
	if err != nil {
		return response, group, err
	}
	return response, group, err
}

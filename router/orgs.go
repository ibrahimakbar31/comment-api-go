package router

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/ibrahimakbar31/comment-api-go/controller/v1"
	"github.com/ibrahimakbar31/comment-api-go/core/model"
	"github.com/ibrahimakbar31/comment-api-go/middleware"
)

//CommentPostV1Output struct
type CommentPostV1Output struct {
	Message string        `json:"message" groups:"comment"`
	Comment model.Comment `json:"comment" groups:"comment"`
}

//CommentsGetV1Output struct
type CommentsGetV1Output struct {
	Message string `json:"message" groups:"orgComments"`
	model.CommentsPagination
}

//CommentsDeleteV1Output struct
type CommentsDeleteV1Output struct {
	Message     string `json:"message" groups:"deleteComments"`
	DeleteCount int64  `json:"delete_count" groups:"deleteComments"`
}

//MembersGetV1Output struct
type MembersGetV1Output struct {
	Message string `json:"message" groups:"membersOrganization"`
	model.MembersPagination
}

//OrgsV1 router
func OrgsV1(routeGroup *gin.RouterGroup, app *middleware.App) {
	orgsV1 := routeGroup.Group("/orgs/:orgCode")
	orgsV1.Use(app.UserValidate("auth"))
	orgsCommentsV1 := orgsV1.Group("/comments")
	orgsCommentsV1.Use(app.UserValidate("memberInOrganization"))
	orgsCommentsV1.POST("/", func(c *gin.Context) {
		var commentCreate model.CommentCreate
		c.ShouldBindJSON(&commentCreate)
		comment, err := v1.CreateComment(commentCreate, app, c)
		output := CommentPostV1Output{
			Message: "success",
			Comment: comment,
		}
		var outputSend = middleware.Output{
			Ctx:        c,
			Err:        err,
			Group:      "comment",
			StructData: output,
			App:        app,
		}
		outputSend.Handler()
	})
	orgsCommentsV1.GET("/", app.InputValidation("getPaginationQuery"), func(c *gin.Context) {
		commentsPagination, err := v1.GetCommentsByOrganization(app, c)
		output := CommentsGetV1Output{
			Message:            "success",
			CommentsPagination: commentsPagination,
		}
		var outputSend = middleware.Output{
			Ctx:        c,
			Err:        err,
			Group:      "orgComments",
			StructData: output,
			App:        app,
		}
		outputSend.Handler()
	})
	orgsCommentsV1.DELETE("/", func(c *gin.Context) {
		deleteCount, err := v1.DeleteCommentByOrganization(app, c)
		output := CommentsDeleteV1Output{
			Message:     "success",
			DeleteCount: deleteCount,
		}
		var outputSend = middleware.Output{
			Ctx:        c,
			Err:        err,
			Group:      "deleteComments",
			StructData: output,
			App:        app,
		}
		outputSend.Handler()
	})

	orgsMembersV1 := orgsV1.Group("/members")
	orgsMembersV1.Use(app.UserValidate("memberInOrganization"))
	orgsMembersV1.GET("/", app.InputValidation("getPaginationQuery"), func(c *gin.Context) {
		membersPagination, err := v1.GetMembersByOrganization(app, c)
		output := MembersGetV1Output{
			Message:           "success",
			MembersPagination: membersPagination,
		}
		var outputSend = middleware.Output{
			Ctx:        c,
			Err:        err,
			Group:      "membersOrganization",
			StructData: output,
			App:        app,
		}
		outputSend.Handler()
	})

}

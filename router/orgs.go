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

//Orgs router
func Orgs(routeGroup *gin.RouterGroup, app *middleware.App) {
	Orgs := routeGroup.Group("/orgs/:orgCode")
	Orgs.Use(app.ValidateUser("auth"))
	orgComments := Orgs.Group("/comments")
	orgComments.Use(app.ValidateUser("memberInOrganization"))
	orgComments.POST("/", func(c *gin.Context) {
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
	orgComments.GET("/", app.ValidateInput("getPaginationQuery"), func(c *gin.Context) {
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
	orgComments.DELETE("/", func(c *gin.Context) {
		deleteCount, err := v1.DeleteCommentsByOrganization(app, c)
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

	orgMembers := Orgs.Group("/members")
	orgMembers.Use(app.ValidateUser("memberInOrganization"))
	orgMembers.GET("/", app.ValidateInput("getPaginationQuery"), func(c *gin.Context) {
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

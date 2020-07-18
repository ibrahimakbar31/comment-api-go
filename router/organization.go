package router

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/ibrahimakbar31/comment-api-go/controller/v1"
	"github.com/ibrahimakbar31/comment-api-go/middleware"
)

//Orgs router
func Orgs(routeGroup *gin.RouterGroup, app *middleware.App) {
	Orgs := routeGroup.Group("/orgs/:orgCode")
	Orgs.Use(app.ValidateUser("auth"), app.ValidateUser("memberInOrganization"))

	orgComments := Orgs.Group("/comments")
	orgComments.POST("/", app.ResponseHandler(v1.CreateComment))
	orgComments.GET("/", app.ValidateInput("getPaginationQuery"), app.ResponseHandler(v1.GetCommentsByOrganization))
	orgComments.DELETE("/", app.ResponseHandler(v1.DeleteCommentsByOrganization))

	orgMembers := Orgs.Group("/members")
	orgMembers.GET("/", app.ValidateInput("getPaginationQuery"), app.ResponseHandler(v1.GetMembersByOrganization))
}

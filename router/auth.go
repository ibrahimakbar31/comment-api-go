package router

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/ibrahimakbar31/comment-api-go/controller/v1"
	"github.com/ibrahimakbar31/comment-api-go/middleware"
)

//Auth router
func Auth(routeGroup *gin.RouterGroup, app *middleware.App) {
	Auth := routeGroup.Group("/auth")
	Auth.POST("/login", app.ResponseHandler(v1.SubmitLogin))
}

package router

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/ibrahimakbar31/comment-api-go/controller/v1"
	"github.com/ibrahimakbar31/comment-api-go/middleware"
)

//AuthLoginV1Output struct
type AuthLoginV1Output struct {
	Message string           `json:"message" groups:"login"`
	Token   middleware.Token `json:"token" groups:"login"`
}

//AuthV1 router
func AuthV1(routeGroup *gin.RouterGroup, app *middleware.App) {
	authV1 := routeGroup.Group("/auth")
	authV1.POST("/login", func(c *gin.Context) {
		var err error
		var output AuthLoginV1Output
		var loginCredential v1.LoginCredential
		c.ShouldBindJSON(&loginCredential)
		memberToken, err := v1.LoginSubmit(loginCredential, app)
		output.Message = "success"
		output.Token = memberToken.Token
		var outputSend = middleware.Output{
			Ctx:        c,
			Err:        err,
			Group:      "login",
			StructData: output,
			App:        app,
		}
		outputSend.Handler()
	})
}

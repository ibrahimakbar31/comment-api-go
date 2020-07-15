package router

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/ibrahimakbar31/comment-api-go/controller/v1"
	"github.com/ibrahimakbar31/comment-api-go/middleware"
)

//AuthLoginResponse struct
type AuthLoginResponse struct {
	Message string           `json:"message" groups:"login"`
	Token   middleware.Token `json:"token" groups:"login"`
}

//Auth router
func Auth(routeGroup *gin.RouterGroup, app *middleware.App) {
	Auth := routeGroup.Group("/auth")
	Auth.POST("/login", func(c *gin.Context) {
		var err error
		var output AuthLoginResponse
		var loginCredential v1.LoginCredential
		c.ShouldBindJSON(&loginCredential)
		memberToken, err := v1.SubmitLogin(loginCredential, app)
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

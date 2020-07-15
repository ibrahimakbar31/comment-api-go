package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ibrahimakbar31/comment-api-go/core/model"
	"github.com/ibrahimakbar31/comment-api-go/middleware"
)

//InitAllRoutes for initial route V1
func InitAllRoutes(app *middleware.App) {
	app.Router = gin.New()
	app.Router.NoRoute(HandleNoRoute())
	version1 := app.Router.Group("/v1")
	InitRouterV1(version1, app)
}

//HandleNoRoute function to handle if no match route
func HandleNoRoute() gin.HandlerFunc {
	return func(c *gin.Context) {
		var err model.APIError
		err.ID = 1
		err.Code = 404
		err.Name = "PAGE_NOT_FOUND"
		err.Message = "Page Not Found"
		var marshalData, _ = middleware.GenerateMarshal([]string{"error"}, err)
		c.JSON(err.Code, marshalData)
	}
}

package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ibrahimakbar31/comment-api-go/middleware"
)

//InitRouterV1 for initial route V1
func InitRouterV1(routeGroup *gin.RouterGroup, app *middleware.App) {
	routeGroup.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "success"})
	})
	Auth(routeGroup, app)
	Orgs(routeGroup, app)
}

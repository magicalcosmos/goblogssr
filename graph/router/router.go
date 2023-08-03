package router

import (
	"github.com/magicalcosmos/goblogssr/graph/controller"

	"github.com/gin-gonic/gin"
)

// Init init
func Init(r *gin.Engine) {
	r.POST("/api/query", controller.GraphqlHandler())
	r.GET("/api", controller.PlaygroundHandler())
}

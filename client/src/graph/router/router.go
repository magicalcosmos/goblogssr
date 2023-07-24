package router

import (
	"github.com/gin-gonic/gin"
	"github.com/magicalcosmos/goblogssr/graph/controller"
)

// Init init
func Init(r *gin.Engine) {
	r.POST("/api/query", controller.GraphqlHandler())
	r.GET("/api", controller.PlaygroundHandler())
}

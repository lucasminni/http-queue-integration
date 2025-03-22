package v1

import (
	orderV1 "http-gateway/pkg/routes/v1/order"

	"github.com/gin-gonic/gin"
)

func Grouper(e *gin.Engine) {
	api := e.Group("/api")
	v1 := api.Group("/v1")

	Register(v1)
}

func Register(g *gin.RouterGroup) {
	orderV1.Register(g)
}

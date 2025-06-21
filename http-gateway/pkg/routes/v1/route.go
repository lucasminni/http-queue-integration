package v1

import (
	"http-gateway/pkg/routes"
	orderV1 "http-gateway/pkg/routes/v1/order"

	"github.com/gin-gonic/gin"
)

func Grouper(r *gin.Engine) {
	r.GET("/ping", routes.HealthCheck)

	api := r.Group("/api")
	v1 := api.Group("/v1")

	Register(v1)
}

func Register(g *gin.RouterGroup) {
	orderV1.Register(g)
}

package v1

import (
	itemV1 "http-gateway/pkg/routes/v1/item"
	orderV1 "http-gateway/pkg/routes/v1/order"

	"github.com/gin-gonic/gin"
)

func Register(g *gin.RouterGroup) {
	orderV1.Register(g)
	itemV1.Register(g)
}

package v1

import (
	"http-gateway/pkg/routes/v1/item"
	"http-gateway/pkg/routes/v1/order"

	"github.com/gin-gonic/gin"
)

func Register(r *gin.RouterGroup) {
	order.Register(r)
	item.Register(r)
}

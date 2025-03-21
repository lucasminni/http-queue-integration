package v1

import (
	"github.com/gin-gonic/gin"
)

func Grouper(r *gin.Engine) {
	api := r.Group("/api")
	v1 := api.Group("/v1")

	Register(v1)
}

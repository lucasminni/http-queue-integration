package v1

import (
	"github.com/gin-gonic/gin"
)

func Grouper(e *gin.Engine) {
	api := e.Group("/api")
	v1 := api.Group("/v1")

	Register(v1)
}

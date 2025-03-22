package infra

import (
	v1 "http-gateway/pkg/routes/v1"

	"github.com/gin-gonic/gin"
)

func StartHTTPServer() error {
	r := gin.Default()

	v1.Grouper(r)

	if err := r.Run(); err != nil {
		return nil
	} else {
		return err
	}
}

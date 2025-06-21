package infra

import (
	"fmt"
	"http-gateway/internal/settings"
	v1 "http-gateway/pkg/routes/v1"
	"log"

	"github.com/gin-gonic/gin"
)

func StartHTTPServer() error {
	r := gin.Default()

	v1.Grouper(r)

	port := settings.GetEnvs().HTTPServerPort

	if err := r.Run(fmt.Sprintf(":%s", port)); err != nil {
		log.Fatal("Error on server start: ", err)
	}

	return nil
}

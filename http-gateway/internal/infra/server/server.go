package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func StartHTTPServer() error {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hello",
		})
	})

	if err := r.Run(); err != nil {
		return nil
	} else {
		return err
	}
}

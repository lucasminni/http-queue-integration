package v1

import (
	item "http-gateway/internal/domain/models/item"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(g *gin.RouterGroup) {
	g.POST("/item", create)
	g.GET("/item", list)
	g.PUT("/item", update)
	g.DELETE("/item/:id", delete)
}

func create(ctx *gin.Context) {
	item := &item.Item{}

	err := ctx.ShouldBindJSON(item)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "error trying to create a new item - " + err.Error(),
		})
		log.Println("Binding JSON error - " + err.Error())
	} else {
		// create on database
		ctx.JSON(http.StatusCreated, item)
		log.Println("New item created")
	}
}

func list(ctx *gin.Context) {
	items := []item.Item{}

	ctx.JSON(http.StatusOK, items)
}

func update(ctx *gin.Context) {
	item := &item.Item{}

	err := ctx.ShouldBindJSON(item)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Unable to update the item id " + item.ID.String() + " due to error " + err.Error(),
		})
		log.Println("Binding JSON error - " + err.Error())
	} else {
		ctx.JSON(http.StatusOK, item)
	}
}

func delete(ctx *gin.Context) {
	id := ctx.Param("id")

	// delete on database
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Item id " + id + " sucessfully deleted",
	})
}

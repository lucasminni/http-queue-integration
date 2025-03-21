package item

import (
	item "http-gateway/internal/domain/models/item"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(r *gin.RouterGroup) {
	r.POST("/item", create)
	r.GET("/item", list)
	r.PUT("/item", update)
	r.DELETE("/item/:id", delete)
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

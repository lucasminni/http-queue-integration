package order

import (
	order "http-gateway/internal/domain/models/order"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(r *gin.RouterGroup) {
	r.POST("/order", create)
	r.GET("/order", list)
	r.PUT("/order", update)
	r.DELETE("/order/:id", delete)
}

func create(ctx *gin.Context) {
	order := &order.Order{}

	err := ctx.ShouldBindJSON(order)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "error trying to create a new order - " + err.Error(),
		})
		log.Println("Binding JSON error - " + err.Error())
	} else {
		// create on database
		ctx.JSON(http.StatusCreated, order)
		log.Println("New order created")
	}
}

func list(ctx *gin.Context) {
	orders := []order.Order{}

	ctx.JSON(http.StatusOK, orders)
}

func update(ctx *gin.Context) {
	order := &order.Order{}

	err := ctx.ShouldBindJSON(order)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Unable to update the order id " + order.ID.String() + " due to error " + err.Error(),
		})
		log.Println("Binding JSON error - " + err.Error())
	} else {
		ctx.JSON(http.StatusOK, order)
	}
}

func delete(ctx *gin.Context) {
	id := ctx.Param("id")

	// delete on database
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Order id " + id + " sucessfully deleted",
	})
}

package v1

import (
	model "http-gateway/internal/domain/models/order"
	schema "http-gateway/internal/domain/schemas/order"
	service "http-gateway/internal/domain/services"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(g *gin.RouterGroup) {
	g.POST("/order", create)
	g.GET("/order", list)
	g.PUT("/order", update)
	g.DELETE("/order/:id", delete)
}

func create(ctx *gin.Context) {
	json := &model.Order{}

	err := ctx.ShouldBindJSON(json)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "error trying to create a new order - " + err.Error(),
		})
		log.Println("Binding JSON error - " + err.Error())
	}

	order, err := service.CreateOrder(*json)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "error trying to create a new order - " + err.Error(),
		})
		log.Println("Binding JSON error - " + err.Error())
	} else {
		ctx.JSON(http.StatusCreated, order)
	}
}

func list(ctx *gin.Context) {
	orders, err := service.GetOrders()

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "error trying to list the orders - " + err.Error(),
		})
		log.Println("Listing orders error - " + err.Error())
	} else {
		ctx.JSON(http.StatusOK, orders)
	}
}

func update(ctx *gin.Context) {
	json := &schema.BodyUpdateOrder{}

	err := ctx.ShouldBindJSON(json)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Unable to update the order due to error " + err.Error(),
		})
		log.Println("Binding JSON error - " + err.Error())
	}

	order, err := service.UpdateOrder(*json)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Error trying to update -" + err.Error(),
		})
		log.Println("Update order error - " + err.Error())
	} else {
		ctx.JSON(http.StatusOK, order)
	}
}

func delete(ctx *gin.Context) {
	id := ctx.Param("id")

	err := service.DeleteOrder(id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Error trying to delete order - " + err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Order id " + id + " sucessfully deleted",
		})
	}
}

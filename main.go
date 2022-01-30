package main

import (
	"customer-orders/controllers"
	"customer-orders/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()

	r.PUT("/orders/:orderid", controllers.UpdateOrder)
	r.POST("/orders", controllers.CreateOrder)
	r.GET("/orders", controllers.FindOrders)
	r.DELETE("/orders/:orderid", controllers.DeleteOrderByID)
	r.GET("/orders/:orderid", controllers.FindOrderByID)

	r.Run()
}

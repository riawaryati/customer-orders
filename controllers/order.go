package controllers

import (
	"customer-orders/models"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func FindOrders(c *gin.Context) {
	orders := []models.Order{}
	// models.DB.Joins("JOIN items ON items.order_id = orders.order_id").Find(&orders)
	models.DB.Find(&orders)

	retOrders := []models.Order{}
	for _, order := range orders {
		items := []models.Item{}
		models.DB.Where("order_id = ?", order.OrderID).Find(&items)
		order.Items = items
		retOrders = append(retOrders, order)
	}

	c.JSON(http.StatusOK, retOrders)
}

func CreateOrder(c *gin.Context) {
	var request models.OrderRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println(request)

	orderedAt, err := time.Parse(time.RFC3339, request.OrderedAt)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	order := &models.Order{CustomerName: request.CustomerName, OrderedAt: orderedAt}
	if err := models.DB.Create(order).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	fmt.Println(request.Items)
	for _, item := range request.Items {
		if err := models.DB.Create(&models.Item{OrderID: order.OrderID, ItemCode: item.ItemCode, Description: item.Description, Quantity: item.Quantity}).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

	}

	c.JSON(http.StatusOK, gin.H{"data": "success"})
}

func UpdateOrder(c *gin.Context) {
	var request models.OrderRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	order := models.Order{}

	if err := models.DB.Where("order_id = ?", c.Param("orderid")).First(&order).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Order not found!"})
		return
	}

	orderedAt, err := time.Parse(time.RFC3339, request.OrderedAt)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	if err := models.DB.Model(&order).Updates(&models.Order{CustomerName: request.CustomerName, OrderedAt: orderedAt}).Error; err != nil {
		// panic(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for _, item := range request.Items {
		oldItem := models.Item{}

		if err := models.DB.Where("order_id = ? AND item_id = ?", c.Param("orderid"), item.ItemID).First(&oldItem).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Order Item not found!"})
			return
		}

		if err = models.DB.Model(&oldItem).Updates(&models.Item{OrderID: oldItem.OrderID, ItemID: oldItem.ItemID, ItemCode: item.ItemCode, Description: item.Description, Quantity: item.Quantity}).Error; err != nil {
			panic(err)
		}
	}

	c.JSON(http.StatusOK, gin.H{"data": "success"})
}

func DeleteOrderByID(c *gin.Context) {
	order := models.Order{}
	items := []models.Item{}

	if err := models.DB.Where("order_id = ?", c.Param("orderid")).First(&items).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Item not found!"})
		return
	}

	if err := models.DB.Where("order_id = ?", c.Param("orderid")).First(&order).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Order not found!"})
		return
	}

	fmt.Println(&items)

	models.DB.Delete(models.Item{}, "order_id = ?", c.Param("orderid"))
	models.DB.Delete(&order)

	c.JSON(http.StatusOK, gin.H{"data": "success"})
}

func FindOrderByID(c *gin.Context) {
	order := models.Order{}

	if err := models.DB.Where("order_id = ?", c.Param("orderid")).First(&order).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	items := []models.Item{}
	models.DB.Where("order_id = ?", order.OrderID).Find(&items)
	order.Items = items

	c.JSON(http.StatusOK, gin.H{"data": order})
}

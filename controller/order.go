package controller

import (
	"net/http"

	"online-ordering-system/model"

	"github.com/gin-gonic/gin"
)

var (
	order  model.Order
	review model.Review
)

func (p *Controller) CreateOrderHandler(c *gin.Context) {

	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return

	}
	p.md.CreateOrder(order)

	c.JSON(http.StatusOK, gin.H{"message": "Order created successfully"})
}

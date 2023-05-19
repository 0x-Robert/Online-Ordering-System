package controller

import (
	"net/http"

	"online-ordering-system/types"

	"github.com/gin-gonic/gin"
)

var (
	order  types.Order
	review types.Review
)

func (p *Controller) CreateOrderHandler(c *gin.Context) {

	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return

	}
	p.md.CreateOrder(order)

	c.JSON(http.StatusOK, gin.H{"message": "Order created successfully"})
}

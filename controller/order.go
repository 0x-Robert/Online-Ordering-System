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

// 특정 주문을 만드는 함수
// CreateOrderHandler godoc
// @Summary
// @Tags Crete Order
// @Description 특정 메뉴를 주문한다. 주문데이터는 model.Order를 참고한다.
// @name CreateOrderHandler
// @Accept  json
// @Produce  json
// @Param menuname query string true "menuname"
// @Param customer query string true "customer"
// @Param phonenumber query string true "phonenumber"
// @Param address query string true "address"
// @Param quantity query string true "quantity"
// @Param paymentinformation query string true "paymentinformation"
// @Router /v01/user/order [post]
// @Success 200 {array} model.Order
func (p *Controller) CreateOrderHandler(c *gin.Context) {

	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return

	}
	p.md.CreateOrder(order)

	c.JSON(http.StatusOK, gin.H{"message": "Order created successfully"})
}

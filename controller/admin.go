package controller

import (
	"net/http"
	model "online-ordering-system/model"

	"github.com/gin-gonic/gin"
)

var (
	admin model.Admin
)

func (p *Controller) AdminRegisterHandler(c *gin.Context) {

	if err := c.ShouldBindJSON(&admin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return

	}
	p.md.RegisterAdmin(admin)
	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

func (p *Controller) AdminLoginHandler(c *gin.Context) {

	if err := c.ShouldBindJSON(&admin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return

	}

	if p.md.LoginAdmin(admin) {
		c.JSON(http.StatusOK, gin.H{"message": "User is logged in"})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check user login"})

	}

}

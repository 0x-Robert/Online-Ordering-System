package controller

import (
	"net/http"
	"online-ordering-system/types"

	"github.com/gin-gonic/gin"
)

var (
	user types.User
)

func (p *Controller) UserRegisterHandler(c *gin.Context) {

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return

	}
	p.md.RegisterUser(user)
	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

func (p *Controller) LoginHandler(c *gin.Context) {

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return

	}

	if p.md.LoginUser(user) {
		c.JSON(http.StatusOK, gin.H{"message": "User is logged in"})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check user login"})

	}

}

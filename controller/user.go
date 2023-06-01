package controller

import (
	"net/http"
	"online-ordering-system/model"

	"github.com/gin-gonic/gin"
)

var (
	user model.User
)

// 유저를 등록하는 함수
// UserRegisterHandler godoc
// @Summary
// @Tags Register user
// @Description 유저를 등록하는 함수다.
// @name UserRegisterHandler
// @Accept  json
// @Produce  json
// @Param id query string true "id"
// @Param password query string true "password"
// @Router /v01/user/register [post]
// @Success 200 {array} model.User
func (p *Controller) UserRegisterHandler(c *gin.Context) {

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return

	}
	p.md.RegisterUser(user)
	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

// 로그인하는 함수
// LoginHandler godoc
// @Summary
// @Tags Login User
// @Description 로그인하는 함수
// @name LoginHandler
// @Accept  json
// @Produce  json
// @Param id query string true "id"
// @Param password query string true "password"
// @Router /v01/user/login [post]
// @Success 200 {array} model.User
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

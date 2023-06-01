package controller

import (
	"net/http"
	model "online-ordering-system/model"

	"github.com/gin-gonic/gin"
)

var (
	admin model.Admin
)

// 관리자 등록하는 함수
// AdminRegisterHandler godoc
// @Summary
// @Tags Admin Register
// @Description 관리자를 등록하기 위한 함수
// @name AdminRegisterHandler
// @Accept  json
// @Produce  json
// @Param id 	   query string true "id"
// @Param password query string true "password"
// @Router /v01/admin/register [post]
// @Success 200 {array} model.Admin
func (p *Controller) AdminRegisterHandler(c *gin.Context) {

	if err := c.ShouldBindJSON(&admin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return

	}
	p.md.RegisterAdmin(admin)
	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

// 관리자가 로그인하는 함수
// AdminLoginHandler godoc
// @Summary
// @Tags Admin Login
// @Description 관리자가 로그인 하기 위한 함수
// @name AdminLoginHandler
// @Accept  json
// @Produce  json
// @Param id query string true "id"
// @Param password query string true "password"
// @Router /v01/admin/login [post]
// @Success 200 {array} model.Admin
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

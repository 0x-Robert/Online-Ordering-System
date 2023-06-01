package controller

import (
	"fmt"
	"log"
	"net/http"
	model "online-ordering-system/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

var menu model.Menu

// 메뉴를 만드는 함수
// CreateMenuHandler godoc
// @Summary
// @Tags Create Menu
// @Description 메뉴를 만드는 함수
// @name CreateMenuHandler
// @Accept  json
// @Produce  json
// @Param menuid query string true "menuid"
// @Param imageurl query string true "imageurl"
// @Param name query string true "name"
// @Param quantity query string true "quantity"
// @Param price query string true "price"
// @Param recommendation query string true "recommendation"
// @Param admin query string true "admin"
// @Param score query string true "score"
// @Param review query string true "review"
// @Router /v01/admin/menu/create [post]
// @Success 200 {array} model.Menu
func (p *Controller) CreateMenuHandler(c *gin.Context) {

	if err := c.ShouldBindJSON(&menu); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return

	}

	p.md.CreateMenu(menu)
	c.JSON(http.StatusOK, gin.H{"message": "Menu created successfully"})
}

// 메뉴를 삭제하는 함수
// DeleteMenuHandler godoc
// @Summary
// @Tags Delete Menu
// @Description 메뉴를 삭제하는 함수
// @name DeleteMenuHandler
// @Accept  json
// @Produce  json
// @Param menuid query string true "menuid"
// @Router /v01/admin/menu/delete [post]
// @Success 200 {array} model.Menu
func (p *Controller) DeleteMenuHandler(c *gin.Context) {
	if err := c.ShouldBindJSON(&menu); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return

	}

	p.md.DeleteMenu(menu)
	c.JSON(http.StatusOK, gin.H{"message": "Menu deleted successfully"})
}

// 메뉴를 전부 가져오는 함수
// GetAllMenuHandler godoc
// @Summary
// @Tags Get All Menu
// @Description 메뉴를 전부 가져오는 함수
// @name GetAllMenuHandler
// @Accept  json
// @Produce  json
// @Router /v01/admin/menu/status [get]
// @Success 200 {array} model.Menu
func (p *Controller) GetAllMenuHandler(c *gin.Context) {

	result := p.md.GetAllMenu()

	c.JSON(http.StatusOK, result)
}

// 특정 메뉴만 가져오는 함수 디테일 메뉴
// DetailMenuHandler godoc
// @Summary
// @Tags Get  Menu Detail
// @Description 메뉴를 한개만 가져온다. 해당하는 메뉴만 가져온다.
// @name DetailMenuHandler
// @Accept  json
// @Produce  json
// @Param menuid query string true "menuid"
// @Router /v01/user/menu/detail [get]
// @Success 200 {array} model.Menu
func (p *Controller) DetailMenuHandler(c *gin.Context) {

	menuID, err := strconv.Atoi(c.Query("menuid")) // URL에서 menuId 값을 가져옴
	if err != nil {
		fmt.Errorf(err.Error())
		log.Println(err)
	}

	menuDetail := p.md.DetailMenu(menuID)

	c.JSON(http.StatusOK, menuDetail)
}

// 특정 메뉴를 추천하는 함수
// RecommendHandler godoc
// @Summary
// @Tags Recommend Order
// @Description 특정 메뉴를 추천해주는 함수다.
// @name RecommendHandler
// @Accept  json
// @Produce  json
// @Param recommendation query string true "recommendation"
// @Router /v01/admin/menu/recom [post]
// @Success 200 {array} model.Menu
func (p *Controller) RecommendHandler(c *gin.Context) {

	// 요청된 JSON 데이터 파싱
	if err := c.ShouldBindJSON(&menu); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse JSON data"})
		return
	}

	resultCount := p.md.RecommendMenu(menu)

	c.JSON(http.StatusOK, resultCount)

}

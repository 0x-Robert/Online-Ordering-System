package controller

import (
	"fmt"
	"log"
	"net/http"
	"online-ordering-system/types"
	"strconv"

	"github.com/gin-gonic/gin"
)

var menu types.Menu

func (p *Controller) CreateMenuHandler(c *gin.Context) {

	if err := c.ShouldBindJSON(&menu); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return

	}

	p.md.CreateMenu(menu)
	c.JSON(http.StatusOK, gin.H{"message": "Menu created successfully"})
}

func (p *Controller) DeleteMenuHandler(c *gin.Context) {
	if err := c.ShouldBindJSON(&menu); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return

	}

	p.md.DeleteMenu(menu)
	c.JSON(http.StatusOK, gin.H{"message": "Menu deleted successfully"})
}

func (p *Controller) GetAllMenuHandler(c *gin.Context) {

	result := p.md.GetAllMenu()

	c.JSON(http.StatusOK, result)
}

func (p *Controller) DetailMenuHandler(c *gin.Context) {

	menuID, err := strconv.Atoi(c.Query("menuid")) // URL에서 menuId 값을 가져옴
	if err != nil {
		fmt.Errorf(err.Error())
		log.Println(err)
	}

	menuDetail := p.md.DetailMenu(menuID)

	c.JSON(http.StatusOK, menuDetail)
}

func (p *Controller) RecommendHandler(c *gin.Context) {

	// 요청된 JSON 데이터 파싱
	if err := c.ShouldBindJSON(&menu); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse JSON data"})
		return
	}

	resultCount := p.md.RecommendMenu(menu)

	c.JSON(http.StatusOK, resultCount)

}

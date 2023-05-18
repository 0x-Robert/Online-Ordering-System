package controller

import (
	ctl "online-ordering-system/controller"
	"testing"
)

type Router struct {
	ct *ctl.Controller
}

func TestSolution(t *testing.T) {
	controller := ctl.NewController()

	// Router 초기화
	router := Router{
		ct: controller,
	}

	menu := Menu{
		MenuId:         4,
		ImageUrl:       "https://example.com/image.jpg",
		Name:           "Sample Menu",
		Quantity:       10,
		Price:          5000,
		Recommendation: false,
		Admin:          "admin",
	}
	//테스트코드 공부 필요!!
	router.ct.CreateMenuHandler(menu)
}

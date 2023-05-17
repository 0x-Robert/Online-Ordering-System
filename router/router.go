package router

import (
	"fmt"
	ctl "online-ordering-system/controller"

	"github.com/gin-gonic/gin"
)

type Router struct {
	ct *ctl.Controller
}

func NewRouter(ctl *ctl.Controller) (*Router, error) {
	r := &Router{ct: ctl} //controller 포인터를 ct로 복사, 할당

	return r, nil
}

// cross domain을 위해 사용
func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		//~ 생략
		c.Next()
	}
}

// 임의 인증을 위한 함수
func LiteAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		//~ 생략
		c.Next()
	}
}

// 실제 라우팅
func (p *Router) Idx() *gin.Engine {
	//~생략
	e := gin.New()
	account := e.Group("acc/v01", LiteAuth())
	{
		fmt.Println(account)
		account.GET("/ok", p.ct.GetOK) // controller 패키지의 실제 처리 함수
	}


	



	return e
}

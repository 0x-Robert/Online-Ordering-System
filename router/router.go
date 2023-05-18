package router

import (
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

	// e := gin.New()
	// account := e.Group("admin/v01", LiteAuth())
	// {
	// 	fmt.Println(account)
	// 	//account.GET("/ok", p.ct.GetOK) // controller 패키지의 실제 처리 함수
	// 	// account.POST("/register", p.ct.Register)
	// 	account.POST("/register", p.ct.Register)
	// }

	router := gin.Default()

	//어드민 레지스터
	router.POST("/admin/v01/register", p.ct.RegisterHandler)
	//어드민 로그인
	router.POST("/admin/v01/login", p.ct.LoginHandler)
	//메뉴 생성
	router.POST("/admin/v01/menu/create", p.ct.CreateMenuHandler)
	//메뉴 삭제
	router.POST("/admin/v01/menu/delete", p.ct.DeleteMenuHandler)
	//메뉴 추천, 비추천
	router.POST("/admin/v01/menu/recom", p.ct.RecommendHandler)
	//menu 상태 보고 전체 내역 가져오기
	router.GET("/admin/v01/menu/status", p.ct.MenuStatusHandler)
	//매뉴 디테일
	router.GET("/v01/menu/detail", p.ct.DetailMenuHandler)
	//유저 회원가입
	router.POST("/v01/register", p.ct.UserRegisterHandler)
	//유저 로그인
	router.POST("/v01/login", p.ct.UserLoginHandler)
	//주문 넣기
	router.POST("/v01/order", p.ct.CreateOrderHandler)

	//리뷰 남기기
	router.POST("/v01/order/review", p.ct.CreateOrderReviewHandler)

	//주문 수정시 상태가 조리중 or 배달 중일 경우 실패
	//주문 수정
	//router.POST("/v01/order/edit", p.ct.CreateOrderHandler)

	//추천이 많은 것을 기준으로  필터링

	//평점으로 내림차순 식으로 필터링

	//주문 수로 필터링

	//최신 날짜 기준으로 필터링 >> 메뉴에 날짜 추가?

	return router
}

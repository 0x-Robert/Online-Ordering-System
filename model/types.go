package model

type Admin struct {
	ID       string `json:"id"`
	Password string `json:"password"`
}

type User struct {
	ID       string `json:"id"`
	Password string `json:"password"`
}

// 메뉴 : 오더 = 1:N
// 메뉴 : 리뷰 = 1:N
// 오더 : 평점 = 1:1
type Menu struct {
	MenuId         int      `json:"menuid"`         //메뉴 아이디, 추후 중복방지를 위해 둠
	ImageUrl       string   `json:"imageurl"`       //이미지 url
	Name           string   `json:"name"`           //메뉴이름
	Quantity       int      `json:"quantity"`       //총 개수
	Price          int      `json:"price"`          //메뉴가격
	Recommendation bool     `json:"recommendation"` //메뉴 추천/비추천
	Admin          string   `json:"admin"`          //관리자 이름
	Score          int      `json:"score"`          //점수
	CreateTime     string   `json:"create_time"`    //생성시간
	Orders         []Order  `json:"orders"`         //추후 통계를 위해 개발 필요 오더스를 취합해서 평점 계산필요
	Reviews        []Review `json:"reviews"`
}

/*
```
order status
- intake / true or false #주문 or 주문 취소
- cooking / true or false # 조리 중
- delivering / true of false # 배달 중
- complete / true or false #배달완료
- user string #주문자
```
*/
type Order struct {
	OrderId            int    `json:"orderid"`
	MenuName           string `json:"menuname"`           //메뉴이름
	Customer           string `json:"customer"`           //고객
	PhoneNumber        string `json:"phonenumber"`        //번호
	Address            string `json:"address"`            //주소
	Quantity           int    `json:"quantity"`           //개수
	PaymentInformation string `json:"paymentinformation"` //결제정보
	Status             string `json:"status"`             //order status
	Rating             string `json:"rating"`             // 평점
}

type Review struct {
	ReviewId      int    `json:"reviewid"`
	ReviewTitle   string `json:"review_title"`   //리뷰 제목
	ReviewContent string `json:"review_content"` //리뷰 콘텐츠
}

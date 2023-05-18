package controller

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"online-ordering-system/model"
	"reflect"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

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
	ReviewTitle   string `json:"review_title"`   //리뷰 제목
	ReviewContent string `json:"review_content"` //리뷰 콘텐츠
}

type Controller struct {
	md *model.Model
}

func NewCTL(rep *model.Model) (*Controller, error) {
	r := &Controller{md: rep}
	return r, nil
}

func (p *Controller) RespOK(c *gin.Context, resp interface{}) {
	c.JSON(http.StatusOK, resp)
}

func (p *Controller) GetOK(c *gin.Context) {
	c.JSON(200, gin.H{"msg": "ok"})
	return
}

// DB 연결 및 콜렉션 체크하는 부분 중복제거 필요함..
// func ConMongo(collection_type string) (m *mongo.Collection) {
// 	c *gin.Context
// 	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to MongoDB"})
// 		return
// 	}
// 	fmt.Println("client", client)

// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()

// 	if err := client.Connect(ctx); err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to MongoDB"})
// 		return
// 	}

// 	defer func() {
// 		if err := client.Disconnect(ctx); err != nil {
// 			log.Println("Failed to disconnect from MongoDB")
// 		}
// 	}()
// 	if collection_type == "menu" {
// 		// MongoDB 컬렉션 선택
// 		collection := client.Database("mini-oss").Collection("menu")
// 		fmt.Println("collection check menu", collection)

// 	} else if collection_type == "user_account" {
// 		collection := client.Database("mini-oss").Collection("user_account")
// 		fmt.Println("collection check user_account", collection)

// 	}

// }

func (p *Controller) RegisterHandler(c *gin.Context) {
	var user User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return

	}

	// MongoDB 연결 설정
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to MongoDB"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := client.Connect(ctx); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to MongoDB"})
		return
	}

	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			log.Println("Failed to disconnect from MongoDB")
		}
	}()

	// MongoDB 컬렉션 선택
	collection := client.Database("mini-oss").Collection("admin_account")
	fmt.Println("collection check", collection)
	// // 데이터 삽입
	if _, err := collection.InsertOne(ctx, user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert data into MongoDB"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

func (p *Controller) LoginHandler(c *gin.Context) {
	var user User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return

	}

	// MongoDB 연결 설정
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to MongoDB"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := client.Connect(ctx); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to MongoDB"})
		return
	}

	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			log.Println("Failed to disconnect from MongoDB")
		}
	}()

	// MongoDB 컬렉션 선택
	collection := client.Database("mini-oss").Collection("admin_account")

	// 사용자 로그인 여부 확인
	filter := bson.M{"id": user.ID, "password": user.Password}
	var existingUser User
	err = collection.FindOne(context.Background(), filter).Decode(&existingUser)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid credentials"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check user login"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User is logged in"})
}

func (p *Controller) UserRegisterHandler(c *gin.Context) {
	var user User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return

	}

	// MongoDB 연결 설정
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to MongoDB"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := client.Connect(ctx); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to MongoDB"})
		return
	}

	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			log.Println("Failed to disconnect from MongoDB")
		}
	}()

	// MongoDB 컬렉션 선택
	collection := client.Database("mini-oss").Collection("user_account")
	fmt.Println("collection check", collection)
	// // 데이터 삽입
	if _, err := collection.InsertOne(ctx, user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert data into MongoDB"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

func (p *Controller) UserLoginHandler(c *gin.Context) {
	var user User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return

	}

	// MongoDB 연결 설정
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to MongoDB"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := client.Connect(ctx); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to MongoDB"})
		return
	}

	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			log.Println("Failed to disconnect from MongoDB")
		}
	}()

	// MongoDB 컬렉션 선택
	collection := client.Database("mini-oss").Collection("user_account")

	// 사용자 로그인 여부 확인
	filter := bson.M{"id": user.ID, "password": user.Password}
	var existingUser User
	err = collection.FindOne(context.Background(), filter).Decode(&existingUser)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid credentials"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check user login"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User is logged in"})
}

func (p *Controller) CreateMenuHandler(c *gin.Context) {
	var menu Menu

	//프론트에서 어떻게 시간요청이 올지 감이 안와서 그냥 다 디비에 넣기
	// now := time.Now()
	// custom := now.Format("2006-01-02 15:04:05")
	// fmt.Println("custom", custom)

	// // 시간 파싱
	// nowTime := "2006-01-02 15:04:05"
	// t, err := time.Parse(layout, menu.Time)
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid time format"})
	// 	return
	// }

	// menu.Time = t.String()

	if err := c.ShouldBindJSON(&menu); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return

	}

	// MongoDB 연결 설정
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to MongoDB"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := client.Connect(ctx); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to MongoDB"})
		return
	}

	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			log.Println("Failed to disconnect from MongoDB")
		}
	}()

	// MongoDB 컬렉션 선택
	collection := client.Database("mini-oss").Collection("menu")

	// 데이터 삽입
	if _, err := collection.InsertOne(ctx, menu); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert data into MongoDB"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Menu created successfully"})
}

func (p *Controller) DeleteMenuHandler(c *gin.Context) {
	var menu Menu

	fmt.Println("menu before")
	if err := c.ShouldBindJSON(&menu); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return

	}

	// MongoDB 연결 설정
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to MongoDB"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := client.Connect(ctx); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to MongoDB"})
		return
	}

	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			log.Println("Failed to disconnect from MongoDB")
		}
	}()

	// MongoDB 컬렉션 선택
	collection := client.Database("mini-oss").Collection("menu")
	fmt.Println("collection check", collection)

	// menu_id를 기반으로 삭제할 데이터 필터 생성
	filter := bson.M{"menuid": menu.MenuId}

	// 데이터 삭제
	result, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete data from MongoDB"})
		return
	}

	if result.DeletedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "Menu not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Menu deleted successfully"})
}

func (p *Controller) MenuStatusHandler(c *gin.Context) {

	var menus []Menu

	fmt.Println("menu before")

	// MongoDB 연결 설정
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to MongoDB"})
		return
	}
	fmt.Println("client", client)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := client.Connect(ctx); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to MongoDB"})
		return
	}

	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			log.Println("Failed to disconnect from MongoDB")
		}
	}()

	// MongoDB 컬렉션 선택
	collection := client.Database("mini-oss").Collection("menu")
	fmt.Println("collection check", collection)

	//대량 검색 후 조회 이후 반환
	cursor, err := collection.Find(nil, bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve data from MongoDB"})
		return
	}
	defer cursor.Close(nil)

	if err := cursor.All(nil, &menus); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode data from MongoDB"})
		return
	}

	c.JSON(http.StatusOK, menus)
}

func (p *Controller) DetailMenuHandler(c *gin.Context) {

	menuID, err := strconv.Atoi(c.Query("menuid")) // URL에서 menuId 값을 가져옴
	if err != nil {
		fmt.Errorf(err.Error())
	}
	fmt.Println("menuID", menuID)
	fmt.Println("menuID", reflect.TypeOf(menuID))

	fmt.Println("menuID", reflect.TypeOf(menuID))
	var menu Menu

	// MongoDB 연결 설정
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to MongoDB"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := client.Connect(ctx); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to MongoDB"})
		return
	}
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			log.Println("Failed to disconnect from MongoDB")
		}
	}()

	// MongoDB 컬렉션 선택
	collection := client.Database("mini-oss").Collection("menu")
	fmt.Println("collection", collection)
	// 메뉴 조회
	err = collection.FindOne(ctx, bson.M{"menuid": menuID}).Decode(&menu)
	fmt.Println("menu", menu)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, gin.H{"error": "Menu not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve data from MongoDB"})
		}
		return
	}

	c.JSON(http.StatusOK, menu)
}

func (p *Controller) RecommendHandler(c *gin.Context) {

	var menu Menu

	// 요청된 JSON 데이터 파싱
	if err := c.ShouldBindJSON(&menu); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse JSON data"})
		return
	}

	fmt.Println("menu", menu)

	// MongoDB 연결 설정
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to MongoDB"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := client.Connect(ctx); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to MongoDB"})
		return
	}

	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			log.Println("Failed to disconnect from MongoDB")
		}
	}()

	// MongoDB 컬렉션 선택
	collection := client.Database("mini-oss").Collection("menu")
	update := bson.M{}

	if menu.Recommendation == true {
		// 업데이트할 필드 및 값 설정
		update = bson.M{
			"$set": bson.M{
				"recommendation": false,
			},
		}
	} else {
		update = bson.M{
			"$set": bson.M{
				"recommendation": true,
			},
		}
	}
	fmt.Println("update", update)
	// 업데이트 쿼리 실행
	filter := bson.M{"menuid": menu.MenuId} // 업데이트할 문서를 선택하는 필터 조건
	result, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update data in MongoDB"})
		return
	}

	if result.ModifiedCount == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No documents matched the update filter"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Menu updated successfully"})

}

func (p *Controller) CreateOrderHandler(c *gin.Context) {
	var order Order

	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return

	}

	// MongoDB 연결 설정
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to MongoDB"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := client.Connect(ctx); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to MongoDB"})
		return
	}

	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			log.Println("Failed to disconnect from MongoDB")
		}
	}()

	// MongoDB 컬렉션 선택
	collection := client.Database("mini-oss").Collection("order")

	// 데이터 삽입
	if _, err := collection.InsertOne(ctx, order); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert data into MongoDB"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order created successfully"})
}

func (p *Controller) CreateOrderReviewHandler(c *gin.Context) {
	var review Review

	if err := c.ShouldBindJSON(&review); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return

	}

	// MongoDB 연결 설정
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to MongoDB"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := client.Connect(ctx); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to MongoDB"})
		return
	}

	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			log.Println("Failed to disconnect from MongoDB")
		}
	}()

	// MongoDB 컬렉션 선택
	collection := client.Database("mini-oss").Collection("order")

	// 데이터 삽입
	if _, err := collection.InsertOne(ctx, review); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert data into MongoDB"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order created successfully"})
}

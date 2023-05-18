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

type Menu struct {
	MenuId         int    `json:"menuid"`
	ImageUrl       string `json:"imageurl"`
	Name           string `json:"name"`
	Quantity       int    `json:"quantity"`
	Price          int    `json:"price"`
	Recommendation bool   `json:"recommendation"`
	Admin          string `json:"admin"`
	Score          int    `json:"score"`
	Review         string `json:"review"`
}

type Order struct {
	MenuName           string `json:"menuname"`
	Customer           string `json:"customer"`
	PhoneNumber        string `json:"phonenumber"`
	Address            string `json:"address"`
	Quantity           int    `json:"quantity"`
	PaymentInformation string `json:"paymentinformation"`
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

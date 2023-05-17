package controller

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"online-ordering-system/model"
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
	MenuId         int    `json:"menu_id"`
	ImageUrl       string `json:"image_url"`
	Name           string `json:"name"`
	Quantity       int    `json:"quantity"`
	Price          int    `json:"price"`
	Recommendation bool   `json:"recommendation"`
	Admin          string `json:"admin"`
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

func (p *Controller) RegisterHandler(c *gin.Context) {
	var user User

	fmt.Println("user before")
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return

	}
	fmt.Println("user after", user)

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
	collection := client.Database("mini-oss").Collection("user_account")
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

	fmt.Println("user login before")
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return

	}
	fmt.Println("user login after", user)

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
	collection := client.Database("mini-oss").Collection("user_account")
	fmt.Println("collection check", collection)

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

	fmt.Println("menu before")
	if err := c.ShouldBindJSON(&menu); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return

	}
	fmt.Println("menu after", menu)

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

	filter := bson.M{"menu_id": menu.MenuId}
	update := bson.M{"$set": menu}
	upsert := false

	opts := options.UpdateOptions{
		Upsert: &upsert, // 중복 데이터가 있는 경우에는 업데이트를 수행하지 않음
	}
	fmt.Println("opts", opts)
	_, err = collection.UpdateOne(ctx, filter, update, &opts)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert data into MongoDB"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Menu created successfully"})

	// // 데이터 삽입
	// if _, err := collection.InsertOne(ctx, menu); err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert data into MongoDB"})
	// 	return
	// }

	// c.JSON(http.StatusOK, gin.H{"message": "Menu created successfully"})
}

func (p *Controller) DeleteMenuHandler(c *gin.Context) {
	var menu Menu

	fmt.Println("menu before")
	if err := c.ShouldBindJSON(&menu); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return

	}
	fmt.Println("menu after", menu)

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
	// // 데이터 삽입
	if _, err := collection.InsertOne(ctx, menu); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert data into MongoDB"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Menu registered successfully"})
}

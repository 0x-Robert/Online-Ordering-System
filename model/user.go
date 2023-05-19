package model

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)



func (m *Model) RegisterUser(user User) error {
	filter := bson.D{{Key: "id", Value: user.ID}}
	update := bson.D{{Key: "$set", Value: user}}

	opts := options.Update().SetUpsert(true)
	if _, err := m.colUserAccount.UpdateOne(context.TODO(), filter, update, opts); err != nil {
		log.Println("Failed to insert data in user_account")
		return fmt.Errorf("fail to register user: %w", err)
	}
	return nil
}

func (m *Model) LoginUser(user User) bool {
	// 사용자 로그인 여부 확인
	filter := bson.M{"id": user.ID, "password": user.Password}

	err := m.colUserAccount.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Println(err)
		}
	}
	return true
}

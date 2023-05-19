package model

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)



func (m *Model) RegisterAdmin(admin Admin) error {
	filter := bson.D{{Key: "id", Value: admin.ID}}
	update := bson.D{{Key: "$set", Value: admin}}

	opts := options.Update().SetUpsert(true)
	if _, err := m.colAdminAccount.UpdateOne(context.TODO(), filter, update, opts); err != nil {
		log.Println("Failed to insert data in user_account")
		return fmt.Errorf("fail to register user: %w", err)
	}
	return nil
}

func (m *Model) LoginAdmin(admin Admin) bool {
	// 사용자 로그인 여부 확인
	filter := bson.M{"id": admin.ID, "password": admin.Password}

	err := m.colAdminAccount.FindOne(context.Background(), filter).Decode(&admin)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Println(err)
		}
	}
	return true
}

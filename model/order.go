package model

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (m *Model) CreateOrder(order Order) error {

	// 데이터 삽입

	filter := bson.D{{Key: "id", Value: order.OrderId}}
	update := bson.D{{Key: "$set", Value: order}}

	opts := options.Update().SetUpsert(true)
	if _, err := m.colOrder.UpdateOne(context.TODO(), filter, update, opts); err != nil {
		log.Println(err)
		fmt.Errorf("fail to create menu: %w", err)
	}
	return nil
}


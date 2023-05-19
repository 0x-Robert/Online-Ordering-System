package model

import (
	"context"
	conf "online-ordering-system/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Model struct {
	client          *mongo.Client
	colMenu         *mongo.Collection
	colOrder        *mongo.Collection
	colOrderStatus  *mongo.Collection
	colUserAccount  *mongo.Collection
	colAdminAccount *mongo.Collection
}

func NewModel() (*Model, error) {

	config := conf.GetConfig("./config/.config.toml")
	r := &Model{}
	var err error
	mgUrl := config.DB["user"]["host"].(string)
	dbName := config.DB["user"]["name"].(string)
	if r.client, err = mongo.Connect(context.Background(), options.Client().ApplyURI(mgUrl)); err != nil {
		return nil, err
	} else if err := r.client.Ping(context.Background(), nil); err != nil {
		return nil, err
	} else {
		db := r.client.Database(dbName)
		r.colMenu = db.Collection("menu")
		r.colOrder = db.Collection("order")
		r.colOrderStatus = db.Collection("order_status")
		r.colAdminAccount = db.Collection("admin_account")
		r.colUserAccount = db.Collection("user_account")

	}

	return r, nil
}

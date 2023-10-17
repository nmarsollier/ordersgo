package order_proj

import (
	"context"
	"log"

	"github.com/nmarsollier/ordersgo/tools/db"
	"github.com/nmarsollier/ordersgo/tools/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Define mongo Collection
var collection *mongo.Collection

func dbCollection() (*mongo.Collection, error) {
	if collection != nil {
		return collection, nil
	}

	database, err := db.Get()
	if err != nil {
		return nil, err
	}

	col := database.Collection("order_projection")

	_, err = col.Indexes().CreateOne(
		context.Background(),
		mongo.IndexModel{
			Keys:    bson.M{"orderId": ""},
			Options: options.Index().SetUnique(true),
		},
	)
	if err != nil {
		log.Output(1, err.Error())
	}

	collection = col
	return collection, nil
}

func insert(order *Order) (*Order, error) {
	if err := order.ValidateSchema(); err != nil {
		return nil, err
	}

	var collection, err = dbCollection()
	if err != nil {
		return nil, err
	}

	_id, _ := primitive.ObjectIDFromHex(order.OrderId)
	filter := bson.M{"_id": _id}
	collection.DeleteOne(context.Background(), filter)
	if _, err := collection.InsertOne(context.Background(), order); err != nil {
		return nil, err
	}

	return order, nil
}

func FindById(orderId string) (*Order, error) {
	var collection, err = dbCollection()
	if err != nil {
		return nil, err
	}

	_id, err := primitive.ObjectIDFromHex(orderId)
	if err != nil {
		return nil, errors.ErrID
	}

	order := &Order{}
	filter := bson.M{"_id": _id}
	if err = collection.FindOne(context.Background(), filter).Decode(order); err != nil {
		return nil, err
	}

	return order, nil
}

// FindAll devuelve todos los eventos por order id
func FindByUserId(userId string) ([]*Order, error) {
	var collection, err = dbCollection()
	if err != nil {
		return nil, err
	}

	filter := bson.M{"userId": userId}
	cur, err := collection.Find(context.Background(), filter, nil)
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.Background())

	orders := []*Order{}
	for cur.Next(context.Background()) {
		order := &Order{}
		if err := cur.Decode(order); err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}

	return orders, nil
}

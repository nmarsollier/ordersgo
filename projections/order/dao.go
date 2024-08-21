package order

import (
	"context"

	"github.com/nmarsollier/ordersgo/log"
	"github.com/nmarsollier/ordersgo/tools/db"
	"github.com/nmarsollier/ordersgo/tools/errs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Define mongo Collection
var collection *mongo.Collection

func dbCollection(ctx ...interface{}) (*mongo.Collection, error) {
	if collection != nil {
		return collection, nil
	}

	database, err := db.Get(ctx...)
	if err != nil {
		log.Get(ctx...).Error(err)
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
		log.Get(ctx...).Error(err)
	}

	collection = col
	return collection, nil
}

func insert(order *Order, ctx ...interface{}) (*Order, error) {
	if err := order.ValidateSchema(); err != nil {
		log.Get(ctx...).Error(err)
		return nil, err
	}

	var collection, err = dbCollection(ctx...)
	if err != nil {
		log.Get(ctx...).Error(err)
		return nil, err
	}

	filter := bson.M{"orderId": order.OrderId}
	updateOptions := options.Update().SetUpsert(true)
	document := upsertOrder{
		Set: order,
	}

	if _, err := collection.UpdateOne(context.Background(), filter, document, updateOptions); err != nil {
		log.Get(ctx...).Error(err)
		return nil, err
	}
	return order, nil
}

type upsertOrder struct {
	Set *Order `bson:"$set"`
}

func FindByOrderId(orderId string, ctx ...interface{}) (*Order, error) {
	var collection, err = dbCollection(ctx...)
	if err != nil {
		log.Get(ctx...).Error(err)
		return nil, err
	}

	order := &Order{}
	filter := bson.M{"orderId": orderId}
	if err = collection.FindOne(context.Background(), filter).Decode(order); err != nil {
		log.Get(ctx...).Error(err)
		if err.Error() == "mongo: no documents in result" {
			return nil, errs.NotFound
		}
		return nil, err
	}

	return order, nil
}

// FindAll devuelve todos los eventos por order id
func FindByUserId(userId string, ctx ...interface{}) ([]*Order, error) {
	var collection, err = dbCollection(ctx...)
	if err != nil {
		log.Get(ctx...).Error(err)
		return nil, err
	}

	filter := bson.M{"userId": userId}
	cur, err := collection.Find(context.Background(), filter, nil)
	if err != nil {
		log.Get(ctx...).Error(err)
		return nil, err
	}
	defer cur.Close(context.Background())

	orders := []*Order{}
	for cur.Next(context.Background()) {
		order := &Order{}
		if err := cur.Decode(order); err != nil {
			log.Get(ctx...).Error(err)
			return nil, err
		}
		orders = append(orders, order)
	}

	return orders, nil
}

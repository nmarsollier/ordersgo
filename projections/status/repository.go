package status

import (
	"context"

	"github.com/nmarsollier/ordersgo/tools/db"
	"github.com/nmarsollier/ordersgo/tools/errs"
	"github.com/nmarsollier/ordersgo/tools/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Define mongo Collection
var collection *mongo.Collection

func dbCollection(deps ...interface{}) (*mongo.Collection, error) {
	if collection != nil {
		return collection, nil
	}

	database, err := db.Get(deps...)
	if err != nil {
		log.Get(deps...).Error(err)
		return nil, err
	}

	col := database.Collection("status_projection")

	_, err = col.Indexes().CreateOne(
		context.Background(),
		mongo.IndexModel{
			Keys:    bson.M{"orderId": ""},
			Options: options.Index().SetUnique(true),
		},
	)

	if err != nil {
		log.Get(deps...).Error(err)
	}

	collection = col
	return collection, nil
}

func insert(order *OrderStatus, deps ...interface{}) (*OrderStatus, error) {
	if err := order.ValidateSchema(); err != nil {
		log.Get(deps...).Error(err)
		return nil, err
	}

	var collection, err = dbCollection()
	if err != nil {
		log.Get(deps...).Error(err)
		return nil, err
	}

	filter := bson.M{"orderId": order.OrderId}
	updateOptions := options.Update().SetUpsert(true)
	document := upsertOrder{
		Set: order,
	}

	if _, err := collection.UpdateOne(context.Background(), filter, document, updateOptions); err != nil {
		log.Get(deps...).Error(err)
		return nil, err
	}
	return order, nil
}

type upsertOrder struct {
	Set *OrderStatus `bson:"$set"`
}

func FindByOrderId(orderId string, deps ...interface{}) (*OrderStatus, error) {
	var collection, err = dbCollection()
	if err != nil {
		log.Get(deps...).Error(err)
		return nil, err
	}

	order := &OrderStatus{}
	filter := bson.M{"orderId": orderId}
	if err = collection.FindOne(context.Background(), filter).Decode(order); err != nil {
		if err.Error() == "mongo: no documents in result" {
			return nil, errs.NotFound
		}
		log.Get(deps...).Error(err)
		return nil, err
	}

	return order, nil
}

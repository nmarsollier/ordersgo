package status

import (
	"context"

	"github.com/golang/glog"
	"github.com/nmarsollier/ordersgo/tools/db"
	"github.com/nmarsollier/ordersgo/tools/errs"
	"go.mongodb.org/mongo-driver/bson"
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
		glog.Error(err)
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
		glog.Error(err)
	}

	collection = col
	return collection, nil
}

func insert(order *OrderStatus) (*OrderStatus, error) {
	if err := order.ValidateSchema(); err != nil {
		glog.Error(err)
		return nil, err
	}

	var collection, err = dbCollection()
	if err != nil {
		glog.Error(err)
		return nil, err
	}

	filter := bson.M{"orderId": order.OrderId}
	updateOptions := options.Update().SetUpsert(true)
	document := upsertOrder{
		Set: order,
	}

	if _, err := collection.UpdateOne(context.Background(), filter, document, updateOptions); err != nil {
		glog.Error(err)
		return nil, err
	}
	return order, nil
}

type upsertOrder struct {
	Set *OrderStatus `bson:"$set"`
}

func FindByOrderId(orderId string) (*OrderStatus, error) {
	var collection, err = dbCollection()
	if err != nil {
		glog.Error(err)
		return nil, err
	}

	order := &OrderStatus{}
	filter := bson.M{"orderId": orderId}
	if err = collection.FindOne(context.Background(), filter).Decode(order); err != nil {
		glog.Error(err)
		if err.Error() == "mongo: no documents in result" {
			return nil, errs.NotFound
		}
		return nil, err
	}

	return order, nil
}

// FindAll devuelve todos los eventos por order id
func FindByUserId(userId string) ([]*OrderStatus, error) {
	var collection, err = dbCollection()
	if err != nil {
		glog.Error(err)
		return nil, err
	}

	filter := bson.M{"userId": userId}
	cur, err := collection.Find(context.Background(), filter, nil)
	if err != nil {
		glog.Error(err)
		return nil, err
	}
	defer cur.Close(context.Background())

	orders := []*OrderStatus{}
	for cur.Next(context.Background()) {
		order := &OrderStatus{}
		if err := cur.Decode(order); err != nil {
			glog.Error(err)
			return nil, err
		}
		orders = append(orders, order)
	}

	return orders, nil
}

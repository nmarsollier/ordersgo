package events

import (
	"context"

	"github.com/nmarsollier/ordersgo/tools/db"
	"github.com/nmarsollier/ordersgo/tools/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
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

	col := database.Collection("events")

	_, err = col.Indexes().CreateOne(
		context.Background(),
		mongo.IndexModel{
			Keys: bson.M{
				"orderId": 1, // index in ascending order
			}, Options: nil,
		},
	)
	if err != nil {
		log.Get(ctx...).Error(err)
		return nil, err
	}

	collection = col
	return collection, nil
}

func insert(event *Event, ctx ...interface{}) (*Event, error) {
	if err := event.ValidateSchema(); err != nil {
		log.Get(ctx...).Error(err)
		return nil, err
	}

	var collection, err = dbCollection(ctx...)
	if err != nil {
		log.Get(ctx...).Error(err)
		return nil, err
	}

	if _, err := collection.InsertOne(context.Background(), event); err != nil {
		log.Get(ctx...).Error(err)
		return nil, err
	}

	return event, nil
}

// findPlaceByCartId lee un usuario desde la db
func findPlaceByCartId(cartId string, ctx ...interface{}) (*Event, error) {
	var collection, err = dbCollection(ctx...)
	if err != nil {
		log.Get(ctx...).Error(err)
		return nil, err
	}

	event := &Event{}
	filter := bson.D{
		{Key: "$and",
			Value: bson.A{
				bson.M{"placeEvent.cartId": cartId},
				bson.M{"type": Place},
			},
		},
	}
	if err = collection.FindOne(context.Background(), filter).Decode(event); err != nil {
		if err.Error() != "mongo: no documents in result" {
			log.Get(ctx...).Error(err)
		}
		return nil, err
	}

	return event, nil
}

// FindAll devuelve todos los eventos por order id
func FindByOrderId(orderId string, ctx ...interface{}) ([]*Event, error) {
	var collection, err = dbCollection(ctx...)
	if err != nil {
		log.Get(ctx...).Error(err)
		return nil, err
	}

	filter := bson.M{"orderId": orderId}
	cur, err := collection.Find(context.Background(), filter, nil)
	if err != nil {
		log.Get(ctx...).Error(err)
		return nil, err
	}
	defer cur.Close(context.Background())

	events := []*Event{}
	for cur.Next(context.Background()) {
		event := &Event{}
		if err := cur.Decode(event); err != nil {
			log.Get(ctx...).Error(err)
			return nil, err
		}
		events = append(events, event)
	}

	return events, nil
}

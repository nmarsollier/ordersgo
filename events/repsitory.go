package events

import (
	"context"
	"errors"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/sso/types"
	"github.com/nmarsollier/ordersgo/tools/db"
	"github.com/nmarsollier/ordersgo/tools/errs"
	"github.com/nmarsollier/ordersgo/tools/log"
)

var tableName = "order_events"

func insert(event *Event, deps ...interface{}) (*Event, error) {
	if err := event.ValidateSchema(); err != nil {
		log.Get(deps...).Error(err)
		return nil, err
	}

	tokenToInsert, err := attributevalue.MarshalMap(event)
	if err != nil {
		log.Get(deps...).Error(err)
		return nil, err
	}

	_, err = db.Get(deps...).PutItem(
		context.TODO(),
		&dynamodb.PutItemInput{
			TableName: &tableName,
			Item:      tokenToInsert,
		},
	)
	if err != nil {
		log.Get(deps...).Error(err)
		return nil, err
	}
	return event, nil
}

// findPlaceByCartId lee un usuario desde la db
func findPlaceByCartId(cartId string, deps ...interface{}) (event *Event, err error) {
	expr, err := expression.NewBuilder().WithKeyCondition(
		expression.Key("idx_place_cart").Equal(expression.Value(cartId)),
	).Build()

	if err != nil {
		log.Get(deps...).Error(err)

		return
	}

	response, err := db.Get(deps...).Query(context.TODO(), &dynamodb.QueryInput{
		TableName:                 &tableName,
		IndexName:                 aws.String("idx_place_cart-index"),
		KeyConditionExpression:    expr.KeyCondition(),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
	})

	if temp := new(types.ResourceNotFoundException); err != nil && !errors.As(err, &temp) {
		log.Get(deps...).Error(err)

		return nil, errs.NotFound
	}

	if err != nil || len(response.Items) == 0 {
		return nil, errs.NotFound
	}

	err = attributevalue.UnmarshalMap(response.Items[0], &event)

	return
}

// FindAll devuelve todos los eventos por order id
func FindByOrderId(orderId string, deps ...interface{}) (events []*Event, err error) {
	expr, err := expression.NewBuilder().WithKeyCondition(
		expression.Key("orderId").Equal(expression.Value(orderId)),
	).Build()
	if err != nil {
		log.Get(deps...).Error(err)

		return
	}

	response, err := db.Get(deps...).Query(context.TODO(), &dynamodb.QueryInput{
		TableName:                 &tableName,
		IndexName:                 aws.String("orderId-index"),
		KeyConditionExpression:    expr.KeyCondition(),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
	})

	if err != nil || len(response.Items) == 0 {
		log.Get(deps...).Error(err)

		return
	}

	err = attributevalue.UnmarshalListOfMaps(response.Items, &events)
	if err != nil {
		log.Get(deps...).Error(err)
	}
	return
}

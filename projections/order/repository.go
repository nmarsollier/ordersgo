package order

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/nmarsollier/ordersgo/tools/db"
	"github.com/nmarsollier/ordersgo/tools/errs"
	"github.com/nmarsollier/ordersgo/tools/log"
)

var tableName = "orders_projection_order"

func insert(order *Order, deps ...interface{}) (orderResult *Order, err error) {
	if err = order.ValidateSchema(); err != nil {
		log.Get(deps...).Error(err)
		return
	}

	orderToInsert, err := attributevalue.MarshalMap(order)
	if err != nil {
		log.Get(deps...).Error(err)
		return
	}

	_, err = db.Get(deps...).PutItem(
		context.TODO(),
		&dynamodb.PutItemInput{
			TableName: &tableName,
			Item:      orderToInsert,
		},
	)

	if err != nil {
		log.Get(deps...).Error(err)
	}

	return
}

func FindByOrderId(orderId string, deps ...interface{}) (order *Order, err error) {
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

	if err != nil {
		log.Get(deps...).Error(err)

		return
	}

	if len(response.Items) == 0 {
		log.Get(deps...).Error(err)

		return nil, errs.NotFound
	}

	err = attributevalue.UnmarshalMap(response.Items[0], &order)
	if err != nil {
		log.Get(deps...).Error(err)
	}

	return
}

// FindAll devuelve todos los eventos por order id
func FindByUserId(userId string, deps ...interface{}) (orders []*Order, err error) {
	expr, err := expression.NewBuilder().WithKeyCondition(
		expression.Key("userId").Equal(expression.Value(userId)),
	).Build()
	if err != nil {
		log.Get(deps...).Error(err)

		return
	}

	response, err := db.Get(deps...).Query(context.TODO(), &dynamodb.QueryInput{
		TableName:                 &tableName,
		IndexName:                 aws.String("userId-index"),
		KeyConditionExpression:    expr.KeyCondition(),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
	})

	if err != nil || len(response.Items) == 0 {
		log.Get(deps...).Error(err)

		return
	}

	err = attributevalue.UnmarshalListOfMaps(response.Items, &orders)
	if err != nil {
		log.Get(deps...).Error(err)
	}

	return
}

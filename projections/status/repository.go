package status

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

var tableName = "orders_projection_status"

func insert(order *OrderStatus, deps ...interface{}) (status *OrderStatus, err error) {
	if err = order.ValidateSchema(); err != nil {
		log.Get(deps...).Error(err)
		return
	}

	stateToInsert, err := attributevalue.MarshalMap(order)
	if err != nil {
		log.Get(deps...).Error(err)

		return
	}

	_, err = db.Get(deps...).PutItem(
		context.TODO(),
		&dynamodb.PutItemInput{
			TableName: &tableName,
			Item:      stateToInsert,
		},
	)
	if err != nil {
		log.Get(deps...).Error(err)
	}

	return
}

func FindByOrderId(orderId string, deps ...interface{}) (status *OrderStatus, err error) {
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
		ScanIndexForward:          aws.Bool(true),
	})

	if err != nil {
		log.Get(deps...).Error(err)
		return
	}

	if len(response.Items) == 0 {
		return nil, errs.NotFound
	}

	err = attributevalue.UnmarshalMap(response.Items[0], &status)
	if err != nil {
		log.Get(deps...).Error(err)
	}

	return
}

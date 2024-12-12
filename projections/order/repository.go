package order

import (
	"context"

	"github.com/nmarsollier/ordersgo/tools/db"
	"github.com/nmarsollier/ordersgo/tools/errs"
	"github.com/nmarsollier/ordersgo/tools/log"
	"github.com/nmarsollier/ordersgo/tools/strs"
)

func insert(order *Order, deps ...interface{}) (orderResult *Order, err error) {
	if err = order.ValidateSchema(); err != nil {
		log.Get(deps...).Error(err)
		return
	}

	conn, err := db.GetPostgresClient(deps...)
	if err != nil {
		log.Get(deps...).Error(err)
		return nil, err
	}

	query := `
        INSERT INTO ordersgo.Orders (ID, OrderId, Status, UserId, CartId, Articles, Payments, Created, Updated)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
				ON CONFLICT (id) DO UPDATE SET
						Status = COALESCE(EXCLUDED.Status, Orders.Status),
            Articles = COALESCE(EXCLUDED.Articles, Orders.Articles),
            Payments = COALESCE(EXCLUDED.Payments, Orders.Payments),
            Updated = EXCLUDED.Updated
    `

	articles := strs.ToJson(order.Articles)
	payments := strs.ToJson(order.Payments)

	_, err = conn.Exec(context.Background(), query, order.ID, order.OrderId, order.Status, order.UserId, order.CartId, articles, payments, order.Created, order.Updated)
	if err != nil {
		log.Get(deps...).Error(err)
		return nil, err
	}

	return order, nil
}

func FindByOrderId(orderId string, deps ...interface{}) (*Order, error) {
	query := `
        SELECT ID, OrderId, Status, UserId, CartId, Articles, Payments, Created, Updated
        FROM ordersgo.Orders
        WHERE OrderId = $1
        ORDER BY Created ASC
        LIMIT 1
    `

	conn, err := db.GetPostgresClient(deps...)
	if err != nil {
		log.Get(deps...).Error(err)
		return nil, err
	}

	row := conn.QueryRow(context.Background(), query, orderId)

	var order Order
	var articles, payments []byte

	err = row.Scan(&order.ID, &order.OrderId, &order.Status, &order.UserId, &order.CartId, &articles, &payments, &order.Created, &order.Updated)
	if err != nil {
		if err.Error() == "no rows in result set" {
			return nil, errs.NotFound
		}
		log.Get(deps...).Error(err)
		return nil, err
	}

	strs.FromJson(string(articles), &order.Articles)
	strs.FromJson(string(payments), &order.Payments)

	return &order, nil
}

// FindAll devuelve todos los eventos por order id
func FindByUserId(userId string, deps ...interface{}) ([]*Order, error) {
	query := `
        SELECT ID, OrderId, Status, UserId, CartId, Articles, Payments, Created, Updated
        FROM ordersgo.Orders
        WHERE UserId = $1
        ORDER BY Created ASC
    `

	conn, err := db.GetPostgresClient(deps...)
	if err != nil {
		log.Get(deps...).Error(err)
		return nil, err
	}

	rows, err := conn.Query(context.Background(), query, userId)
	if err != nil {
		log.Get(deps...).Error(err)
		return nil, err
	}
	defer rows.Close()

	var orders []*Order

	for rows.Next() {
		var order Order
		var articles, payments []byte

		err := rows.Scan(&order.ID, &order.OrderId, &order.Status, &order.UserId, &order.CartId, &articles, &payments, &order.Created, &order.Updated)
		if err != nil {
			log.Get(deps...).Error(err)
			return nil, err
		}

		strs.FromJson(string(articles), &order.Articles)
		strs.FromJson(string(payments), &order.Payments)

		orders = append(orders, &order)
	}

	if err := rows.Err(); err != nil {
		log.Get(deps...).Error(err)
		return nil, err
	}

	return orders, nil
}

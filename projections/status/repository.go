package status

import (
	"context"

	"github.com/nmarsollier/ordersgo/tools/db"
	"github.com/nmarsollier/ordersgo/tools/errs"
	"github.com/nmarsollier/ordersgo/tools/log"
)

func insert(order *OrderStatus, deps ...interface{}) (status *OrderStatus, err error) {
	if err = order.ValidateSchema(); err != nil {
		log.Get(deps...).Error(err)
		return
	}

	query := `
        INSERT INTO ordersgo.OrderStatus (id, OrderId, UserId, Placed, PartialValidated, Validated, PartialPayment, PaymentCompleted, Created, Updated)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
        ON CONFLICT (id) DO UPDATE SET
            Placed = EXCLUDED.Placed,
            PartialValidated = EXCLUDED.PartialValidated,
            Validated = EXCLUDED.Validated,
            PartialPayment = EXCLUDED.PartialPayment,
            PaymentCompleted = EXCLUDED.PaymentCompleted,
            Updated = EXCLUDED.Updated
    `

	conn, err := db.GetPostgresClient(deps...)
	if err != nil {
		log.Get(deps...).Error(err)
		return nil, err
	}

	_, err = conn.Exec(context.Background(), query, order.ID, order.OrderId, order.UserId, order.Placed, order.PartialValidated, order.Validated, order.PartialPayment, order.PaymentCompleted, order.Created, order.Updated)
	if err != nil {
		log.Get(deps...).Error(err)
		return nil, err
	}

	return order, nil
}

func FindByOrderId(orderId string, deps ...interface{}) (*OrderStatus, error) {
	query := `
        SELECT id, OrderId, UserId, Placed, PartialValidated, Validated, PartialPayment, PaymentCompleted, Created, Updated
        FROM ordersgo.OrderStatus
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

	var status OrderStatus

	err = row.Scan(&status.ID, &status.OrderId, &status.UserId, &status.Placed, &status.PartialValidated, &status.Validated, &status.PartialPayment, &status.PaymentCompleted, &status.Created, &status.Updated)
	if err != nil {
		if err.Error() == "no rows in result set" {
			return nil, errs.NotFound
		}
		log.Get(deps...).Error(err)
		return nil, err
	}

	return &status, nil
}

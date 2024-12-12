package events

import (
	"context"
	"encoding/json"

	"github.com/nmarsollier/ordersgo/tools/db"
	"github.com/nmarsollier/ordersgo/tools/errs"
	"github.com/nmarsollier/ordersgo/tools/log"
	"github.com/nmarsollier/ordersgo/tools/strs"
)

func insert(event *Event, deps ...interface{}) (*Event, error) {
	if err := event.ValidateSchema(); err != nil {
		log.Get(deps...).Error(err)
		return nil, err
	}

	query := `
        INSERT INTO ordersgo.Events (id, OrderId, Type, PlaceEvent, Validation, Payment, Created, Updated)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
    `

	placeEvent := strs.ToJson(event.PlaceEvent)
	validation := strs.ToJson(event.Validation)
	payment := strs.ToJson(event.Payment)

	conn, err := db.GetPostgresClient(deps...)
	if err != nil {
		log.Get(deps...).Error(err)
		return nil, err
	}

	_, err = conn.Exec(context.Background(), query, event.ID, event.OrderId, event.Type, placeEvent, validation, payment, event.Created, event.Updated)
	if err != nil {
		log.Get(deps...).Error(err)
		return nil, err
	}

	return event, nil
}

// findPlaceByCartId lee un usuario desde la db
func findPlaceByCartId(cartId string, deps ...interface{}) (*Event, error) {
	query := `
        SELECT id, OrderId, Type, PlaceEvent, Validation, Payment, Created, Updated
        FROM ordersgo.Events
        WHERE PlaceEvent->>'cartId' = $1
        ORDER BY Created ASC
        LIMIT 1
    `

	conn, err := db.GetPostgresClient(deps...)
	if err != nil {
		log.Get(deps...).Error(err)
		return nil, err
	}

	row := conn.QueryRow(context.Background(), query, cartId)

	var event Event
	var placeEvent, validation, payment []byte

	err = row.Scan(&event.ID, &event.OrderId, &event.Type, &placeEvent, &validation, &payment, &event.Created, &event.Updated)
	if err != nil {
		if err.Error() == "no rows in result set" {
			return nil, errs.NotFound
		}
		log.Get(deps...).Error(err)
		return nil, err
	}

	if err := json.Unmarshal(placeEvent, &event.PlaceEvent); err != nil {
		log.Get(deps...).Error(err)
		return nil, err
	}

	if err := json.Unmarshal(validation, &event.Validation); err != nil {
		log.Get(deps...).Error(err)
		return nil, err
	}

	if err := json.Unmarshal(payment, &event.Payment); err != nil {
		log.Get(deps...).Error(err)
		return nil, err
	}

	return &event, nil
}

// FindAll devuelve todos los eventos por order id
func FindByOrderId(orderId string, deps ...interface{}) ([]*Event, error) {
	query := `
        SELECT id, OrderId, Type, PlaceEvent, Validation, Payment, Created, Updated
        FROM ordersgo.Events
        WHERE OrderId = $1
        ORDER BY Created ASC
    `

	conn, err := db.GetPostgresClient(deps...)
	if err != nil {
		log.Get(deps...).Error(err)
		return nil, err
	}

	rows, err := conn.Query(context.Background(), query, orderId)
	if err != nil {
		log.Get(deps...).Error(err)
		return nil, err
	}
	defer rows.Close()

	var events []*Event

	for rows.Next() {
		var event Event
		var placeEvent, validation, payment []byte

		err := rows.Scan(&event.ID, &event.OrderId, &event.Type, &placeEvent, &validation, &payment, &event.Created, &event.Updated)
		if err != nil {
			log.Get(deps...).Error(err)
			return nil, err
		}

		strs.FromJson(string(placeEvent), &event.PlaceEvent)
		strs.FromJson(string(validation), &event.Validation)
		strs.FromJson(string(payment), &event.Payment)

		events = append(events, &event)
	}

	if err := rows.Err(); err != nil {
		log.Get(deps...).Error(err)
		return nil, err
	}

	return events, nil
}

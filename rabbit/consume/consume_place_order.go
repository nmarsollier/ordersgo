package consume

import (
	"encoding/json"

	"github.com/nmarsollier/ordersgo/events"
	"github.com/nmarsollier/ordersgo/log"
	"github.com/nmarsollier/ordersgo/services"
	"github.com/nmarsollier/ordersgo/tools/env"
	uuid "github.com/satori/go.uuid"
	"github.com/streadway/amqp"
)

//	@Summary		Mensage Rabbit place_order/order_place_order
//	@Description	Cuando se consume place_order se genera la orden y se inicia el proceso.
//	@Tags			Rabbit
//	@Accept			json
//	@Produce		json
//	@Param			place_order	body	consumePlaceDataMessage	true	"Consume place_order/order_place_order"
//	@Router			/rabbit/place_order [get]
//
// Validar Artículos
func consumePlaceOrder() error {
	logger := log.Get().
		WithField("Controller", "Rabbit").
		WithField("Queue", "place_order").
		WithField("Method", "Consume")

	conn, err := amqp.Dial(env.Get().RabbitURL)
	if err != nil {
		logger.Error(err)
		return err
	}
	defer conn.Close()

	chn, err := conn.Channel()
	if err != nil {
		logger.Error(err)
		return err
	}
	defer chn.Close()

	err = chn.ExchangeDeclare(
		"place_order", // name
		"direct",      // type
		false,         // durable
		false,         // auto-deleted
		false,         // internal
		false,         // no-wait
		nil,           // arguments
	)
	if err != nil {
		logger.Error(err)
		return err
	}

	queue, err := chn.QueueDeclare(
		"order_place_order", // name
		false,               // durable
		false,               // delete when unused
		false,               // exclusive
		false,               // no-wait
		nil,                 // arguments
	)
	if err != nil {
		logger.Error(err)
		return err
	}

	err = chn.QueueBind(
		queue.Name,    // queue name
		"place_order", // routing key
		"place_order", // exchange
		false,
		nil)
	if err != nil {
		logger.Error(err)
		return err
	}

	mgs, err := chn.Consume(
		queue.Name, // queue
		"",         // consumer
		false,      // auto-ack
		false,      // exclusive
		false,      // no-local
		false,      // no-wait
		nil,        // args
	)
	if err != nil {
		logger.Error(err)
		return err
	}

	logger.Info("RabbitMQ consumeOrdersChannel conectado")

	go func() {
		for d := range mgs {
			newMessage := &consumePlaceDataMessage{}
			body := d.Body
			logger.Info("Incomming place_order : ", string(body))

			err = json.Unmarshal(body, newMessage)
			if err == nil {
				l := logger.WithField("CorrelationId", getOrderPlacedCorrelationId(newMessage))

				processPlaceOrder(newMessage, l)

				if err := d.Ack(false); err != nil {
					l.Info("Failed ACK :", string(body), err)
				} else {
					l.Info("Consumed place_order :", string(body))
				}
			} else {
				logger.Error(err)
			}
		}
	}()

	logger.Info("Closed connection: ", <-conn.NotifyClose(make(chan *amqp.Error)))

	return nil
}

func processPlaceOrder(newMessage *consumePlaceDataMessage, ctx ...interface{}) {
	data := newMessage.Message

	_, err := services.PocessPlaceOrder(data, ctx...)
	if err != nil {
		log.Get(ctx...).Error(err)
		return
	}
}

type consumePlaceDataMessage struct {
	CorrelationId string `json:"correlation_id" example:"123123" `
	Message       *events.PlacedOrderData
}

func getOrderPlacedCorrelationId(c *consumePlaceDataMessage) string {
	value := c.CorrelationId

	if len(value) == 0 {
		value = uuid.NewV4().String()
	}

	return value
}

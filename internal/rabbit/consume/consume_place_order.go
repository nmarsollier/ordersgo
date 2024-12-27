package consume

import (
	"encoding/json"

	"github.com/nmarsollier/ordersgo/internal/engine/log"
	"github.com/nmarsollier/ordersgo/internal/events"
	"github.com/nmarsollier/ordersgo/internal/services"
	uuid "github.com/satori/go.uuid"
	"github.com/streadway/amqp"
)

type OrderPlacedConsumer interface {
	ConsumeOrderPlaced() error
}

func NewOrderPlacedConsumer(fluentUrl string, rabbitUrl string, service services.Service) OrderPlacedConsumer {
	logger := log.Get(fluentUrl).
		WithField(log.LOG_FIELD_CONTROLLER, "Rabbit").
		WithField(log.LOG_FIELD_RABBIT_QUEUE, "cart_order_placed").
		WithField(log.LOG_FIELD_RABBIT_EXCHANGE, "order_placed").
		WithField(log.LOG_FIELD_RABBIT_ACTION, "Consume")

	return &orderPlacedConsumer{
		logger:    logger,
		rabbitUrl: rabbitUrl,
		service:   service,
	}
}

type orderPlacedConsumer struct {
	logger    log.LogRusEntry
	rabbitUrl string
	service   services.Service
}

//	@Summary		Mensage Rabbit place_order/order_place_order
//	@Description	Cuando se consume place_order se genera la orden y se inicia el proceso.
//	@Tags			Rabbit
//	@Accept			json
//	@Produce		json
//	@Param			place_order	body	consumePlaceDataMessage	true	"Consume place_order/order_place_order"
//	@Router			/rabbit/place_order [get]
//
// Validar Artículos
func (r *orderPlacedConsumer) ConsumeOrderPlaced() error {
	conn, err := amqp.Dial(r.rabbitUrl)
	if err != nil {
		r.logger.Error(err)
		return err
	}
	defer conn.Close()

	chn, err := conn.Channel()
	if err != nil {
		r.logger.Error(err)
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
		r.logger.Error(err)
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
		r.logger.Error(err)
		return err
	}

	err = chn.QueueBind(
		queue.Name,    // queue name
		"place_order", // routing key
		"place_order", // exchange
		false,
		nil)
	if err != nil {
		r.logger.Error(err)
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
		r.logger.Error(err)
		return err
	}

	r.logger.Info("RabbitMQ consumeOrdersChannel conectado")

	go func() {
		for d := range mgs {
			newMessage := &consumePlaceDataMessage{}
			body := d.Body

			err = json.Unmarshal(body, newMessage)
			if err == nil {
				r.logger.WithField(log.LOG_FIELD_CORRELATION_ID, r.getOrderPlacedCorrelationId(newMessage))
				r.logger.Info("Incomming place_order : ", string(body))

				r.processPlaceOrder(newMessage)

				if err := d.Ack(false); err != nil {
					r.logger.Info("Failed ACK :", string(body), err)
				} else {
					r.logger.Info("Consumed place_order :", string(body))
				}
			} else {
				r.logger.Error(err)
			}
		}
	}()

	r.logger.Info("Closed connection: ", <-conn.NotifyClose(make(chan *amqp.Error)))

	return nil
}

func (r *orderPlacedConsumer) processPlaceOrder(newMessage *consumePlaceDataMessage) {
	data := newMessage.Message

	_, err := r.service.PocessPlaceOrder(data)
	if err != nil {
		r.logger.Error(err)
		return
	}
}

type consumePlaceDataMessage struct {
	CorrelationId string `json:"correlation_id" example:"123123" `
	Message       *events.PlacedOrderData
}

func (r *orderPlacedConsumer) getOrderPlacedCorrelationId(c *consumePlaceDataMessage) string {
	value := c.CorrelationId

	if len(value) == 0 {
		value = uuid.NewV4().String()
	}

	return value
}

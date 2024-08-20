package consume

import (
	"encoding/json"

	"github.com/golang/glog"
	"github.com/nmarsollier/ordersgo/events"
	"github.com/nmarsollier/ordersgo/services"
	"github.com/nmarsollier/ordersgo/tools/env"
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
	conn, err := amqp.Dial(env.Get().RabbitURL)
	if err != nil {
		glog.Error(err)
		return err
	}
	defer conn.Close()

	chn, err := conn.Channel()
	if err != nil {
		glog.Error(err)
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
		glog.Error(err)
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
		glog.Error(err)
		return err
	}

	err = chn.QueueBind(
		queue.Name,    // queue name
		"place_order", // routing key
		"place_order", // exchange
		false,
		nil)
	if err != nil {
		glog.Error(err)
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
		glog.Error(err)
		return err
	}

	glog.Info("RabbitMQ consumeOrdersChannel conectado")

	go func() {
		for d := range mgs {
			newMessage := &consumePlaceDataMessage{}
			body := d.Body
			glog.Info("Incomming place_order : ", string(body))

			err = json.Unmarshal(body, newMessage)
			if err == nil {
				processPlaceOrder(newMessage)

				if err := d.Ack(false); err != nil {
					glog.Info("Failed ACK :", string(body), err)
				} else {
					glog.Info("Consumed place_order :", string(body))
				}
			} else {
				glog.Error(err)
			}
		}
	}()

	glog.Info("Closed connection: ", <-conn.NotifyClose(make(chan *amqp.Error)))

	return nil
}

func processPlaceOrder(newMessage *consumePlaceDataMessage) {
	data := newMessage.Message

	_, err := services.PocessPlaceOrder(data)
	if err != nil {
		glog.Error(err)
		return
	}
}

type consumePlaceDataMessage struct {
	Message *events.PlacedOrderData
}

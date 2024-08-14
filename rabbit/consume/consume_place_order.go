package consume

import (
	"encoding/json"

	"github.com/golang/glog"
	"github.com/nmarsollier/ordersgo/events"
	"github.com/nmarsollier/ordersgo/services"
	"github.com/nmarsollier/ordersgo/tools/env"
	"github.com/streadway/amqp"
)

//	@Summary		Mensage Rabbit order/article-data
//	@Description	Cuando se consume place-order se genera la orden y se inicia el proceso.
//	@Tags			Rabbit
//	@Accept			json
//	@Produce		json
//	@Param			place-order	body	consumePlaceDataMessage	true	"Message para Type = place-order"
//	@Router			/rabbit/article-data [get]
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
		"order",  // name
		"direct", // type
		false,    // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)
	if err != nil {
		glog.Error(err)
		return err
	}

	queue, err := chn.QueueDeclare(
		"order", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	if err != nil {
		glog.Error(err)
		return err
	}

	err = chn.QueueBind(
		queue.Name, // queue name
		"order",    // routing key
		"order",    // exchange
		false,
		nil)
	if err != nil {
		glog.Error(err)
		return err
	}

	mgs, err := chn.Consume(
		queue.Name,     // queue
		"order_placed", // consumer
		true,           // auto-ack
		false,          // exclusive
		false,          // no-local
		false,          // no-wait
		nil,            // args
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
			glog.Info("Rabbit Consume : ", string(body))

			err = json.Unmarshal(body, newMessage)
			if err == nil {
				switch newMessage.Type {
				case "place-order":
					processPlaceOrder(newMessage)
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

	event, err := services.PocessPlaceOrder(data)
	if err != nil {
		glog.Error(err)
		return
	}

	glog.Info("Order placed completed : ", event)
}

type consumePlaceDataMessage struct {
	Type     string `json:"type" example:"place-order"`
	Version  int    `json:"version"`
	Queue    string `json:"queue"`
	Exchange string `json:"exchange"`
	Message  *events.PlacedOrderData
}

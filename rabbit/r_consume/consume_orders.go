package r_consume

import (
	"encoding/json"

	"github.com/golang/glog"
	"github.com/nmarsollier/ordersgo/events"
	"github.com/nmarsollier/ordersgo/services"
	"github.com/nmarsollier/ordersgo/tools/env"
	"github.com/streadway/amqp"
)

func consumeOrders() error {
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
		queue.Name, // queue
		"",         // consumer
		true,       // auto-ack
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
			newMessage := &ConsumeMessage{}
			body := d.Body
			glog.Info("Rabbit Consume : ", string(body))

			err = json.Unmarshal(body, newMessage)
			if err == nil {
				switch newMessage.Type {
				case "article-data":
					articleMessage := &ConsumeArticleDataMessage{}
					if err := json.Unmarshal(body, articleMessage); err != nil {
						glog.Error("Error decoding Article Data", err)
						return
					}

					processArticleData(articleMessage)
				case "place-order":
					placeMessage := &ConsumePlaceDataMessage{}
					if err := json.Unmarshal(body, placeMessage); err != nil {
						glog.Error("Error decoding Place Data", err)
						return
					}
					err = json.Unmarshal(body, newMessage)
					processPlaceOrder(placeMessage)
				}
			} else {
				glog.Error(err)
			}
		}
	}()

	glog.Info("Closed connection: ", <-conn.NotifyClose(make(chan *amqp.Error)))

	return nil
}

// @Summary		Mensage Rabbit order/article-data
// @Description	Antes de iniciar las operaciones se validan los artículos contra el catalogo.
// @Tags			Rabbit
// @Accept			json
// @Produce		json
// @Param			article-data	body	ConsumeArticleDataMessage	true	"Message para Type = article-data"
// @Router			/rabbit/article-data [get]
//
// Validar Artículos
func processArticleData(newMessage *ConsumeArticleDataMessage) {
	data := newMessage.Message

	event, err := services.ProcessArticleData(data)
	if err != nil {
		glog.Error(err)
		return
	}

	glog.Info("Article exist completed : ", event.ID.Hex())
}

type ConsumeArticleDataMessage struct {
	Type     string `json:"type"`
	Version  int    `json:"version"`
	Queue    string `json:"queue"`
	Exchange string `json:"exchange"`
	Message  *events.ValidationEvent
}

// @Summary		Mensage Rabbit order/article-data
// @Description	Antes de iniciar las operaciones se validan los artículos contra el catalogo.
// @Tags			Rabbit
// @Accept			json
// @Produce		json
// @Param			place-order	body	ConsumePlaceDataMessage	true	"Message para Type = place-order"
// @Router			/rabbit/article-data [get]
//
// Validar Artículos
func processPlaceOrder(newMessage *ConsumePlaceDataMessage) {
	data := newMessage.Message

	event, err := services.PocessPlaceOrder(data)
	if err != nil {
		glog.Error(err)
		return
	}

	glog.Info("Order placed completed : ", event)
}

type ConsumePlaceDataMessage struct {
	Type     string `json:"type"`
	Version  int    `json:"version"`
	Queue    string `json:"queue"`
	Exchange string `json:"exchange"`
	Message  *events.PlacedOrderData
}

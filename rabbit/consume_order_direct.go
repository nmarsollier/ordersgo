package rabbit

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/nmarsollier/ordersgo/events"
	"github.com/nmarsollier/ordersgo/services"
	"github.com/nmarsollier/ordersgo/tools/env"
	"github.com/streadway/amqp"
)

func consumeOrdersChannel() error {
	conn, err := amqp.Dial(env.Get().RabbitURL)
	if err != nil {
		return err
	}
	defer conn.Close()

	chn, err := conn.Channel()
	if err != nil {
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
		return err
	}

	err = chn.QueueBind(
		queue.Name, // queue name
		"order",    // routing key
		"order",    // exchange
		false,
		nil)
	if err != nil {
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
		return err
	}

	fmt.Println("RabbitMQ consumeOrdersChannel conectado")

	go func() {
		for d := range mgs {
			newMessage := &ConsumeMessage{}
			body := d.Body

			fmt.Println(string(body))
			err = json.Unmarshal(body, newMessage)
			if err == nil {
				switch newMessage.Type {
				case "article-data":
					articleMessage := &ConsumeArticleDataMessage{}
					if err := json.Unmarshal(body, articleMessage); err != nil {
						log.Print("Error decoding Article Data")
						return
					}

					processArticleData(articleMessage)
				case "place-order":
					placeMessage := &ConsumePlaceDataMessage{}
					if err := json.Unmarshal(body, placeMessage); err != nil {
						log.Print("Error decoding Place Data")
						return
					}
					err = json.Unmarshal(body, newMessage)
					processPlaceOrder(placeMessage)
				}
			}
		}
	}()

	fmt.Println("Closed connection: ", <-conn.NotifyClose(make(chan *amqp.Error)))

	return nil
}

type ConsumeArticleDataMessage struct {
	Type     string `json:"type"`
	Version  int    `json:"version"`
	Queue    string `json:"queue"`
	Exchange string `json:"exchange"`
	Message  *events.ValidationEvent
}

// Validar Artículos
//
//	@Summary		Mensage Rabbit order/article-data
//	@Description	Antes de iniciar las operaciones se validan los artículos contra el catalogo.
//	@Tags			Rabbit
//	@Accept			json
//	@Produce		json
//	@Param			article-data	body	ConsumeArticleDataMessage	true	"Message para Type = article-data"
//
//	@Router			/rabbit/article-data [put]
func processArticleData(newMessage *ConsumeArticleDataMessage) {
	data := newMessage.Message

	event, err := services.ProcessArticleData(data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	log.Print("Article exist completed : " + event.ID.Hex())
}

type ConsumePlaceDataMessage struct {
	Type     string `json:"type"`
	Version  int    `json:"version"`
	Queue    string `json:"queue"`
	Exchange string `json:"exchange"`
	Message  *events.PlacedOrderData
}

// Validar Artículos
//
//	@Summary		Mensage Rabbit order/article-data
//	@Description	Antes de iniciar las operaciones se validan los artículos contra el catalogo.
//	@Tags			Rabbit
//	@Accept			json
//	@Produce		json
//	@Param			place-order	body	ConsumePlaceDataMessage	true	"Message para Type = place-order"
//
//	@Router			/rabbit/article-data [put]
func processPlaceOrder(newMessage *ConsumePlaceDataMessage) {
	data := newMessage.Message

	event, err := services.PocessPlaceOrder(data)
	if err != nil {
		log.Print("Invalid Article Data " + err.Error())
		return
	}

	EmitOrderPlaced(event)

	for _, article := range event.PlaceEvent.Articles {
		go SendArticleValidation(ArticleValidationData{
			ReferenceId: event.OrderId,
			ArticleId:   article.ArticleId,
		})
	}

	log.Print("Order placed completed : " + event.OrderId)
}

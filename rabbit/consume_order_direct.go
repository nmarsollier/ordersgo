package rabbit

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/nmarsollier/ordersgo/events"
	"github.com/nmarsollier/ordersgo/order_proj"
	"github.com/nmarsollier/ordersgo/tools/env"
	"github.com/streadway/amqp"
)

/**
 *
 * @api {direct} order/article-data Validar Artículos
 *
 * @apiGroup RabbitMQ GET
 *
 * @apiDescription Antes de iniciar las operaciones se validan los artículos contra el catalogo.
 *
 * @apiExample {json} Mensaje
 *     {
 *     "type": "article-data",
 *     "message" : {
 *         "cartId": "{cartId}",
 *         "articleId": "{articleId}",
 *         "valid": True|False
 *        }
 *     }
 */

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
		true,    // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	if err != nil {
		return err
	}

	err = chn.QueueBind(
		queue.Name, // queue name
		"",         // routing key
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
			newMessage := &consume_message{}
			err = json.Unmarshal(d.Body, newMessage)
			if err == nil {
				switch newMessage.Type {
				case "article-data":
					processArticleData(newMessage)
				case "place-order":
					processPlaceOrder(newMessage)
				}
			}
		}
	}()

	fmt.Println("Closed connection: ", <-conn.NotifyClose(make(chan *amqp.Error)))

	return nil
}

func processArticleData(newMessage *consume_message) {
	data := &events.ValidationEvent{}

	if err := json.Unmarshal([]byte(newMessage.Message), data); err != nil {
		log.Print("Error decoding Article Data")
		return
	}

	event, err := events.SaveArticleExist(data)
	if err != nil {
		log.Print("Invalid Article Data " + err.Error())
		return
	}

	log.Print("Article exist completed : " + event.ID.Hex())

	go order_proj.UpdateOrderProjection(event.OrderId)
}

func processPlaceOrder(newMessage *consume_message) {
	data := &events.PlacedOrderData{}

	if err := json.Unmarshal([]byte(newMessage.Message), data); err != nil {
		log.Print("Error decoding Article Data")
		return
	}

	event, err := events.SavePlaceOrder(data)
	if err != nil {
		log.Print("Invalid Article Data " + err.Error())
		return
	}

	EmitOrderPlaced(event)

	log.Print("Order placed completed : " + event.OrderId)

	go order_proj.UpdateOrderProjection(event.OrderId)
}

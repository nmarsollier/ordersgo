package rabbit

import (
	"encoding/json"
	"log"

	"github.com/nmarsollier/ordersgo/events"
	"github.com/streadway/amqp"
)

/**
 *
 * @api {topic} order/order-placed Orden Creada
 *
 * @apiGroup RabbitMQ POST
 *
 * @apiDescription Envía de mensajes order-placed desde Order con el topic "order_placed".
 *
 * @apiSuccessExample {json} Mensaje
 *     {
 *     "type": "order-placed",
 *     "message" : {
 *         "cartId": "{cartId}",
 *         "orderId": "{orderId}"
 *         "articles": [{
 *              "articleId": "{article id}"
 *              "quantity" : {quantity}
 *          }, ...]
 *        }
 *     }
 *
 */
func EmitOrderPlaced(data *events.Event) error {
	type message struct {
		Type     string          `json:"type"`
		Exchange string          `json:"exchange"`
		Queue    string          `json:"queue"`
		Message  orderPlacedData `json:"message"`
	}

	send := message{
		Type:     "order-placed",
		Exchange: "order",
		Queue:    "order",
		Message:  *toPlaceData(data),
	}

	chn, err := getChannel()
	if err != nil {
		chn = nil
		return err
	}

	err = chn.ExchangeDeclare(
		"order_placed", // name
		"direct",       // type
		false,          // durable
		false,          // auto-deleted
		false,          // internal
		false,          // no-wait
		nil,            // arguments
	)
	if err != nil {
		chn = nil
		return err
	}

	body, err := json.Marshal(send)
	if err != nil {
		return err
	}

	err = chn.Publish(
		"sell_flow", // exchange
		"",          // routing key
		false,       // mandatory
		false,       // immediate
		amqp.Publishing{
			Body: []byte(body),
		})
	if err != nil {
		chn = nil
		return err
	}

	log.Output(1, "Rabbit order placed enviado")
	return nil
}

type orderPlacedData struct {
	OrderId string `json:"orderId"`

	CartId string `json:"cartId"`

	Articles []articlePlacedData `json:"articles"`
}

type articlePlacedData struct {
	ArticleId string `json:"articleId"`

	Quantity int `json:"quantity"`
}

func toPlaceData(event *events.Event) *orderPlacedData {

	articles := make([]articlePlacedData, len(event.PlaceEvent.Articles))
	for index, article := range event.PlaceEvent.Articles {
		articles[index] = articlePlacedData{
			ArticleId: article.ArticleId,
			Quantity:  article.Quantity,
		}
	}

	return &orderPlacedData{
		OrderId:  event.OrderId,
		CartId:   event.PlaceEvent.CartId,
		Articles: articles,
	}
}

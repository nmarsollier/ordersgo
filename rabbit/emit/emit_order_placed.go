package emit

import (
	"encoding/json"

	"github.com/golang/glog"
	"github.com/nmarsollier/ordersgo/events"
	"github.com/streadway/amqp"
)

//	@Summary		Emite order_placed/order_placed
//	@Description	Emite order_placed, un broadcast a rabbit con order_placed. Esto no es Rest es RabbitMQ.
//	@Tags			Rabbit
//	@Accept			json
//	@Produce		json
//	@Param			body	body	message	true	"Order Placed Event"
//	@Router			/rabbit/order_placed [put]
//
// SendOrderPlaced envía un broadcast a rabbit con logout
func EmitOrderPlaced(data *events.Event) error {
	send := message{
		Message: *toPlaceData(data),
	}

	chn, err := getChannel()
	if err != nil {
		glog.Error(err)
		chn = nil
		return err
	}

	err = chn.ExchangeDeclare(
		"order_placed", // name
		"fanout",       // type
		false,          // durable
		false,          // auto-deleted
		false,          // internal
		false,          // no-wait
		nil,            // arguments
	)
	if err != nil {
		glog.Error(err)
		chn = nil
		return err
	}

	body, err := json.Marshal(send)
	if err != nil {
		glog.Error(err)
		return err
	}

	err = chn.Publish(
		"order_placed", // exchange
		"",             // routing key
		false,          // mandatory
		false,          // immediate
		amqp.Publishing{
			Body: []byte(body),
		})
	if err != nil {
		glog.Error(err)
		chn = nil
		return err
	}

	glog.Info("Emit order_placed :", string(body))
	return nil
}

type message struct {
	Message orderPlacedData `json:"message"`
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

package r_emit

import (
	"encoding/json"

	"github.com/golang/glog"
	"github.com/nmarsollier/ordersgo/events"
	"github.com/streadway/amqp"
)

// SendOrderPlaced envía un broadcast a rabbit con logout
//
//	@Summary		Mensage Rabbit
//	@Description	SendOrderPlaced envía un broadcast a rabbit con logout. Esto no es Rest es RabbitMQ.
//	@Tags			Rabbit
//	@Accept			json
//	@Produce		json
//	@Param			body	body	message	true	"Order Placed Event"
//	@Router			/rabbit/logout [put]
func EmitOrderPlaced(data *events.Event) error {
	send := message{
		Type:     "order-placed",
		Exchange: "",
		Queue:    "",
		Message:  *toPlaceData(data),
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

	glog.Info("Rabbit order placed enviado ", string(body))
	return nil
}

type message struct {
	Type     string          `json:"type"`
	Exchange string          `json:"exchange"`
	Queue    string          `json:"queue"`
	Message  orderPlacedData `json:"message"`
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

package emit

import (
	"encoding/json"

	"github.com/nmarsollier/ordersgo/events"
	"github.com/nmarsollier/ordersgo/log"
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
func EmitOrderPlaced(data *events.Event, ctx ...interface{}) error {
	logger := log.Get(ctx...).
		WithField(log.LOG_FIELD_CONTROLLER, "Rabbit").
		WithField(log.LOG_FIELD_RABBIT_EXCHANGE, "order_placed").
		WithField(log.LOG_FIELD_RABBIT_QUEUE, "order_placed").
		WithField(log.LOG_FIELD_RABBIT_ACTION, "Emit")

	corrId, _ := logger.Data[log.LOG_FIELD_CORRELATION_ID].(string)
	send := message{
		CorrelationId: corrId,
		Message:       *toPlaceData(data),
	}

	chn, err := getChannel(logger)
	if err != nil {
		logger.Error(err)
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
		logger.Error(err)
		chn = nil
		return err
	}

	body, err := json.Marshal(send)
	if err != nil {
		logger.Error(err)
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
		logger.Error(err)
		chn = nil
		return err
	}

	logger.Info(string(body))
	return nil
}

type message struct {
	CorrelationId string          `json:"correlation_id" example:"123123" `
	Message       orderPlacedData `json:"message"`
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

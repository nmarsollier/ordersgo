package emit

import (
	"encoding/json"

	"github.com/nmarsollier/ordersgo/internal/engine/log"
	"github.com/nmarsollier/ordersgo/internal/events"
)

type RabbitEmit interface {
	EmitArticleValidation(data ArticleValidationData) error
	EmitOrderPlaced(data *events.Event) error
}

func NewRabbitEmit(log log.LogRusEntry, channel RabbitChannel) RabbitEmit {
	return &rabbitEmit{
		log:     log,
		channel: channel,
	}
}

type rabbitEmit struct {
	log     log.LogRusEntry
	channel RabbitChannel
}

//	@Summary		Emite article_exist/article_exist
//	@Description	Antes de iniciar las operaciones se validan los artículos contra el catalogo.
//	@Tags			Rabbit
//	@Accept			json
//	@Produce		json
//	@Param			body	body	SendValidationMessage	true	"Mensage de validacion"
//	@Router			/rabbit/cart/article_exist [put]
//
// Emite Validar Artículos a Cart
func (r *rabbitEmit) EmitArticleValidation(data ArticleValidationData) error {
	r.log.
		WithField(log.LOG_FIELD_CONTROLLER, "Rabbit").
		WithField(log.LOG_FIELD_RABBIT_EXCHANGE, "article_exist").
		WithField(log.LOG_FIELD_RABBIT_QUEUE, "article_exist").
		WithField(log.LOG_FIELD_RABBIT_ACTION, "Emit")

	corrId, _ := r.log.Data()[log.LOG_FIELD_CORRELATION_ID].(string)

	send := SendValidationMessage{
		CorrelationId: corrId,
		Exchange:      "article_exist",
		RoutingKey:    "order_article_exist",
		Message:       data,
	}

	err := r.channel.ExchangeDeclare(
		"article_exist", // name
		"direct",        // type
	)
	if err != nil {
		r.log.Error(err)
		return err
	}

	body, err := json.Marshal(send)
	if err != nil {
		r.log.Error(err)
		return err
	}

	err = r.channel.Publish(
		"article_exist", // exchange
		"article_exist", // routing key
		body)
	if err != nil {
		r.log.Error(err)
		return err
	}

	r.log.Info(string(body))

	return nil
}

type ArticleValidationData struct {
	ReferenceId string `json:"referenceId"`

	ArticleId string `json:"articleId"`
}

type SendValidationMessage struct {
	CorrelationId string                `json:"correlation_id" example:"123123" `
	Exchange      string                `json:"exchange"`
	RoutingKey    string                `json:"routing_key" example:"Remote RoutingKey to Reply"`
	Message       ArticleValidationData `json:"message"`
}

//	@Summary		Emite order_placed/order_placed
//	@Description	Emite order_placed, un broadcast a rabbit con order_placed. Esto no es Rest es RabbitMQ.
//	@Tags			Rabbit
//	@Accept			json
//	@Produce		json
//	@Param			body	body	message	true	"Order Placed Event"
//	@Router			/rabbit/order_placed [put]
//
// SendOrderPlaced envía un broadcast a rabbit con logout
func (r *rabbitEmit) EmitOrderPlaced(data *events.Event) error {
	r.log.
		WithField(log.LOG_FIELD_CONTROLLER, "Rabbit").
		WithField(log.LOG_FIELD_RABBIT_EXCHANGE, "order_placed").
		WithField(log.LOG_FIELD_RABBIT_QUEUE, "order_placed").
		WithField(log.LOG_FIELD_RABBIT_ACTION, "Emit")

	corrId, _ := r.log.Data()[log.LOG_FIELD_CORRELATION_ID].(string)
	send := message{
		CorrelationId: corrId,
		Message:       *toPlaceData(data),
	}

	err := r.channel.ExchangeDeclare(
		"order_placed", // name
		"fanout",       // type
	)
	if err != nil {
		r.log.Error(err)
		return err
	}

	body, err := json.Marshal(send)
	if err != nil {
		r.log.Error(err)
		return err
	}

	err = r.channel.Publish(
		"order_placed", // exchange
		"",             // routing key
		body)
	if err != nil {
		r.log.Error(err)
		return err
	}

	r.log.Info(string(body))
	return nil
}

type message struct {
	CorrelationId string          `json:"correlation_id" example:"123123" `
	Message       orderPlacedData `json:"message"`
}

type orderPlacedData struct {
	OrderId string `json:"orderId"`

	CartId string `json:"cartId"`

	UserId string `json:"userId"`

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
		UserId:   event.PlaceEvent.UserId,
		Articles: articles,
	}
}

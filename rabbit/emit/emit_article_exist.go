package emit

import (
	"encoding/json"

	"github.com/nmarsollier/ordersgo/tools/log"
	"github.com/streadway/amqp"
)

//	@Summary		Emite article_exist/article_exist
//	@Description	Antes de iniciar las operaciones se validan los artículos contra el catalogo.
//	@Tags			Rabbit
//	@Accept			json
//	@Produce		json
//	@Param			body	body	SendValidationMessage	true	"Mensage de validacion"
//	@Router			/rabbit/cart/article_exist [put]
//
// Emite Validar Artículos a Cart
func EmitArticleValidation(data ArticleValidationData, ctx ...interface{}) error {
	logger := log.Get(ctx...).
		WithField(log.LOG_FIELD_CONTROLLER, "Rabbit").
		WithField(log.LOG_FIELD_RABBIT_EXCHANGE, "article_exist").
		WithField(log.LOG_FIELD_RABBIT_QUEUE, "article_exist").
		WithField(log.LOG_FIELD_RABBIT_ACTION, "Emit")

	corrId, _ := logger.Data[log.LOG_FIELD_CORRELATION_ID].(string)

	send := SendValidationMessage{
		CorrelationId: corrId,
		Exchange:      "article_exist",
		RoutingKey:    "order_article_exist",
		Message:       data,
	}

	chn, err := getChannel(logger)
	if err != nil {
		logger.Error(err)
		chn = nil
		return err
	}

	err = chn.ExchangeDeclare(
		"article_exist", // name
		"direct",        // type
		false,           // durable
		false,           // auto-deleted
		false,           // internal
		false,           // no-wait
		nil,             // arguments
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
		"article_exist", // exchange
		"article_exist", // routing key
		false,           // mandatory
		false,           // immediate
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

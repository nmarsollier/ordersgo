package rabbit

import (
	"encoding/json"
	"log"

	"github.com/streadway/amqp"
)

type ArticleValidationData struct {
	ReferenceId string `json:"referenceId"`

	ArticleId string `json:"articleId"`
}

/**
 *
 * @api {direct} cart/article-data Validación de Artículos
 *
 * @apiGroup RabbitMQ POST
 *
 * @apiDescription Antes de iniciar las operaciones se validan los artículos contra el catalogo.
 *
 * @apiSuccessExample {json} Mensaje
 *     {
 *     "type": "article-data",
 *     "message" : {
 *         "cartId": "{cartId}",
 *         "articleId": "{articleId}",
 *        }
 *     }
 */
func SendArticleValidation(data ArticleValidationData) error {
	type message struct {
		Type     string                `json:"type"`
		Exchange string                `json:"exchange"`
		Queue    string                `json:"queue"`
		Message  ArticleValidationData `json:"message"`
	}

	send := message{
		Type:     "article-data",
		Exchange: "order",
		Queue:    "order",
		Message:  data,
	}

	chn, err := getChannel()
	if err != nil {
		chn = nil
		return err
	}

	err = chn.ExchangeDeclare(
		"catalog", // name
		"direct",  // type
		false,     // durable
		false,     // auto-deleted
		false,     // internal
		false,     // no-wait
		nil,       // arguments
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
		"catalog", // exchange
		"",        // routing key
		false,     // mandatory
		false,     // immediate
		amqp.Publishing{
			Body: []byte(body),
		})
	if err != nil {
		chn = nil
		return err
	}

	log.Output(1, "Rabbit article validation enviado")
	return nil
}

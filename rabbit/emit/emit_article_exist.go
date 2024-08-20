package emit

import (
	"encoding/json"

	"github.com/golang/glog"
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
func EmitArticleValidation(data ArticleValidationData) error {

	send := SendValidationMessage{
		Exchange:   "article_exist",
		RoutingKey: "order_article_exist",
		Message:    data,
	}

	chn, err := getChannel()
	if err != nil {
		glog.Error(err)
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
		"article_exist", // exchange
		"article_exist", // routing key
		false,           // mandatory
		false,           // immediate
		amqp.Publishing{
			Body: []byte(body),
		})
	if err != nil {
		glog.Error(err)
		chn = nil
		return err
	}

	glog.Info("Emit article_exist :", string(body))

	return nil
}

type ArticleValidationData struct {
	ReferenceId string `json:"referenceId"`

	ArticleId string `json:"articleId"`
}

type SendValidationMessage struct {
	Exchange   string                `json:"exchange"`
	RoutingKey string                `json:"routing_key" example:"Remote RoutingKey to Reply"`
	Message    ArticleValidationData `json:"message"`
}

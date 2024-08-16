package emit

import (
	"encoding/json"

	"github.com/golang/glog"
	"github.com/streadway/amqp"
)

// @Summary		Emite Validar Artículos a Cart cart/article-data
// @Description	Antes de iniciar las operaciones se validan los artículos contra el catalogo.
// @Tags			Rabbit
// @Accept			json
// @Produce		json
// @Param			body	body	SendValidationMessage	true	"Mensage de validacion"
// @Router			/rabbit/cart/article-data [put]
//
// Emite Validar Artículos a Cart
func EmitArticleValidation(data ArticleValidationData) error {

	send := SendValidationMessage{
		Type:     "article-data",
		Exchange: "order",
		Queue:    "order",
		Message:  data,
	}

	chn, err := getChannel()
	if err != nil {
		glog.Error(err)
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
		"catalog", // exchange
		"catalog", // routing key
		false,     // mandatory
		false,     // immediate
		amqp.Publishing{
			Body: []byte(body),
		})
	if err != nil {
		glog.Error(err)
		chn = nil
		return err
	}

	glog.Info("Rabbit article validation enviado ", string(body))

	return nil
}

type ArticleValidationData struct {
	ReferenceId string `json:"referenceId"`

	ArticleId string `json:"articleId"`
}

type SendValidationMessage struct {
	Type     string                `json:"type"`
	Exchange string                `json:"exchange"`
	Queue    string                `json:"queue"`
	Message  ArticleValidationData `json:"message"`
}
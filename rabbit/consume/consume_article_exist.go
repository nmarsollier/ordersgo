package consume

import (
	"encoding/json"

	"github.com/nmarsollier/ordersgo/events"
	"github.com/nmarsollier/ordersgo/services"
	"github.com/nmarsollier/ordersgo/tools/env"
	"github.com/nmarsollier/ordersgo/tools/log"
	uuid "github.com/satori/go.uuid"
	"github.com/streadway/amqp"
)

//	@Summary		Mensage Rabbit article_exist/order_article_exist
//	@Description	Antes de iniciar las operaciones se validan los artículos contra el catalogo.
//	@Tags			Rabbit
//	@Accept			json
//	@Produce		json
//	@Param			article_exist	body	consumeArticleDataMessage	true	"Consume article_exist/order_article_exist"
//	@Router			/rabbit/article_exist [get]
//
// Validar Artículos
func consumeArticleData() error {
	logger := log.Get().
		WithField(log.LOG_FIELD_CONTROLLER, "Rabbit").
		WithField(log.LOG_FIELD_RABBIT_EXCHANGE, "article_exist").
		WithField(log.LOG_FIELD_RABBIT_QUEUE, "order_article_exist").
		WithField(log.LOG_FIELD_RABBIT_ACTION, "Consume")

	conn, err := amqp.Dial(env.Get().RabbitURL)
	if err != nil {
		logger.Error(err)
		return err
	}
	defer conn.Close()

	chn, err := conn.Channel()
	if err != nil {
		logger.Error(err)
		return err
	}
	defer chn.Close()

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
		return err
	}

	queue, err := chn.QueueDeclare(
		"order_article_exist", // name
		false,                 // durable
		false,                 // delete when unused
		false,                 // exclusive
		false,                 // no-wait
		nil,                   // arguments
	)
	if err != nil {
		logger.Error(err)
		return err
	}

	err = chn.QueueBind(
		queue.Name,            // queue name
		"order_article_exist", // routing key
		"article_exist",       // exchange
		false,
		nil)
	if err != nil {
		logger.Error(err)
		return err
	}

	mgs, err := chn.Consume(
		queue.Name, // queue
		"",         // consumer
		false,      // auto-ack
		false,      // exclusive
		false,      // no-local
		false,      // no-wait
		nil,        // args
	)
	if err != nil {
		logger.Error(err)
		return err
	}

	logger.Info("RabbitMQ consumeOrdersChannel conectado")

	go func() {
		for d := range mgs {
			newMessage := &consumeArticleDataMessage{}
			body := d.Body

			err = json.Unmarshal(body, newMessage)
			if err == nil {
				l := logger.WithField(log.LOG_FIELD_CORRELATION_ID, getArticleExistCorrelationId(newMessage))
				l.Info("Incoming article_exist : ", string(body))

				processArticleData(newMessage, l)

				if err := d.Ack(false); err != nil {
					l.Info("Failed ACK :", string(body), err)
				} else {
					l.Info("Consumed article_exist :", string(body))
				}
			} else {
				logger.Error(err)
			}
		}
	}()

	logger.Info("Closed connection: ", <-conn.NotifyClose(make(chan *amqp.Error)))

	return nil
}

func processArticleData(newMessage *consumeArticleDataMessage, ctx ...interface{}) {
	data := newMessage.Message

	_, err := services.ProcessArticleData(data, ctx...)
	if err != nil {
		log.Get(ctx...).Error(err)
		return
	}
}

type consumeArticleDataMessage struct {
	CorrelationId string `json:"correlation_id" example:"123123" `
	Message       *events.ValidationEvent
}

func getArticleExistCorrelationId(c *consumeArticleDataMessage) string {
	value := c.CorrelationId

	if len(value) == 0 {
		value = uuid.NewV4().String()
	}

	return value
}

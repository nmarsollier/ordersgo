package consume

import (
	"encoding/json"

	"github.com/nmarsollier/ordersgo/internal/engine/log"
	"github.com/nmarsollier/ordersgo/internal/events"
	"github.com/nmarsollier/ordersgo/internal/services"
	uuid "github.com/satori/go.uuid"
	"github.com/streadway/amqp"
)

type ArticleExistConsumer interface {
	ConsumeArticleExist() error
}

func NewArticleExistConsumer(fluentUrl string, rabbitUrl string, service services.Service) ArticleExistConsumer {
	log := log.Get(fluentUrl).
		WithField(log.LOG_FIELD_CONTROLLER, "Rabbit").
		WithField(log.LOG_FIELD_RABBIT_EXCHANGE, "article_exist").
		WithField(log.LOG_FIELD_RABBIT_QUEUE, "order_article_exist").
		WithField(log.LOG_FIELD_RABBIT_ACTION, "Consume")

	return &articleExistConsumer{
		log:       log,
		rabbitUrl: rabbitUrl,
		service:   service,
	}
}

type articleExistConsumer struct {
	log       log.LogRusEntry
	rabbitUrl string
	service   services.Service
}

//	@Summary		Mensage Rabbit article_exist/order_article_exist
//	@Description	Antes de iniciar las operaciones se validan los artículos contra el catalogo.
//	@Tags			Rabbit
//	@Accept			json
//	@Produce		json
//	@Param			article_exist	body	consumeArticleDataMessage	true	"Consume article_exist/order_article_exist"
//	@Router			/rabbit/article_exist [get]
//
// Validar Artículos
func (r *articleExistConsumer) ConsumeArticleExist() error {

	conn, err := amqp.Dial(r.rabbitUrl)
	if err != nil {
		r.log.Error(err)
		return err
	}
	defer conn.Close()

	chn, err := conn.Channel()
	if err != nil {
		r.log.Error(err)
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
		r.log.Error(err)
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
		r.log.Error(err)
		return err
	}

	err = chn.QueueBind(
		queue.Name,            // queue name
		"order_article_exist", // routing key
		"article_exist",       // exchange
		false,
		nil)
	if err != nil {
		r.log.Error(err)
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
		r.log.Error(err)
		return err
	}

	r.log.Info("RabbitMQ consumeOrdersChannel conectado")

	go func() {
		for d := range mgs {
			newMessage := &consumeArticleDataMessage{}
			body := d.Body

			err = json.Unmarshal(body, newMessage)
			if err == nil {
				r.log.WithField(log.LOG_FIELD_CORRELATION_ID, r.getArticleExistCorrelationId(newMessage))
				r.log.Info("Incoming article_exist : ", string(body))

				r.processArticleData(newMessage)

				if err := d.Ack(false); err != nil {
					r.log.Info("Failed ACK :", string(body), err)
				} else {
					r.log.Info("Consumed article_exist :", string(body))
				}
			} else {
				r.log.Error(err)
			}
		}
	}()

	r.log.Info("Closed connection: ", <-conn.NotifyClose(make(chan *amqp.Error)))

	return nil
}

func (r *articleExistConsumer) processArticleData(newMessage *consumeArticleDataMessage, deps ...interface{}) {
	data := newMessage.Message

	_, err := r.service.ProcessArticleData(data)
	if err != nil {
		r.log.Error(err)
		return
	}
}

type consumeArticleDataMessage struct {
	CorrelationId string `json:"correlation_id" example:"123123" `
	Message       *events.ValidationEvent
}

func (r *articleExistConsumer) getArticleExistCorrelationId(c *consumeArticleDataMessage) string {
	value := c.CorrelationId

	if len(value) == 0 {
		value = uuid.NewV4().String()
	}

	return value
}

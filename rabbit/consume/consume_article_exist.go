package consume

import (
	"encoding/json"

	"github.com/golang/glog"
	"github.com/nmarsollier/ordersgo/events"
	"github.com/nmarsollier/ordersgo/services"
	"github.com/nmarsollier/ordersgo/tools/env"
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
	conn, err := amqp.Dial(env.Get().RabbitURL)
	if err != nil {
		glog.Error(err)
		return err
	}
	defer conn.Close()

	chn, err := conn.Channel()
	if err != nil {
		glog.Error(err)
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
		glog.Error(err)
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
		glog.Error(err)
		return err
	}

	err = chn.QueueBind(
		queue.Name,            // queue name
		"order_article_exist", // routing key
		"article_exist",       // exchange
		false,
		nil)
	if err != nil {
		glog.Error(err)
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
		glog.Error(err)
		return err
	}

	glog.Info("RabbitMQ consumeOrdersChannel conectado")

	go func() {
		for d := range mgs {
			newMessage := &consumeArticleDataMessage{}
			body := d.Body
			glog.Info("Incomming article_exist : ", string(body))

			err = json.Unmarshal(body, newMessage)
			if err == nil {
				processArticleData(newMessage)

				if err := d.Ack(false); err != nil {
					glog.Info("Failed ACK :", string(body), err)
				} else {
					glog.Info("Consumed article_exist :", string(body))
				}
			} else {
				glog.Error(err)
			}
		}
	}()

	glog.Info("Closed connection: ", <-conn.NotifyClose(make(chan *amqp.Error)))

	return nil
}

func processArticleData(newMessage *consumeArticleDataMessage) {
	data := newMessage.Message

	_, err := services.ProcessArticleData(data)
	if err != nil {
		glog.Error(err)
		return
	}
}

type consumeArticleDataMessage struct {
	Message *events.ValidationEvent
}

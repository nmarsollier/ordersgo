package consume

import (
	"time"

	"github.com/nmarsollier/ordersgo/log"
)

func Init() {
	logger := log.Get().
		WithField("Controller", "Rabbit").
		WithField("Method", "Consume")
	go func() {
		for {
			err := consumePlaceOrder()
			if err != nil {
				logger.Error(err)
			}
			logger.Info("RabbitMQ consumePlaceOrder conectando en 5 segundos.")
			time.Sleep(5 * time.Second)
		}
	}()

	go func() {
		for {
			err := consumeLogout()
			if err != nil {
				logger.Error(err)
			}
			logger.Info("RabbitMQ listenLogout conectando en 5 segundos.")
			time.Sleep(5 * time.Second)
		}
	}()

	go func() {
		for {
			err := consumeArticleData()
			if err != nil {
				logger.Error(err)
			}
			logger.Info("RabbitMQ consumeArticleData conectando en 5 segundos.")
			time.Sleep(5 * time.Second)
		}
	}()
}

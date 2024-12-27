package rabbit

import (
	"time"

	"github.com/nmarsollier/ordersgo/internal/engine/di"
	"github.com/nmarsollier/ordersgo/internal/engine/log"
)

func Init(deps di.Injector) {
	logger := deps.Logger().
		WithField(log.LOG_FIELD_CONTROLLER, "Rabbit").
		WithField(log.LOG_FIELD_RABBIT_ACTION, "Init")
	go func() {
		for {
			err := deps.OrderPlacedConsumer().ConsumeOrderPlaced()
			if err != nil {
				logger.Error(err)
			}
			logger.Info("RabbitMQ consumePlaceOrder conectando en 5 segundos.")
			time.Sleep(5 * time.Second)
		}
	}()

	go func() {
		for {
			err := deps.LogoutConsumer().ConsumeLogout()
			if err != nil {
				logger.Error(err)
			}
			logger.Info("RabbitMQ listenLogout conectando en 5 segundos.")
			time.Sleep(5 * time.Second)
		}
	}()

	go func() {
		for {
			err := deps.ArticleExistConsumer().ConsumeArticleExist()
			if err != nil {
				logger.Error(err)
			}
			logger.Info("RabbitMQ consumeArticleData conectando en 5 segundos.")
			time.Sleep(5 * time.Second)
		}
	}()
}

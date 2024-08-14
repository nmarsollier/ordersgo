package r_consume

import (
	"time"

	"github.com/golang/glog"
)

func Init() {
	go func() {
		for {
			err := consumePlaceOrder()
			if err != nil {
				glog.Error(err)
			}
			glog.Info("RabbitMQ consumePlaceOrder conectando en 5 segundos.")
			time.Sleep(5 * time.Second)
		}
	}()

	go func() {
		for {
			err := consumeLogout()
			if err != nil {
				glog.Error(err)
			}
			glog.Info("RabbitMQ listenLogout conectando en 5 segundos.")
			time.Sleep(5 * time.Second)
		}
	}()

	go func() {
		for {
			err := consumeArticleData()
			if err != nil {
				glog.Error(err)
			}
			glog.Info("RabbitMQ consumeArticleData conectando en 5 segundos.")
			time.Sleep(5 * time.Second)
		}
	}()
}

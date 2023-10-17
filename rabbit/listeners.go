package rabbit

import (
	"fmt"
	"time"
)

func Init() {
	go func() {
		for {
			consumeOrdersChannel()
			fmt.Println("RabbitMQ consumeOrdersChannel conectando en 5 segundos.")
			time.Sleep(5 * time.Second)
		}
	}()

	go func() {
		for {
			listenLogout()
			fmt.Println("RabbitMQ listenLogout conectando en 5 segundos.")
			time.Sleep(5 * time.Second)
		}
	}()
}

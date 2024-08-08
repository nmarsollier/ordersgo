package rabbit

import (
	"fmt"
	"time"
)

func Init() {
	go func() {
		for {
			err := consumeOrdersChannel()
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println("RabbitMQ consumeOrdersChannel conectando en 5 segundos.")
			time.Sleep(5 * time.Second)
		}
	}()

	go func() {
		for {
			err := listenLogout()
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println("RabbitMQ listenLogout conectando en 5 segundos.")
			time.Sleep(5 * time.Second)
		}
	}()
}

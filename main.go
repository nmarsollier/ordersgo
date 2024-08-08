package main

import (
	"github.com/nmarsollier/ordersgo/rabbit"
	routes "github.com/nmarsollier/ordersgo/rest"
)

//	@title			OrdersGo
//	@version		1.0
//	@description	Microservicio de Ordenes.

//	@contact.name	Nestor Marsollier
//	@contact.email	nmarsollier@gmail.com

// @host		localhost:3004
// @BasePath	/v1
func main() {
	rabbit.Init()
	routes.Start()
}

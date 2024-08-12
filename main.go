package main

import (
	"flag"

	"github.com/nmarsollier/ordersgo/rabbit/r_consume"
	routes "github.com/nmarsollier/ordersgo/rest"
)

//	@title			OrdersGo
//	@version		1.0
//	@description	Microservicio de Ordenes.
//	@contact.name	Nestor Marsollier
//	@contact.email	nmarsollier@gmail.com
//
//	@host			localhost:3004
//	@BasePath		/v1
//
// Main Method
func main() {
	// For logging
	flag.Parse()
	flag.Set("logtostderr", "true")
	flag.Set("v", "2")

	r_consume.Init()
	routes.Start()
}

package main

import (
	"flag"

	"github.com/golang/glog"
	"github.com/nmarsollier/ordersgo/rabbit/consume"
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
	flag.Set("stderrthreshold", "ERROR")
	flag.Set("v", "2")
	defer glog.Flush()

	consume.Init()
	routes.Start()
}

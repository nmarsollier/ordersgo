package main

import (
	"github.com/nmarsollier/ordersgo/rabbit"
	routes "github.com/nmarsollier/ordersgo/rest"
)

func main() {
	rabbit.Init()
	routes.Start()
}

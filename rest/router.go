package rest

import (
	"fmt"

	"github.com/nmarsollier/ordersgo/rest/server"
	"github.com/nmarsollier/ordersgo/tools/env"
)

// Start this server
func Start() {
	server.Router().Run(fmt.Sprintf(":%d", env.Get().Port))
}

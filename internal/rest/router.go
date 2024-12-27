package rest

import (
	"fmt"

	"github.com/nmarsollier/ordersgo/internal/engine/env"
	"github.com/nmarsollier/ordersgo/internal/rest/engine"
)

// Start this server
func Start() {
	engine.Router().Run(fmt.Sprintf(":%d", env.Get().Port))
}

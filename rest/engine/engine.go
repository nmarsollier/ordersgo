package engine

import (
	"time"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
	_ "github.com/nmarsollier/ordersgo/docs"
	"github.com/nmarsollier/ordersgo/rest/middlewares"
	"github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var engine *gin.Engine = nil

func Router() *gin.Engine {
	if engine == nil {

		engine = gin.Default()

		engine.Use(gzip.Gzip(gzip.DefaultCompression))

		engine.Use(cors.Middleware(cors.Config{
			Origins:         "*",
			Methods:         "GET, PUT, POST, DELETE",
			RequestHeaders:  "Origin, Authorization, Content-Type",
			ExposedHeaders:  "",
			MaxAge:          50 * time.Second,
			Credentials:     false,
			ValidateHeaders: false,
		}))

		engine.Use(middlewares.ErrorHandler)

		engine.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	return engine
}

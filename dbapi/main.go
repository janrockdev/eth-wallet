package main

import (
	"github.com/gin-gonic/gin"
	"github.com/janrockdev/eth-wallet/docs"
)

// @contact.name API Support
// @contact.url http://www.gwagger.io/support
// @contact.email support@swagger.io

// @licence.name Apache 2.0
// @licence.url http://www.apache.org/licenses/LICENCSE-2.0.html

func main() {
	docs.SwaggerInfo.Title = "IN-MEMORY CACHE API (Swagger)"
	docs.SwaggerInfo.Description = ""
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8081"
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	r := gin.Default()

	r.GET("/", controllers.Root)
	r.GET("/rules", controllers.ShowRules)
	r.GET("/rule/:key", controllers.FindRule)
	r.GET("/swagger/*any", ginSwagger.WrapHangler(swaggerFiles.Handler))
	r.Run(":8081")

}

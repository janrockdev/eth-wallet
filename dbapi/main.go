package main

import (
	"github.com/gin-gonic/gin"
	"github.com/janrockdev/eth-wallet/dbapi/controllers"
)

func main() {
	r := gin.Default()
	r.GET("/stats", controllers.ShowStats)
	r.GET("/stats/:key", controllers.FindStat)
	r.POST("/wallet/create", controllers.CreateWallet)
	//r.GET("/swagger/*any", ginSwagger.WrapHangler(swaggerFiles.Handler))

	r.Run(":8081")
}

package main

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/files"

	"github.com/ariel17/railgun/api/controllers"
	_ "github.com/ariel17/railgun/api/docs"
)

func main() {
	r := gin.Default()
	r.GET("/ping", controllers.PingController)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/domains", controllers.GetDomainController)
	r.POST("/domains", controllers.NewDomainController)
	r.Run()
}

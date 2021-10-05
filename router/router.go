package router

import (
	"github.com/gin-gonic/gin"
	"joke/controller"
)

func RouterInit() *gin.Engine {
	r := gin.Default()
	r.Static("/static", "static")
	r.LoadHTMLGlob("templates/*")

	r.GET("/", controller.GetRandom)

	r.GET("/add.html", controller.AddHtml)

	r.POST("/add", controller.Add)

	return r
}

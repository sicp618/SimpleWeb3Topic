package main

import (
	"SimpleBlog/controller"
	"SimpleBlog/model"
	"github.com/gin-gonic/gin"
)

func main() {
	model.Init()
	controller.Init()

	r := gin.Default()
	r.GET("/ping", controller.Ping)
	r.POST("/login", controller.Login)
	r.POST("/register", controller.Register)
	r.Run()
}

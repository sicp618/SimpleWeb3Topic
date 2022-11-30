package main

import (
	"SimpleBlog/controller"
	"SimpleBlog/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"sync"
)

func main() {
	model.Init()
	controller.Init()

	r := gin.Default()
	r.GET("/ping", controller.Ping)
	r.POST("/login", controller.Login)
	r.GET("/login", controller.LoginGet)
	r.OPTIONS("/login", controller.LoginOptions)
	r.POST("/register", controller.Register)
	r.Run()
}

var wg sync.WaitGroup

func main1() {
	chDog := make(chan struct{})
	chCat := make(chan struct{})
	chFish := make(chan struct{})
	fmt.Println("start")
	chCat <- struct{}{}
	fmt.Println("end")
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go cat(chCat, chDog)
		go dog(chDog, chFish)
		go fish(chFish, chCat)
		fmt.Println("create gorount ", i)
	}
	wg.Wait()
}

func dog(ch, chNext chan struct{}) {
	<-ch
	fmt.Println("dog")
	chNext <- struct{}{}
}
func cat(ch, chNext chan struct{}) {
	<-ch
	fmt.Println("cat")
	chNext <- struct{}{}
}
func fish(ch, chNext chan struct{}) {
	<-ch
	fmt.Println("fish")
	chNext <- struct{}{}
	wg.Done()
}

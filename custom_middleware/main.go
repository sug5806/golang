package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang/custom_middleware/handler"
	"golang/custom_middleware/middleware"
)

var myHandler handler.MyHandler

func main() {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(middleware.CM)

	r.GET("/ping", myHandler.GetPing)

	r.Run("127.0.0.1:7089")
}

func init() {
	fmt.Println("init init")
	myHandler = handler.MyHandler{}
}

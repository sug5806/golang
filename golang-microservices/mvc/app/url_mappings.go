package app

import (
	"golang/golang-microservices/mvc/controller"
)

func mapUrls() {
	router.GET("/users/:user_id", controller.GetUser)
}

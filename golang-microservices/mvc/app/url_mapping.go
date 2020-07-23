package app

import "golang/golang-microservices/mvc/controller"

func urlMapping() {
	router.GET("/users/:user_id", controller.GetUser)
}

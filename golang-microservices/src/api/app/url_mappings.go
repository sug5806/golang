package app

import (
	"golang/golang-microservices/src/api/controllers/polo"
	"golang/golang-microservices/src/api/controllers/repositories"
)

func mapUrls() {
	router.GET("/marco", polo.Marco)
	router.POST("/repositories", repositories.CreateRepo)
}

package services

import (
	"golang/golang-microservices/mvc/domain"
	"golang/golang-microservices/mvc/utils"
)

func GetUser(userId int64) (*domain.User, *utils.ApplicationError) {
	return domain.GetUser(userId)
}

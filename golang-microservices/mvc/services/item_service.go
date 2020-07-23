package services

import (
	"golang/golang-microservices/mvc/domain"
	"golang/golang-microservices/mvc/utils"
	"net/http"
)

type itemService struct {
}

var (
	ItemService itemService
)

func (i *itemService) GetItem(itemId string) (*domain.Item, *utils.ApplicationError) {
	return nil, &utils.ApplicationError{
		Message:    "implement me",
		StatusCode: http.StatusInternalServerError,
		Code:       "",
	}
}

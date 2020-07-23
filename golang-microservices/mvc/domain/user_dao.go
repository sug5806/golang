package domain

import (
	"fmt"
	"golang/golang-microservices/mvc/utils"
	"log"
	"net/http"
)

var (
	users = map[int64]*User{
		123: {
			Id:        123,
			FirstName: "Fede",
			LastName:  "Leon",
			Email:     "myemail@gmail.com",
		},
	}

	UserDao userDaoInterface
)

func init() {
	UserDao = &userDao{}
}

type userDaoInterface interface {
	GetUser(int64) (*User, *utils.ApplicationError)
}

type userDao struct{}

func (u *userDao) GetUser(userId int64) (*User, *utils.ApplicationError) {
	log.Println("we're accessing the database")
	if user := users[userId]; user != nil {
		return user, nil
	}

	return nil, &utils.ApplicationError{
		Message:    fmt.Sprintf("user %v does not exists", userId),
		StatusCode: http.StatusBadRequest,
		Code:       "bad_request",
	}
}

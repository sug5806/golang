package controller

import (
	"github.com/gin-gonic/gin"
	"golang/golang-microservices/mvc/services"
	"golang/golang-microservices/mvc/utils"
	"net/http"
	"strconv"
)

func GetUser(c *gin.Context) {
	userId, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		apiErr := &utils.ApplicationError{
			Message:    "user_id must be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}
		utils.ErrorResponse(c, apiErr)
		return

	}

	user, apiErr := services.UserService.GetUser(userId)
	if apiErr != nil {
		utils.ErrorResponse(c, apiErr)
		// Handle the err and return to the client
		return
	}

	// return user to client

	utils.Response(c, user)
}

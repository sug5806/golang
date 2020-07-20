package controller

import (
	"encoding/json"
	"golang/golang-microservices/mvc/services"
	"golang/golang-microservices/mvc/utils"
	"net/http"
	"strconv"
)

func GetUser(resp http.ResponseWriter, req *http.Request) {
	userId, err := strconv.ParseInt(req.URL.Query().Get("user_id"), 10, 64)
	if err != nil {
		apiErr := &utils.ApplicationError{
			Message:    "user_id must be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}

		jsonValue, _ := json.Marshal(apiErr)

		resp.WriteHeader(apiErr.StatusCode)
		resp.Write(jsonValue)
		return

	}

	user, apiErr := services.GetUser(userId)
	if apiErr != nil {
		resp.WriteHeader(apiErr.StatusCode)
		resp.Write([]byte(apiErr.Message))
		// Handle the err and return to the client
		return
	}

	// return user to client

	jsonValue, _ := json.Marshal(user)
	resp.Write(jsonValue)
}

package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type MyHandler struct {
}

type response struct {
	Code    uint8  `json:"code"`
	message string `json:"message"`
}

func (m *MyHandler) GetPing(c *gin.Context) {
	resp := &response{
		Code:    0,
		message: "pong pong",
	}
	c.PureJSON(http.StatusOK, resp)
	return
}

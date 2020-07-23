package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Response(c *gin.Context, body interface{}) {
	if c.GetHeader("Accept") == "application/xml" {
		c.XML(http.StatusOK, body)
		return
	}
	c.JSON(http.StatusOK, body)
}

func ErrorResponse(c *gin.Context, applicationError *ApplicationError) {
	if c.GetHeader("Accept") == "application/xml" {
		c.XML(applicationError.StatusCode, applicationError)
		return
	}
	c.JSON(applicationError.StatusCode, applicationError)
}

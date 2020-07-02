package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func CM(c *gin.Context) {
	fmt.Println("dummy dummy CM!!!")
	c.Next()
}

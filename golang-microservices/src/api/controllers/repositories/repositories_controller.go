package repositories

import (
	"github.com/gin-gonic/gin"
	"golang/golang-microservices/src/api/domain/repositories"
	"golang/golang-microservices/src/api/services"
	"golang/golang-microservices/src/api/utils/errors"
	"net/http"
)

func CreateRepo(c *gin.Context) {
	var request repositories.CreateRepoRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		apiErr := errors.NewBadRequestError("invalid json body")
		c.JSON(apiErr.ApiStatus(), apiErr)
		return
	}

	result, err := services.RepositoryService.CreateRepo(request)

	if err != nil {
		c.JSON(err.ApiStatus(), err)
		return
	}

	c.JSON(http.StatusCreated, result)
}

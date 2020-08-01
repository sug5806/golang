package services

import (
	"github.com/stretchr/testify/assert"
	"golang/golang-microservices/src/api/clients/rest_client"
	"golang/golang-microservices/src/api/domain/repositories"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"testing"
)

func TestMain(m *testing.M) {
	rest_client.StartMockups()
	os.Exit(m.Run())
}

func TestReposService_CreateRepoInvalidInputName(t *testing.T) {
	request := repositories.CreateRepoRequest{}
	result, err := RepositoryService.CreateRepo(request)

	assert.Nil(t, result)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusBadRequest, err.ApiStatus())
	assert.EqualValues(t, "invalid repository name", err.ApiMessage())

}

func TestReposService_CreateRepoErrorFromGithub(t *testing.T) {
	rest_client.FlushMockups()
	rest_client.AddMockup(rest_client.Mock{
		Url:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusUnauthorized,
			Body:       ioutil.NopCloser(strings.NewReader(`{"message": "Requires authentication", "documentation_url": "https://developer.githubcom/v3/repos/docs"}`)),
		},
	})

	request := repositories.CreateRepoRequest{
		Name:        "testing",
		Description: "",
	}

	result, err := RepositoryService.CreateRepo(request)
	assert.Nil(t, result)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusUnauthorized, err.ApiStatus())
	assert.EqualValues(t, "Requires authentication", err.ApiMessage())
}

func TestReposService_CreateRepoNoErrorFromGithub(t *testing.T) {
	rest_client.FlushMockups()
	rest_client.AddMockup(rest_client.Mock{
		Url:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusCreated,
			Body:       ioutil.NopCloser(strings.NewReader(`{"id": 123, "name": "testing", "owner": {"login": "federicoleon"}}`)),
		},
	})

	request := repositories.CreateRepoRequest{
		Name:        "testing",
		Description: "",
	}

	result, err := RepositoryService.CreateRepo(request)
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.EqualValues(t, 123, result.Id)
	assert.EqualValues(t, "testing", result.Name)
	assert.EqualValues(t, "federicoleon", result.Owner)
}

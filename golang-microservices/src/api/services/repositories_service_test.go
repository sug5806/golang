package services

import (
	"github.com/stretchr/testify/assert"
	"golang/golang-microservices/src/api/clients/rest_client"
	"golang/golang-microservices/src/api/domain/repositories"
	"golang/golang-microservices/src/api/utils/errors"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"sync"
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

func TestCreateRepoConcurrentInvalidRequest(t *testing.T) {
	request := repositories.CreateRepoRequest{}

	output := make(chan repositories.CreateRepositoriesResult)

	service := reposService{}
	go service.createRepoConcurrent(request, output)

	result := <-output
	assert.NotNil(t, result)
	assert.Nil(t, result.Response)
	assert.NotNil(t, result.Error)
	assert.EqualValues(t, http.StatusBadRequest, result.Error.ApiStatus())
	assert.EqualValues(t, "invalid repository name", result.Error.ApiMessage())
}

func TestCreateRepoConcurrentErrorFromGithub(t *testing.T) {
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
		Name:        "golang",
		Description: "test",
	}

	output := make(chan repositories.CreateRepositoriesResult)

	service := reposService{}
	go service.createRepoConcurrent(request, output)

	result := <-output
	assert.NotNil(t, result)
	assert.Nil(t, result.Response)
	assert.NotNil(t, result.Error)
	assert.EqualValues(t, http.StatusUnauthorized, result.Error.ApiStatus())
	assert.EqualValues(t, "Requires authentication", result.Error.ApiMessage())
}

func TestCreateRepoConcurrentNoError(t *testing.T) {
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
		Name:        "golang",
		Description: "test",
	}

	output := make(chan repositories.CreateRepositoriesResult)

	service := reposService{}
	go service.createRepoConcurrent(request, output)

	result := <-output
	assert.NotNil(t, result)
	assert.Nil(t, result.Error)
	assert.NotNil(t, result.Response)
	assert.EqualValues(t, 123, result.Response.Id)
	assert.EqualValues(t, "testing", result.Response.Name)
	assert.EqualValues(t, "federicoleon", result.Response.Owner)
}

func TestHandleRepoResults(t *testing.T) {
	input := make(chan repositories.CreateRepositoriesResult)
	output := make(chan repositories.CreateReposResponse)

	var wg sync.WaitGroup

	service := reposService{}
	wg.Add(1)

	go func() {
		input <- repositories.CreateRepositoriesResult{
			Error: errors.NewBadRequestError("invalid repository name"),
		}
	}()

	go service.handleRepoResults(&wg, input, output)

	wg.Wait()
	close(input)

	result := <-output

	assert.NotNil(t, result)
	assert.EqualValues(t, 0, result.StatusCode)

	assert.EqualValues(t, 1, len(result.Results))
	assert.NotNil(t, result.Results[0].Error)
	assert.EqualValues(t, http.StatusBadRequest, result.Results[0].Error.ApiStatus())
	assert.EqualValues(t, "invalid repository name", result.Results[0].Error.ApiMessage())

}

func TestCreateReposInvalidRequests(t *testing.T) {
	requests := []repositories.CreateRepoRequest{
		{},
		{Name: "  "},
	}

	result, err := RepositoryService.CreateRepos(requests)
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.EqualValues(t, http.StatusBadRequest, result.StatusCode)
	assert.EqualValues(t, 2, len(result.Results))

	assert.Nil(t, result.Results[0].Response)
	assert.EqualValues(t, http.StatusBadRequest, result.Results[0].Error.ApiStatus())
	assert.EqualValues(t, "invalid repository name", result.Results[0].Error.ApiMessage())

	assert.Nil(t, result.Results[1].Response)
	assert.EqualValues(t, http.StatusBadRequest, result.Results[1].Error.ApiStatus())
	assert.EqualValues(t, "invalid repository name", result.Results[1].Error.ApiMessage())

}

func TestCreateReposOneSuccessOneFail(t *testing.T) {
	rest_client.FlushMockups()
	rest_client.AddMockup(rest_client.Mock{
		Url:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusCreated,
			Body:       ioutil.NopCloser(strings.NewReader(`{"id": 123, "name": "testing", "owner": {"login": "federicoleon"}}`)),
		},
	})

	requests := []repositories.CreateRepoRequest{
		{},
		{Name: "testing"},
	}

	result, err := RepositoryService.CreateRepos(requests)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.EqualValues(t, http.StatusPartialContent, result.StatusCode)
	assert.EqualValues(t, 2, len(result.Results))

	for _, result := range result.Results {
		if result.Error != nil {
			assert.EqualValues(t, http.StatusBadRequest, result.Error.ApiStatus())
			assert.EqualValues(t, "invalid repository name", result.Error.ApiMessage())
		} else {
			assert.EqualValues(t, 123, result.Response.Id)
			assert.EqualValues(t, "testing", result.Response.Name)
			assert.EqualValues(t, "federicoleon", result.Response.Owner)
		}

	}

}

func TestCreateReposAllSuccess(t *testing.T) {
	rest_client.FlushMockups()
	rest_client.AddMockup(rest_client.Mock{
		Url:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusCreated,
			Body:       ioutil.NopCloser(strings.NewReader(`{"id": 123, "name": "testing", "owner": {"login": "federicoleon"}}`)),
		},
	})

	requests := []repositories.CreateRepoRequest{
		{Name: "testing"},
		{Name: "testing"},
	}

	result, err := RepositoryService.CreateRepos(requests)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.EqualValues(t, http.StatusCreated, result.StatusCode)
	assert.EqualValues(t, 2, len(result.Results))

	assert.Nil(t, result.Results[0].Error)
	assert.EqualValues(t, 123, result.Results[0].Response.Id)
	assert.EqualValues(t, "testing", result.Results[0].Response.Name)
	assert.EqualValues(t, "federicoleon", result.Results[0].Response.Owner)

	assert.Nil(t, result.Results[1].Error)
	assert.EqualValues(t, 123, result.Results[1].Response.Id)
	assert.EqualValues(t, "testing", result.Results[1].Response.Name)
	assert.EqualValues(t, "federicoleon", result.Results[1].Response.Owner)

}

package github_provier

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"golang/golang-microservices/src/api/clients/rest_client"
	"golang/golang-microservices/src/api/domain/github"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"testing"
)

// All Test Execution
func TestMain(m *testing.M) {
	rest_client.StartMockups()
	os.Exit(m.Run())
}

func TestConstants(t *testing.T) {
	assert.EqualValues(t, "Authorization", headerAuthorization)
	assert.EqualValues(t, "token %s", headerAuthorizationFormat)
	assert.EqualValues(t, "https://api.github.com/user/repos", urlCreateRepo)
}

func TestGetAuthorizationHeader(t *testing.T) {
	headerAuthorizationFormat := getAuthorizationHeader("abc123")
	assert.EqualValues(t, "token abc123", headerAuthorizationFormat)
}

func TestCreateRepoErrorRestClient(t *testing.T) {
	rest_client.FlushMockups()
	rest_client.AddMockup(rest_client.Mock{
		Url:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Err:        errors.New("invalid rest_client response"),
	})

	response, err := CreateRepo("", github.CreateRepoRequest{})

	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.StatusCode)
	assert.EqualValues(t, "invalid rest_client response", err.Message)
}

func TestCreateRepoInvalidResponseBody(t *testing.T) {
	rest_client.FlushMockups()
	invalidCloser, _ := os.Open("-asdf")
	defer invalidCloser.Close()

	rest_client.AddMockup(rest_client.Mock{
		Url:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusInternalServerError,
			Body:       invalidCloser,
		},
	})

	response, err := CreateRepo("", github.CreateRepoRequest{})

	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.StatusCode)
	assert.EqualValues(t, "invalid response body", err.Message)
}

func TestCreateRepoInvalidErrorInterface(t *testing.T) {
	rest_client.FlushMockups()

	rest_client.AddMockup(rest_client.Mock{
		Url:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusUnauthorized,
			Body:       ioutil.NopCloser(strings.NewReader(`{"message" : 1`)),
		},
	})

	response, err := CreateRepo("", github.CreateRepoRequest{})

	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.StatusCode)
	assert.EqualValues(t, "invalid json response body", err.Message)
}

func TestCreateRepoUnauthorized(t *testing.T) {
	rest_client.FlushMockups()
	rest_client.AddMockup(rest_client.Mock{
		Url:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusUnauthorized,
			Body:       ioutil.NopCloser(strings.NewReader(`{"message": "Requires authentication", "documentation_url": "https://developer.githubcom/v3/repos/#create"}`)),
		},
		Err: nil,
	})
}

func TestCreateRepoInvalidSuccessResponse(t *testing.T) {
	rest_client.FlushMockups()

	rest_client.AddMockup(rest_client.Mock{
		Url:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusCreated,
			Body:       ioutil.NopCloser(strings.NewReader(`{"id" : 123`)),
		},
	})

	response, err := CreateRepo("", github.CreateRepoRequest{})

	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.StatusCode)
	assert.EqualValues(t, "json unmarshal fail github create repo response", err.Message)
}

func TestCreateReposeNoError(t *testing.T) {
	rest_client.FlushMockups()

	rest_client.AddMockup(rest_client.Mock{
		Url:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusCreated,
			Body:       ioutil.NopCloser(strings.NewReader(`{"id" : 123, "name": "golang_tutorial", "full_name": "golang_tuto"}`)),
		},
	})

	response, err := CreateRepo("", github.CreateRepoRequest{})

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.EqualValues(t, 123, response.Id)
	assert.EqualValues(t, "golang_tutorial", response.Name)
	assert.EqualValues(t, "golang_tuto", response.FullName)

}

package test_utils

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetMockContext(t *testing.T) {
	request, err := http.NewRequest(http.MethodGet, "http://127.0.0.1:8080", nil)
	assert.Nil(t, err)
	response := httptest.NewRecorder()
	request.Header.Add("X-Mock", "true")
	c := GetMockContext(request, response)

	assert.EqualValues(t, http.MethodGet, c.Request.Method)
	assert.EqualValues(t, "http", c.Request.URL.Scheme)
	assert.Nil(t, c.Request.Body)
	assert.EqualValues(t, "127.0.0.1:8080", c.Request.URL.Host)
	assert.EqualValues(t, "8080", c.Request.URL.Port())
	assert.EqualValues(t, 1, len(c.Request.Header))
	assert.EqualValues(t, "true", c.GetHeader("x-mock"))
	assert.EqualValues(t, "true", c.GetHeader("X-Mock"))

}

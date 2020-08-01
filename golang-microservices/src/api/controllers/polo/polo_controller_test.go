package polo

import (
	"github.com/stretchr/testify/assert"
	"golang/golang-microservices/src/api/utils/test_utils"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestConstants(t *testing.T) {
	assert.EqualValues(t, "polo", polo)
}

func TestMarco(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/marco", nil)
	response := httptest.NewRecorder()
	c := test_utils.GetMockContext(request, response)

	Marco(c)

	assert.EqualValues(t, http.StatusOK, response.Code)
	assert.EqualValues(t, "polo", response.Body.String())

}

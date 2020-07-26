package rest_client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
)

var (
	enableMocks = false
	mocks       = make(map[string]*Mock)
)

type Mock struct {
	Url        string
	HttpMethod string
	Response   *http.Response
	Err        error
}

func getMockId(httpMethod string, url string) string {
	return fmt.Sprintf("%s_%s", httpMethod, url)
}

func StartMockups() {
	enableMocks = true
}

func StopMockups() {
	enableMocks = false
}

func AddMockup(mock Mock) {
	mocks[getMockId(mock.HttpMethod, mock.Url)] = &mock
}

func FlushMockups() {
	mocks = make(map[string]*Mock)
}

func Post(url string, body interface{}, headers http.Header) (*http.Response, error) {
	if enableMocks {
		mock := mocks[getMockId(http.MethodPost, url)]
		if mock == nil {
			return nil, errors.New("no mockup found for give request")

		} else {
			return mock.Response, mock.Err
		}

		// TODO: return local mock without calling any external resource!
	}

	jsonBytes, err := json.Marshal(body)
	if err != nil {
		log.Println("err : ", err)
		return nil, err
	}

	request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(jsonBytes))
	if err != nil {
		log.Println("err : ", err)
		return nil, err
	}

	request.Header = headers

	client := http.Client{}
	return client.Do(request)
}

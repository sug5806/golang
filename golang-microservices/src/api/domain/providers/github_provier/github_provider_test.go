package github_provier

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetAuthorizationHeader(t *testing.T) {
	headerAuthorizationFormat := getAuthorizationHeader("abc123")
	assert.EqualValues(t, "token abc123", headerAuthorizationFormat)
}

func TestDefer(t *testing.T) {
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
	defer fmt.Println("4")
	fmt.Println("function's body")
}

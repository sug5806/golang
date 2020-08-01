package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConst(t *testing.T) {
	assert.EqualValues(t, "SECRET_GITHUB_ACCESS_TOKEN", apiGithubAccessToken)
}

func TestGetGithubAccessToken(t *testing.T) {
	token := GetGithubAccessToken()

	assert.EqualValues(t, "", token)
}

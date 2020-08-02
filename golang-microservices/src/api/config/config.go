package config

import (
	"golang/golang-microservices/src/api/domain/github"
)

const (
	apiGithubAccessToken = "SECRET_GITHUB_ACCESS_TOKEN"
)

var (
	githubAccessToken = github.GITHUB_TOKEN
)

func GetGithubAccessToken() string {
	return githubAccessToken
}

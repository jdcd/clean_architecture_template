package config

import (
	"github.com/jdcd9001/clean-architecture-template/config/factory/repository"
	"os"
)

func GetConfigurations() *AppConfiguration {
	return &AppConfiguration{
		repository:    os.Getenv("REPOSITORY"),
		repositoryURL: os.Getenv("CONNECTION_URL"),
	}
}

type AppConfiguration struct {
	repository    string
	repositoryURL string
}

func (c *AppConfiguration) Repository() repository.AvailableRepo {
	return repository.AvailableRepo(c.repository)
}

func (c *AppConfiguration) ConnectionURL() string {
	return c.repositoryURL
}

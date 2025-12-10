package svc

import (
	"core-service/internal/config"
	"rag-smart-serve/pkg/sdk"
)

type ServiceContext struct {
	Config    config.Config
	ApiClient *sdk.ApiClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		ApiClient: sdk.NewApiClient("http://localhost:11434"),
	}
}

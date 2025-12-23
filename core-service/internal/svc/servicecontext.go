package svc

import (
	"core-service/internal/config"
	"rag-smart-serve/pkg/sdk/llm"
)

type ServiceContext struct {
	Config    config.Config
	ApiClient *llm.ApiClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		ApiClient: llm.NewApiClient("http://localhost:11434"),
	}
}

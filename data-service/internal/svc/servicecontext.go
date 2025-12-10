package svc

import (
	"data-service/internal/config"
	"rag-smart-serve/pkg/constant"
	"rag-smart-serve/pkg/database"
	data_service "rag-smart-serve/pkg/database/dto/data-service"
)

type ServiceContext struct {
	Config         config.Config
	KnowledgeModel *database.MongoManager[data_service.KnowledgeDao]
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:         c,
		KnowledgeModel: database.NewMongoManager[data_service.KnowledgeDao](constant.DataServiceDB, constant.KnowledgeCollection),
	}
}

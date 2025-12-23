package svc

import (
	"context"
	"data-service/internal/config"
	"rag-smart-serve/pkg/constant"
	"rag-smart-serve/pkg/database"
	data_service "rag-smart-serve/pkg/database/dto/data-service"
	"rag-smart-serve/pkg/etcd"
)

type ServiceContext struct {
	Config         config.Config
	KnowledgeModel *database.MongoManager[data_service.KnowledgeDao]
}

func NewServiceContext(c config.Config) *ServiceContext {
	mongoUri, err := etcd.GetValue(context.Background(), constant.MongoKey)
	if err != nil {
		return nil
	}
	return &ServiceContext{
		Config:         c,
		KnowledgeModel: database.NewMongoManager[data_service.KnowledgeDao](mongoUri, constant.DataServiceDB, constant.KnowledgeCollection),
	}
}

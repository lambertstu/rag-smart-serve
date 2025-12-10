package knowledgeservicelogic

import (
	"context"
	customErr "rag-smart-serve/pkg/error_code"
	"time"

	"data-service/data"
	"data-service/internal/svc"
	data_service "rag-smart-serve/pkg/database/dto/data-service"

	"github.com/zeromicro/go-zero/core/logx"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateKnowledgeBaseLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateKnowledgeBaseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateKnowledgeBaseLogic {
	return &CreateKnowledgeBaseLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (receiver *CreateKnowledgeBaseLogic) CreateKnowledgeBase(in *data.CreateKnowledgeBaseReq) (*data.CreateKnowledgeBaseResp, error) {
	doc := &data_service.KnowledgeDao{
		Name:           in.GetName(),
		Description:    in.GetDescription(),
		EmbeddingModel: in.GetEmbeddingModel(),
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}
	insertRes, err := receiver.svcCtx.KnowledgeModel.InsertOne(receiver.ctx, doc)
	if err != nil {
		return nil, customErr.DatabaseError.WithError(err)
	}

	oid, ok := insertRes.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, customErr.DataTypeError.WithDetails("CreateKnowledgeBase id")
	}

	return &data.CreateKnowledgeBaseResp{
		Id: oid.Hex(),
	}, nil
}

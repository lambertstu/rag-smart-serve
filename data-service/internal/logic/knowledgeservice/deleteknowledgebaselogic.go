package knowledgeservicelogic

import (
	"context"
	customErr "rag-smart-serve/pkg/error_code"

	"data-service/data"
	"data-service/internal/svc"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteKnowledgeBaseLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteKnowledgeBaseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteKnowledgeBaseLogic {
	return &DeleteKnowledgeBaseLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteKnowledgeBaseLogic) DeleteKnowledgeBase(in *data.DeleteKnowledgeBaseReq) (*data.DeleteKnowledgeBaseResp, error) {
	oid, err := primitive.ObjectIDFromHex(in.GetId())
	if err != nil {
		return nil, customErr.DataTypeError.WithDetails("invalid id format")
	}

	filter := bson.M{"_id": oid}
	deleteRes, err := l.svcCtx.KnowledgeModel.DeleteOne(l.ctx, filter)
	if err != nil {
		return nil, customErr.DatabaseError.WithError(err)
	}

	if deleteRes.DeletedCount == 0 {
		return nil, customErr.DatabaseError.WithDetails("knowledge base not found")
	}

	return &data.DeleteKnowledgeBaseResp{
		Success: true,
	}, nil
}

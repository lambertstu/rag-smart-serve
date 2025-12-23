package knowledgeservicelogic

import (
	"context"
	"time"

	"data-service/data"
	"data-service/internal/svc"
	customErr "rag-smart-serve/pkg/error_code"

	"github.com/zeromicro/go-zero/core/logx"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UpdateKnowledgeBaseLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateKnowledgeBaseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateKnowledgeBaseLogic {
	return &UpdateKnowledgeBaseLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateKnowledgeBaseLogic) UpdateKnowledgeBase(in *data.UpdateKnowledgeBaseReq) (*data.UpdateKnowledgeBaseResp, error) {
	oid, err := primitive.ObjectIDFromHex(in.Id)
	if err != nil {
		return nil, customErr.ParamVerifyError.WithDetails("invalid id")
	}

	updateFields := bson.M{
		"updatedAt": time.Now(),
	}

	if in.Name != "" {
		updateFields["name"] = in.Name
	}
	if in.Description != "" {
		updateFields["description"] = in.Description
	}
	if in.EmbeddingModel != "" {
		updateFields["embeddingModel"] = in.EmbeddingModel
	}

	update := bson.M{"$set": updateFields}
	res, err := l.svcCtx.KnowledgeModel.UpdateOne(l.ctx, bson.M{"_id": oid}, update)
	if err != nil {
		return nil, customErr.DatabaseError.WithError(err)
	}

	if res.MatchedCount == 0 {
		return nil, customErr.ResourceNotFoundError.WithDetails("knowledge base not found")
	}

	return &data.UpdateKnowledgeBaseResp{
		Success: true,
	}, nil
}

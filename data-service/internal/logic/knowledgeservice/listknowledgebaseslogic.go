package knowledgeservicelogic

import (
	"context"
	customErr "rag-smart-serve/pkg/error_code"

	"data-service/data"
	"data-service/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ListKnowledgeBasesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListKnowledgeBasesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListKnowledgeBasesLogic {
	return &ListKnowledgeBasesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListKnowledgeBasesLogic) ListKnowledgeBases(in *data.ListKnowledgeBasesReq) (*data.ListKnowledgeBasesResp, error) {
	page := in.GetPage()
	if page < 1 {
		page = 1
	}
	pageSize := in.GetPageSize()
	if pageSize < 1 {
		pageSize = 10
	}

	skip := int64((page - 1) * pageSize)
	limit := int64(pageSize)

	total, err := l.svcCtx.KnowledgeModel.Count(l.ctx, bson.M{})
	if err != nil {
		return nil, customErr.DatabaseError.WithError(err)
	}

	opts := options.Find().SetSkip(skip).SetLimit(limit).SetSort(bson.M{"created_at": -1}) // 按创建时间倒序
	list, err := l.svcCtx.KnowledgeModel.FindMany(l.ctx, bson.M{}, opts)
	if err != nil {
		return nil, customErr.DatabaseError.WithError(err)
	}

	var pbList []*data.KnowledgeBase
	for _, k := range list {
		pbList = append(pbList, &data.KnowledgeBase{
			Name:           k.Name,
			Description:    k.Description,
			EmbeddingModel: k.EmbeddingModel,
			CreatedAt:      k.CreatedAt.UnixMilli(),
			UpdatedAt:      k.UpdatedAt.UnixMilli(),
		})
	}

	return &data.ListKnowledgeBasesResp{
		List:  pbList,
		Total: total,
	}, nil
}

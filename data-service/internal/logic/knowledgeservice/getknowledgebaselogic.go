package knowledgeservicelogic

import (
	"context"
	customErr "rag-smart-serve/pkg/error_code"

	"data-service/data"
	"data-service/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GetKnowledgeBaseLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetKnowledgeBaseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetKnowledgeBaseLogic {
	return &GetKnowledgeBaseLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetKnowledgeBaseLogic) GetKnowledgeBase(in *data.GetKnowledgeBaseReq) (*data.GetKnowledgeBaseResp, error) {
	oid, err := primitive.ObjectIDFromHex(in.GetId())
	if err != nil {
		return nil, customErr.DataTypeError.WithDetails("invalid id format")
	}

	knowledge, err := l.svcCtx.KnowledgeModel.FindByID(l.ctx, oid)
	if err != nil {
		if err.Error() == "mongo: no documents in result" { // 这里需要根据实际 err 判断，或者检查 FindByID 内部是否返回了特定的 ErrNotFound
			return nil, customErr.DatabaseError.WithDetails("knowledge base not found")
		}
		return nil, customErr.DatabaseError.WithError(err)
	}

	// 此时 knowledge 是 *data_service.KnowledgeDao 类型，需要转换为 proto message
	// 注意：KnowledgeDao 中没有 ID 字段（ID 是 MongoDB 的 _id），这里假设返回的结构体中包含了数据
	// 如果 KnowledgeDao 定义中确实没有 ID 字段，这里只能填入业务字段
	// 另外，created_at 和 updated_at 在 KnowledgeDao 中是 time.Time，在 proto 中是 int64 (假设是 unix timestamp)

	return &data.GetKnowledgeBaseResp{
		KnowledgeBase: &data.KnowledgeBase{
			Name:           knowledge.Name,
			Description:    knowledge.Description,
			EmbeddingModel: knowledge.EmbeddingModel,
			CreatedAt:      knowledge.CreatedAt.UnixMilli(),
			UpdatedAt:      knowledge.UpdatedAt.UnixMilli(),
		},
	}, nil
}

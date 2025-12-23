package knowledgeservicelogic

import (
	"context"

	"data-service/data"
	"data-service/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateDocumentStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateDocumentStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateDocumentStatusLogic {
	return &UpdateDocumentStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateDocumentStatusLogic) UpdateDocumentStatus(in *data.UpdateDocumentStatusReq) (*data.UpdateDocumentStatusResp, error) {
	// todo: add your logic here and delete this line

	return &data.UpdateDocumentStatusResp{}, nil
}

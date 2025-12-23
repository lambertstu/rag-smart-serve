package knowledgeservicelogic

import (
	"context"

	"data-service/data"
	"data-service/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListDocumentsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListDocumentsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListDocumentsLogic {
	return &ListDocumentsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListDocumentsLogic) ListDocuments(in *data.ListDocumentsReq) (*data.ListDocumentsResp, error) {
	// todo: add your logic here and delete this line

	return &data.ListDocumentsResp{}, nil
}

package knowledgeservicelogic

import (
	"context"

	"data-service/data"
	"data-service/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteDocumentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteDocumentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteDocumentLogic {
	return &DeleteDocumentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteDocumentLogic) DeleteDocument(in *data.DeleteDocumentReq) (*data.DeleteDocumentResp, error) {
	// todo: add your logic here and delete this line

	return &data.DeleteDocumentResp{}, nil
}

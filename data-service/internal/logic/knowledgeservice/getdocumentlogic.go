package knowledgeservicelogic

import (
	"context"

	"data-service/data"
	"data-service/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetDocumentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetDocumentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDocumentLogic {
	return &GetDocumentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetDocumentLogic) GetDocument(in *data.GetDocumentReq) (*data.GetDocumentResp, error) {
	// todo: add your logic here and delete this line

	return &data.GetDocumentResp{}, nil
}

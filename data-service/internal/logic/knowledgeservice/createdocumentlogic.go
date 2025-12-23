package knowledgeservicelogic

import (
	"context"

	"data-service/data"
	"data-service/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateDocumentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateDocumentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateDocumentLogic {
	return &CreateDocumentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateDocumentLogic) CreateDocument(in *data.CreateDocumentReq) (*data.CreateDocumentResp, error) {
	// todo: add your logic here and delete this line

	return &data.CreateDocumentResp{}, nil
}

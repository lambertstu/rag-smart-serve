package promptservicelogic

import (
	"context"

	"core-service/core"
	"core-service/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTemplateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetTemplateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTemplateLogic {
	return &GetTemplateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取单个模板
func (l *GetTemplateLogic) GetTemplate(in *core.GetTemplateReq) (*core.GetTemplateResp, error) {
	// todo: add your logic here and delete this line

	return &core.GetTemplateResp{}, nil
}

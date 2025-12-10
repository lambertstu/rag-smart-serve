package promptservicelogic

import (
	"context"

	"core-service/core"
	"core-service/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateTemplateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateTemplateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateTemplateLogic {
	return &UpdateTemplateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新模板
func (l *UpdateTemplateLogic) UpdateTemplate(in *core.UpdateTemplateReq) (*core.UpdateTemplateResp, error) {
	// todo: add your logic here and delete this line

	return &core.UpdateTemplateResp{}, nil
}

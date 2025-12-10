package promptservicelogic

import (
	"context"

	"core-service/core"
	"core-service/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteTemplateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteTemplateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteTemplateLogic {
	return &DeleteTemplateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除模板
func (l *DeleteTemplateLogic) DeleteTemplate(in *core.DeleteTemplateReq) (*core.DeleteTemplateResp, error) {
	// todo: add your logic here and delete this line

	return &core.DeleteTemplateResp{}, nil
}

package promptservicelogic

import (
	"context"

	"core-service/core"
	"core-service/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type RenderTemplateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRenderTemplateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RenderTemplateLogic {
	return &RenderTemplateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 渲染模板（填充变量）
func (l *RenderTemplateLogic) RenderTemplate(in *core.RenderTemplateReq) (*core.RenderTemplateResp, error) {
	// todo: add your logic here and delete this line

	return &core.RenderTemplateResp{}, nil
}

package promptservicelogic

import (
	"context"

	"core-service/core"
	"core-service/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListTemplatesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListTemplatesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListTemplatesLogic {
	return &ListTemplatesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取模板列表
func (l *ListTemplatesLogic) ListTemplates(in *core.ListTemplatesReq) (*core.ListTemplatesResp, error) {
	// todo: add your logic here and delete this line

	return &core.ListTemplatesResp{}, nil
}

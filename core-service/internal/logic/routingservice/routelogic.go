package routingservicelogic

import (
	"context"

	"core-service/core"
	"core-service/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type RouteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRouteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RouteLogic {
	return &RouteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 路由决策（AI继续 or 转人工）
func (l *RouteLogic) Route(in *core.RouteReq) (*core.RouteResp, error) {
	// todo: add your logic here and delete this line

	return &core.RouteResp{}, nil
}

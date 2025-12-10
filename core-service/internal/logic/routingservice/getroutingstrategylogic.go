package routingservicelogic

import (
	"context"

	"core-service/core"
	"core-service/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRoutingStrategyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetRoutingStrategyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRoutingStrategyLogic {
	return &GetRoutingStrategyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取路由策略配置
func (l *GetRoutingStrategyLogic) GetRoutingStrategy(in *core.GetRoutingStrategyReq) (*core.GetRoutingStrategyResp, error) {
	// todo: add your logic here and delete this line

	return &core.GetRoutingStrategyResp{}, nil
}

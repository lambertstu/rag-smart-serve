package emotionservicelogic

import (
	"context"

	"core-service/core"
	"core-service/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DetectRiskLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDetectRiskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DetectRiskLogic {
	return &DetectRiskLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 风险预警检测（检测高危用户）
func (l *DetectRiskLogic) DetectRisk(in *core.DetectRiskReq) (*core.DetectRiskResp, error) {
	// todo: add your logic here and delete this line

	return &core.DetectRiskResp{}, nil
}

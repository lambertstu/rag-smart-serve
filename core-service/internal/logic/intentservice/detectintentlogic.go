package intentservicelogic

import (
	"context"

	"core-service/core"
	"core-service/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DetectIntentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDetectIntentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DetectIntentLogic {
	return &DetectIntentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 识别用户意图
func (l *DetectIntentLogic) DetectIntent(in *core.DetectIntentReq) (*core.DetectIntentResp, error) {
	// todo: add your logic here and delete this line

	return &core.DetectIntentResp{}, nil
}

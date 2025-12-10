package intentservicelogic

import (
	"context"

	"core-service/core"
	"core-service/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type ClassifyMessageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewClassifyMessageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ClassifyMessageLogic {
	return &ClassifyMessageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 消息分类（快速分类：咨询/投诉/闲聊等）
func (l *ClassifyMessageLogic) ClassifyMessage(in *core.ClassifyMessageReq) (*core.ClassifyMessageResp, error) {
	// todo: add your logic here and delete this line

	return &core.ClassifyMessageResp{}, nil
}

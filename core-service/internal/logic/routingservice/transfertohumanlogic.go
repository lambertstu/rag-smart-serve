package routingservicelogic

import (
	"context"

	"core-service/core"
	"core-service/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type TransferToHumanLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTransferToHumanLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TransferToHumanLogic {
	return &TransferToHumanLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 转人工客服
func (l *TransferToHumanLogic) TransferToHuman(in *core.TransferToHumanReq) (*core.TransferToHumanResp, error) {
	// todo: add your logic here and delete this line

	return &core.TransferToHumanResp{}, nil
}

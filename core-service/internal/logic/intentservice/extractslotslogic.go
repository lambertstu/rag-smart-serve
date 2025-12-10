package intentservicelogic

import (
	"context"

	"core-service/core"
	"core-service/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type ExtractSlotsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewExtractSlotsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ExtractSlotsLogic {
	return &ExtractSlotsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 提取槽位信息（NER）
func (l *ExtractSlotsLogic) ExtractSlots(in *core.ExtractSlotsReq) (*core.ExtractSlotsResp, error) {
	// todo: add your logic here and delete this line

	return &core.ExtractSlotsResp{}, nil
}

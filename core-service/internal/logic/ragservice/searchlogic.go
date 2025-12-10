package ragservicelogic

import (
	"context"

	"core-service/core"
	"core-service/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchLogic {
	return &SearchLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchLogic) Search(in *core.SearchReq) (*core.SearchResp, error) {
	// todo: add your logic here and delete this line

	return &core.SearchResp{}, nil
}

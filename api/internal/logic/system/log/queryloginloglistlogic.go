package log

import (
	"context"

	"github.com/feihua/zero-admin-template/api/internal/svc"
	"github.com/feihua/zero-admin-template/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryLoginLogListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryLoginLogListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryLoginLogListLogic {
	return &QueryLoginLogListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryLoginLogListLogic) QueryLoginLogList(req *types.QueryLoginLogListReq) (resp *types.QueryLoginLogListResp, err error) {
	// todo: add your logic here and delete this line

	return
}

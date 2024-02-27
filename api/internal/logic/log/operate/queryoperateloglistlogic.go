package operate

import (
	"context"

	"github.com/feihua/zero-admin-template/api/internal/svc"
	"github.com/feihua/zero-admin-template/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryOperateLogListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryOperateLogListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryOperateLogListLogic {
	return &QueryOperateLogListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryOperateLogListLogic) QueryOperateLogList(req *types.QueryOperateLogListReq) (resp *types.QueryOperateLogListResp, err error) {
	// todo: add your logic here and delete this line

	return
}

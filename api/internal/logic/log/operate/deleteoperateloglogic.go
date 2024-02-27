package operate

import (
	"context"

	"github.com/feihua/zero-admin-template/api/internal/svc"
	"github.com/feihua/zero-admin-template/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteOperateLogLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteOperateLogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteOperateLogLogic {
	return &DeleteOperateLogLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteOperateLogLogic) DeleteOperateLog(req *types.DeleteOperateLogReq) (resp *types.DeleteOperateLogResp, err error) {
	// todo: add your logic here and delete this line

	return
}

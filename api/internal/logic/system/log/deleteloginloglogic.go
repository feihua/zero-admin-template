package log

import (
	"context"

	"github.com/feihua/zero-admin-template/api/internal/svc"
	"github.com/feihua/zero-admin-template/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteLoginLogLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteLoginLogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteLoginLogLogic {
	return &DeleteLoginLogLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteLoginLogLogic) DeleteLoginLog(req *types.DeleteLoginLogReq) (resp *types.DeleteLoginLogResp, err error) {
	// todo: add your logic here and delete this line

	return
}

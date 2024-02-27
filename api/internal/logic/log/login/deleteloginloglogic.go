package login

import (
	"context"
	"github.com/feihua/zero-admin-template/api/internal/svc"
	"github.com/feihua/zero-admin-template/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// DeleteLoginLogLogic 删除登录日志
/*
Author: LiuFeiHua
Date: 2024/2/27 12:13
*/
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

// DeleteLoginLog 删除登录日志
func (l *DeleteLoginLogLogic) DeleteLoginLog(req *types.DeleteLoginLogReq) (resp *types.DeleteLoginLogResp, err error) {
	resp = &types.DeleteLoginLogResp{}
	for _, id := range req.Ids {
		_ = l.svcCtx.LoginLogModel.Delete(l.ctx, id)
	}

	resp.Code = 0
	resp.Msg = "删除操作日志成功"
	return
}

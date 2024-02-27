package operate

import (
	"context"
	"github.com/feihua/zero-admin-template/api/internal/svc"
	"github.com/feihua/zero-admin-template/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// DeleteOperateLogLogic 删除操作日志
/*
Author: LiuFeiHua
Date: 2024/2/27 12:11
*/
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

// DeleteOperateLog 删除操作日志
func (l *DeleteOperateLogLogic) DeleteOperateLog(req *types.DeleteOperateLogReq) (resp *types.DeleteOperateLogResp, err error) {
	resp = &types.DeleteOperateLogResp{}
	for _, id := range req.Ids {
		_ = l.svcCtx.OperateLogModel.Delete(l.ctx, id)
	}

	resp.Code = 0
	resp.Msg = "删除操作日志成功"

	return
}

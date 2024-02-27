package user

import (
	"context"
	"github.com/feihua/zero-admin-template/api/internal/common/errorx"

	"github.com/feihua/zero-admin-template/api/internal/svc"
	"github.com/feihua/zero-admin-template/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// DeleteUserLogic
/*
Author: LiuFeiHua
Date: 2024/2/23 下午3:59
*/
type DeleteUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteUserLogic {
	return &DeleteUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// DeleteUser 删除用户
func (l *DeleteUserLogic) DeleteUser(req *types.DeleteUserReq) (resp *types.DeleteUserResp, err error) {
	resp = &types.DeleteUserResp{}
	for _, id := range req.Ids {
		err = l.svcCtx.UserModel.Delete(l.ctx, id)

		if err != nil {
			logx.WithContext(l.ctx).Errorf("根据userId: %d,删除用户异常:%s", id, err.Error())
			return nil, errorx.NewDefaultError("删除用户异常")
		}
	}

	resp.Code = 0
	resp.Msg = "删除用户成功"
	return
}

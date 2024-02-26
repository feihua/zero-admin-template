package role

import (
	"context"
	"github.com/feihua/zero-admin-template/api/internal/common/errorx"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin-template/api/internal/svc"
	"github.com/feihua/zero-admin-template/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// DeleteRoleLogic
/*
Author: LiuFeiHua
Date: 2024/2/23 下午3:31
*/
type DeleteRoleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteRoleLogic {
	return &DeleteRoleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// DeleteRole 删除角色
func (l *DeleteRoleLogic) DeleteRole(req *types.DeleteRoleReq) (resp *types.DeleteRoleResp, err error) {
	resp = &types.DeleteRoleResp{}
	for _, id := range req.Ids {
		err := l.svcCtx.RoleModel.Delete(l.ctx, id)

		if err != nil {
			logc.Errorf(l.ctx, "根据roleId: %d,删除角色异常:%s", id, err.Error())
			return nil, errorx.NewDefaultError("删除角色异常")
		}
	}

	resp.Code = 0
	resp.Msg = "删除角色成功"
	return
}

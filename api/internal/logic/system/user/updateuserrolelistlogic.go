package user

import (
	"context"
	"github.com/feihua/zero-admin-template/api/internal/common/errorx"
	"github.com/feihua/zero-admin-template/api/internal/model"

	"github.com/feihua/zero-admin-template/api/internal/svc"
	"github.com/feihua/zero-admin-template/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateUserRoleListLogic
/*
Author: LiuFeiHua
Date: 2024/2/23 下午4:24
*/
type UpdateUserRoleListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateUserRoleListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserRoleListLogic {
	return &UpdateUserRoleListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UpdateUserRoleList 更新用户角色关联
func (l *UpdateUserRoleListLogic) UpdateUserRoleList(req *types.UpdateUserRoleListReq) (resp *types.UpdateUserRoleListResp, err error) {
	resp = &types.UpdateUserRoleListResp{}
	if req.UserId == 1 {
		return nil, errorx.NewDefaultError("不能修改超级管理员的角色")
	}

	_ = l.svcCtx.UserRoleModel.DeleteByUserId(l.ctx, req.UserId)

	for _, id := range req.RoleIds {
		_, _ = l.svcCtx.UserRoleModel.Insert(l.ctx, &model.SysUserRole{
			RoleId: id,
			UserId: req.UserId,
		})
	}

	resp.Code = 0
	resp.Msg = "更新用户角色关联成功"
	return
}

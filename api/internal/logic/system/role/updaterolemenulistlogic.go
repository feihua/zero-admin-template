package role

import (
	"context"
	"github.com/feihua/zero-admin-template/api/internal/common/errorx"
	"github.com/feihua/zero-admin-template/api/internal/model"

	"github.com/feihua/zero-admin-template/api/internal/svc"
	"github.com/feihua/zero-admin-template/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateRoleMenuListLogic
/*
Author: LiuFeiHua
Date: 2024/2/23 下午3:46
*/
type UpdateRoleMenuListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateRoleMenuListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateRoleMenuListLogic {
	return &UpdateRoleMenuListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UpdateRoleMenuList 更新角色菜单
func (l *UpdateRoleMenuListLogic) UpdateRoleMenuList(req *types.UpdateRoleMenuListReq) (resp *types.UpdateRoleMenuListResp, err error) {
	resp = &types.UpdateRoleMenuListResp{}
	if req.RoleId == 1 {
		return nil, errorx.NewDefaultError("不能修改超级管理员的权限")
	}

	_ = l.svcCtx.RoleMenuModel.DeleteByRoleId(l.ctx, req.RoleId)

	ids := req.MenuIds
	for _, id := range ids {
		_, _ = l.svcCtx.RoleMenuModel.Insert(l.ctx, &model.SysRoleMenu{
			MenuId: id,
			RoleId: req.RoleId,
		})
	}

	resp.Code = 0
	resp.Msg = "更新角色菜单成功"
	return
}

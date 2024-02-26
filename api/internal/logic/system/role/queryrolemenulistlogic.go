package role

import (
	"context"
	"strconv"

	"github.com/feihua/zero-admin-template/api/internal/svc"
	"github.com/feihua/zero-admin-template/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryRoleMenuListLogic 查询角色菜单列表
/*
Author: LiuFeiHua
Date: 2024/2/23 下午3:46
*/
type QueryRoleMenuListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryRoleMenuListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryRoleMenuListLogic {
	return &QueryRoleMenuListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// QueryRoleMenuList 查询角色菜单列表
func (l *QueryRoleMenuListLogic) QueryRoleMenuList(req *types.QueryRoleMenuListReq) (resp *types.QueryRoleMenuListResp, err error) {
	resp = &types.QueryRoleMenuListResp{}
	findAll, _ := l.svcCtx.MenuModel.FindAll(l.ctx, "")

	var roleMenuList []types.RoleMenuList
	var menuIds []int64

	for _, menu := range *findAll {
		roleMenuList = append(roleMenuList, types.RoleMenuList{
			Id:       menu.Id,
			ParentID: menu.ParentId,
			Title:    menu.MenuName,
			Key:      strconv.FormatInt(menu.Id, 10),
		})

		menuIds = append(menuIds, menu.Id)
	}

	if req.RoleId != 1 {
		menuIds, _ = l.svcCtx.RoleMenuModel.FindAllByRoleId(l.ctx, req.RoleId)
	}

	resp.Code = 0
	resp.Msg = "查询角色菜单成功"
	resp.Data = types.QueryRoleMenuListData{
		RoleMenus:    menuIds,
		MenuRoleList: roleMenuList,
	}
	return
}

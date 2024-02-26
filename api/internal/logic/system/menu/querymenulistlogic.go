package menu

import (
	"context"
	"github.com/feihua/zero-admin-template/api/internal/common/errorx"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin-template/api/internal/svc"
	"github.com/feihua/zero-admin-template/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryMenuListLogic
/*
Author: LiuFeiHua
Date: 2024/2/23 下午3:11
*/
type QueryMenuListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryMenuListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryMenuListLogic {
	return &QueryMenuListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// QueryMenuList 查询菜单列表
func (l *QueryMenuListLogic) QueryMenuList(req *types.QueryMenuListReq) (resp *types.QueryMenuListResp, err error) {
	resp = &types.QueryMenuListResp{}
	menuList, err := l.svcCtx.MenuModel.FindAll(l.ctx, req.MenuName)

	if err != nil {
		logc.Errorf(l.ctx, "参数: %+v,查询菜单列表异常:%s", req, err.Error())
		return nil, errorx.NewDefaultError("查询菜单失败")
	}

	var list []types.QueryMenuListData
	for _, menu := range *menuList {
		list = append(list, types.QueryMenuListData{
			Id:         menu.Id,
			MenuName:   menu.MenuName,
			MenuType:   menu.MenuType,
			StatusId:   menu.StatusId,
			Sort:       menu.Sort,
			ParentId:   menu.ParentId,
			MenuUrl:    menu.MenuUrl,
			ApiUrl:     menu.ApiUrl,
			MenuIcon:   menu.MenuIcon.String,
			Remark:     menu.Remark.String,
			CreateTime: menu.CreateTime.Format("2006-01-02 15:04:05"),
			UpdateTime: menu.UpdateTime.Time.Format("2006-01-02 15:04:05"),
		})
	}

	resp.Code = 0
	resp.Msg = "查询菜单成功"
	resp.Data = list
	return
}

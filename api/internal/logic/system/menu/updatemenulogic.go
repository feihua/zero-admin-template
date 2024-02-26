package menu

import (
	"context"
	"database/sql"
	"github.com/feihua/zero-admin-template/api/internal/common/errorx"
	"github.com/feihua/zero-admin-template/api/internal/model"
	"github.com/feihua/zero-admin-template/api/internal/svc"
	"github.com/feihua/zero-admin-template/api/internal/types"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateMenuLogic
/*
Author: LiuFeiHua
Date: 2024/2/23 下午3:17
*/
type UpdateMenuLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMenuLogic {
	return &UpdateMenuLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UpdateMenu 更新菜单
func (l *UpdateMenuLogic) UpdateMenu(req *types.UpdateMenuReq) (resp *types.UpdateMenuResp, err error) {
	resp = &types.UpdateMenuResp{}
	err = l.svcCtx.MenuModel.Update(l.ctx, &model.SysMenu{
		Id:       req.Id,
		StatusId: req.StatusId,
		Sort:     req.Sort,
		ParentId: req.ParentId,
		MenuName: req.MenuName,
		MenuUrl:  req.MenuUrl,
		ApiUrl:   req.ApiUrl,
		MenuIcon: sql.NullString{
			String: req.MenuIcon,
			Valid:  true,
		},
		Remark: sql.NullString{
			String: req.Remark,
			Valid:  true,
		},
		MenuType: req.MenuType,
	})

	if err != nil {
		logc.Errorf(l.ctx, "参数: %+v,更新菜单异常:%s", req, err.Error())
		return nil, errorx.NewDefaultError("更新菜单失败")
	}

	resp.Code = 0
	resp.Msg = "添加菜单成功"
	return
}

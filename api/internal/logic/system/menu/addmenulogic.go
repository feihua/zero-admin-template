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

// AddMenuLogic
/*
Author: LiuFeiHua
Date: 2024/2/23 下午3:14
*/
type AddMenuLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddMenuLogic {
	return &AddMenuLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// AddMenu 添加菜单
func (l *AddMenuLogic) AddMenu(req *types.AddMenuReq) (resp *types.AddMenuResp, err error) {
	resp = &types.AddMenuResp{}
	_, err = l.svcCtx.MenuModel.Insert(l.ctx, &model.SysMenu{
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
		logc.Errorf(l.ctx, "参数: %+v,添加菜单异常:%s", req, err.Error())
		return nil, errorx.NewDefaultError("添加菜单失败")
	}

	resp.Code = 0
	resp.Msg = "添加菜单成功"
	return
}

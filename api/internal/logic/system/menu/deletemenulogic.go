package menu

import (
	"context"
	"github.com/feihua/zero-admin-template/api/internal/common/errorx"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin-template/api/internal/svc"
	"github.com/feihua/zero-admin-template/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// DeleteMenuLogic
/*
Author: LiuFeiHua
Date: 2024/2/23 下午3:00
*/
type DeleteMenuLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteMenuLogic {
	return &DeleteMenuLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// DeleteMenu 删除菜单
func (l *DeleteMenuLogic) DeleteMenu(req *types.DeleteMenuReq) (resp *types.DeleteMenuResp, err error) {
	resp = &types.DeleteMenuResp{}
	err = l.svcCtx.MenuModel.Delete(l.ctx, req.Id)

	if err != nil {
		logc.Errorf(l.ctx, "根据menuId: %d,删除菜单异常:%s", req.Id, err.Error())
		return nil, errorx.NewDefaultError("删除菜单异常")
	}

	resp.Code = 0
	resp.Msg = "删除菜单成功"
	return
}

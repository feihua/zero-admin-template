package role

import (
	"context"
	"github.com/feihua/zero-admin-template/api/internal/common/errorx"
	"github.com/feihua/zero-admin-template/api/internal/model"
	"github.com/feihua/zero-admin-template/api/internal/svc"
	"github.com/feihua/zero-admin-template/api/internal/types"
	"github.com/zeromicro/go-zero/core/logx"
)

// AddRoleLogic
/*
Author: LiuFeiHua
Date: 2024/2/23 下午3:27
*/
type AddRoleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddRoleLogic {
	return &AddRoleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// AddRole 添加角色
func (l *AddRoleLogic) AddRole(req *types.AddRoleReq) (resp *types.AddRoleResp, err error) {
	resp = &types.AddRoleResp{}
	_, err = l.svcCtx.RoleModel.Insert(l.ctx, &model.SysRole{
		RoleName: req.RoleName,
		StatusId: req.StatusId,
		Sort:     req.Sort,
		Remark:   req.Remark,
	})

	if err != nil {
		return nil, errorx.NewDefaultError("添加角色失败")
	}

	resp.Code = 0
	resp.Msg = "添加角色成功"
	return
}

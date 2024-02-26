package role

import (
	"context"
	"github.com/feihua/zero-admin-template/api/internal/common/errorx"
	"github.com/feihua/zero-admin-template/api/internal/model"
	"github.com/feihua/zero-admin-template/api/internal/svc"
	"github.com/feihua/zero-admin-template/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateRoleLogic
/*
Author: LiuFeiHua
Date: 2024/2/23 下午3:39
*/
type UpdateRoleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateRoleLogic {
	return &UpdateRoleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UpdateRole 更新角色
func (l *UpdateRoleLogic) UpdateRole(req *types.UpdateRoleReq) (resp *types.UpdateRoleResp, err error) {
	resp = &types.UpdateRoleResp{}
	err = l.svcCtx.RoleModel.Update(l.ctx, &model.SysRole{
		Id:       req.Id,
		StatusId: req.StatusId,
		Sort:     req.Sort,
		RoleName: req.RoleName,
		Remark:   req.Remark,
	})

	if err != nil {
		logx.WithContext(l.ctx).Errorf(err.Error())
		return nil, errorx.NewDefaultError("更新角色失败")
	}

	resp.Code = 0
	resp.Msg = "更新角色成功"
	return
}

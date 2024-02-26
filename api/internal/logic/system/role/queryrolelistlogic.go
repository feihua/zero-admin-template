package role

import (
	"context"
	"github.com/feihua/zero-admin-template/api/internal/common/errorx"

	"github.com/feihua/zero-admin-template/api/internal/svc"
	"github.com/feihua/zero-admin-template/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryRoleListLogic
/*
Author: LiuFeiHua
Date: 2024/2/23 下午3:32
*/
type QueryRoleListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryRoleListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryRoleListLogic {
	return &QueryRoleListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// QueryRoleList 查询角色列表
func (l *QueryRoleListLogic) QueryRoleList(req *types.QueryRoleListReq) (resp *types.QueryRoleListResp, err error) {
	resp = &types.QueryRoleListResp{}
	roleList, count, err := l.svcCtx.RoleModel.FindAll(l.ctx, req.Current, req.PageSize, req.RoleName, req.StatusId)

	if err != nil {
		return nil, errorx.NewDefaultError("查询角色列表异常")
	}
	var list []types.RoleListData

	for _, role := range *roleList {
		list = append(list, types.RoleListData{
			Id:         role.Id,
			StatusId:   role.StatusId,
			Sort:       role.Sort,
			RoleName:   role.RoleName,
			Remark:     role.Remark,
			CreateTime: role.CreateTime.Format("2006-01-02 15:04:05"),
			UpdateTime: role.UpdateTime.Time.Format("2006-01-02 15:04:05"),
		})
	}

	resp.Code = 0
	resp.Msg = "查询角色列表成功"
	resp.Data = list
	resp.Success = true
	resp.Total = count
	return
}

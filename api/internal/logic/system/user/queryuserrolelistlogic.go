package user

import (
	"context"

	"github.com/feihua/zero-admin-template/api/internal/svc"
	"github.com/feihua/zero-admin-template/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryUserRoleListLogic 查询用户角色列表
/*
Author: LiuFeiHua
Date: 2024/2/23 下午4:18
*/
type QueryUserRoleListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryUserRoleListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryUserRoleListLogic {
	return &QueryUserRoleListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// QueryUserRoleList 查询用户角色列表
func (l *QueryUserRoleListLogic) QueryUserRoleList(req *types.QueryUserRoleListReq) (resp *types.QueryUserRoleListResp, err error) {
	resp = &types.QueryUserRoleListResp{}
	//所有角色
	allRoleList, _ := l.svcCtx.RoleModel.QueryRoleList(l.ctx)

	var roleList []types.UserRoleList
	var allUserRoleIds []int64

	for _, role := range *allRoleList {
		roleList = append(roleList, types.UserRoleList{
			Id:         role.Id,
			StatusId:   role.StatusId,
			Sort:       role.Sort,
			RoleName:   role.RoleName,
			Remark:     role.Remark,
			CreateTime: role.CreateTime.Format("2006-01-02 15:04:05"),
			UpdateTime: role.UpdateTime.Time.Format("2006-01-02 15:04:05"),
		})
		allUserRoleIds = append(allUserRoleIds, role.Id)
	}

	if req.UserId != 1 {
		allUserRoleIds, _ = l.svcCtx.UserRoleModel.FindAllRoleIdsByUserId(l.ctx, req.UserId)
	}

	resp.Code = 0
	resp.Msg = "查询用户角色列表成功"
	resp.Data = types.QueryUserRoleListData{
		AllRoles:       roleList,
		AllUserRoleIds: allUserRoleIds,
	}
	return
}

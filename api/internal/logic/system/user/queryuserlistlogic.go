package user

import (
	"context"
	"github.com/feihua/zero-admin-template/api/internal/common/errorx"
	"github.com/feihua/zero-admin-template/api/internal/svc"
	"github.com/feihua/zero-admin-template/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryUserListLogic
/*
Author: LiuFeiHua
Date: 2024/2/23 下午4:06
*/
type QueryUserListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryUserListLogic {
	return &QueryUserListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// QueryUserList 查询用户列表
func (l *QueryUserListLogic) QueryUserList(req *types.QueryUserListReq) (resp *types.QueryUserListResp, err error) {
	resp = &types.QueryUserListResp{}
	userList, count, err := l.svcCtx.UserModel.FindAll(l.ctx, req.Current, req.PageSize, req.Mobile, req.StatusId)

	if err != nil {
		return nil, errorx.NewDefaultError("查询角色列表异常")
	}

	var list []types.UserListData

	for _, user := range *userList {
		list = append(list, types.UserListData{
			Id:         user.Id,
			Mobile:     user.Mobile,
			UserName:   user.UserName,
			StatusId:   user.StatusId,
			Sort:       user.Sort,
			Remark:     user.Remark.String,
			CreateTime: user.CreateTime.Format("2006-01-02 15:04:05"),
			UpdateTime: user.UpdateTime.Time.Format("2006-01-02 15:04:05"),
		})
	}

	resp.Code = 0
	resp.Msg = "查询用户列表成功"
	resp.Data = list
	resp.Success = true
	resp.Total = count
	return
}

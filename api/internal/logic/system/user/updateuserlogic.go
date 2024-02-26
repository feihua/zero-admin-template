package user

import (
	"context"
	"database/sql"
	"github.com/feihua/zero-admin-template/api/internal/common/errorx"
	"github.com/feihua/zero-admin-template/api/internal/model"
	"github.com/feihua/zero-admin-template/api/internal/svc"
	"github.com/feihua/zero-admin-template/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateUserLogic
/*
Author: LiuFeiHua
Date: 2024/2/23 下午4:21
*/
type UpdateUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserLogic {
	return &UpdateUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UpdateUser 更新用户
func (l *UpdateUserLogic) UpdateUser(req *types.UpdateUserReq) (resp *types.UpdateUserResp, err error) {
	resp = &types.UpdateUserResp{}
	err = l.svcCtx.UserModel.Update(l.ctx, &model.SysUser{
		Id:       req.Id,
		StatusId: req.StatusId,
		Sort:     req.Sort,
		Mobile:   req.Mobile,
		UserName: req.UserName,
		Remark:   sql.NullString{String: req.Remark, Valid: true},
	})

	if err != nil {
		return nil, errorx.NewDefaultError("更新用户异常")
	}

	resp.Code = 0
	resp.Msg = "更新用户成功"
	return
}

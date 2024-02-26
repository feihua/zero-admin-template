package user

import (
	"context"
	"database/sql"
	"github.com/feihua/zero-admin-template/api/internal/common/errorx"
	"github.com/feihua/zero-admin-template/api/internal/model"
	"time"

	"github.com/feihua/zero-admin-template/api/internal/svc"
	"github.com/feihua/zero-admin-template/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// AddUserLogic
/*
Author: LiuFeiHua
Date: 2024/2/23 下午3:54
*/
type AddUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddUserLogic {
	return &AddUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// AddUser 添加用户
func (l *AddUserLogic) AddUser(req *types.AddUserReq) (resp *types.AddUserResp, err error) {
	resp = &types.AddUserResp{}
	mobile, _ := l.svcCtx.UserModel.FindOneByMobile(l.ctx, req.Mobile)

	if mobile != nil {
		return nil, errorx.NewDefaultError("手机号码已注册")
	}

	_, err = l.svcCtx.UserModel.Insert(l.ctx, &model.SysUser{
		Mobile:   req.Mobile,
		UserName: req.UserName,
		//用户默认密码
		Password: "123456",
		StatusId: req.StatusId,
		Sort:     req.Sort,
		Remark: sql.NullString{
			String: req.Remark,
			Valid:  true,
		},
		CreateTime: time.Time{},
		UpdateTime: sql.NullTime{},
	})

	if err != nil {
		return nil, errorx.NewDefaultError("添加用户异常")
	}

	resp.Code = 0
	resp.Msg = "添加用户成功"
	return
}

package login

import (
	"context"
	"github.com/feihua/zero-admin-template/api/internal/common/errorx"

	"github.com/feihua/zero-admin-template/api/internal/svc"
	"github.com/feihua/zero-admin-template/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryLoginLogListLogic 查询登录日志列表
/*
Author: LiuFeiHua
Date: 2024/2/27 12:14
*/
type QueryLoginLogListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryLoginLogListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryLoginLogListLogic {
	return &QueryLoginLogListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// QueryLoginLogList 查询登录日志列表
func (l *QueryLoginLogListLogic) QueryLoginLogList(req *types.QueryLoginLogListReq) (resp *types.QueryLoginLogListResp, err error) {
	resp = &types.QueryLoginLogListResp{}
	userList, count, err := l.svcCtx.LoginLogModel.FindAll(l.ctx, req.Current, req.PageSize, req.UserName, req.Ip)

	if err != nil {
		return nil, errorx.NewDefaultError("查询登录日志列表异常")
	}

	var list []types.LoginLogListData

	for _, log := range *userList {
		list = append(list, types.LoginLogListData{
			Id:        log.Id,
			UserName:  log.UserName,
			Status:    log.Status,
			Ip:        log.Ip,
			LoginTime: log.LoginTime.Format("2006-01-02 15:04:05"),
		})
	}

	resp.Code = 0
	resp.Msg = "查询登录日志列表成功"
	resp.Data = list
	resp.Success = true
	resp.Total = count
	return
}

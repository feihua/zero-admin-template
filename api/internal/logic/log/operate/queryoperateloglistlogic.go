package operate

import (
	"context"
	"github.com/feihua/zero-admin-template/api/internal/common/errorx"

	"github.com/feihua/zero-admin-template/api/internal/svc"
	"github.com/feihua/zero-admin-template/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryOperateLogListLogic 查询操作日志列表
/*
Author: LiuFeiHua
Date: 2024/2/27 13:37
*/
type QueryOperateLogListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryOperateLogListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryOperateLogListLogic {
	return &QueryOperateLogListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// QueryOperateLogList 查询操作日志列表
func (l *QueryOperateLogListLogic) QueryOperateLogList(req *types.QueryOperateLogListReq) (resp *types.QueryOperateLogListResp, err error) {
	resp = &types.QueryOperateLogListResp{}
	userList, count, err := l.svcCtx.OperateLogModel.FindAll(l.ctx, req.Current, req.PageSize, req.UserName, req.Method, req.Ip)

	if err != nil {
		return nil, errorx.NewDefaultError("查询操作日志列表异常")
	}

	var list []types.OperateLogListData

	for _, log := range *userList {
		list = append(list, types.OperateLogListData{
			Id:             log.Id,
			UserName:       log.UserName,
			Operation:      log.Operation,
			Method:         log.Method,
			RequestParams:  log.RequestParams,
			Time:           log.Time,
			Ip:             log.Ip.String,
			ResponseParams: log.ResponseParams.String,
			OperationTime:  log.OperationTime.Format("2006-01-02 15:04:05"),
		})
	}

	resp.Code = 0
	resp.Msg = "查询操作日志列表成功"
	resp.Data = list
	resp.Success = true
	resp.Total = count
	return
}

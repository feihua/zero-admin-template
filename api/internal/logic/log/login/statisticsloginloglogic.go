package login

import (
	"context"

	"github.com/feihua/zero-admin-template/api/internal/svc"
	"github.com/feihua/zero-admin-template/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type StatisticsLoginLogLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewStatisticsLoginLogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StatisticsLoginLogLogic {
	return &StatisticsLoginLogLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *StatisticsLoginLogLogic) StatisticsLoginLog() (resp *types.StatisticsLoginLogResp, err error) {
	//查询当天登录人数（根据IP）
	dayLoginCount, _ := l.svcCtx.LoginLogModel.StatisticsLoginLog(l.ctx, 1)
	//统计当前周登录人数（根据IP）
	weekLoginCount, _ := l.svcCtx.LoginLogModel.StatisticsLoginLog(l.ctx, 2)
	//统计当前月登录人数（根据IP）
	monthLoginCount, _ := l.svcCtx.LoginLogModel.StatisticsLoginLog(l.ctx, 3)

	resp = &types.StatisticsLoginLogResp{}
	resp.Code = 0
	resp.Message = "0"
	resp.Data = types.StatisticsLoginLogData{
		DayLoginCount:   dayLoginCount,
		WeekLoginCount:  weekLoginCount,
		MonthLoginCount: monthLoginCount,
	}
	return
}

package user

import (
	"context"
	"github.com/feihua/zero-admin-template/api/internal/common/errorx"
	"github.com/feihua/zero-admin-template/api/internal/model"
	"github.com/golang-jwt/jwt/v4"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"strings"
	"time"

	"github.com/feihua/zero-admin-template/api/internal/svc"
	"github.com/feihua/zero-admin-template/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// UserLogin
/*
Author: LiuFeiHua
Date: 2024/2/23 下午4:25
*/
type UserLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLoginLogic {
	return &UserLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UserLogin 用户登陆
func (l *UserLoginLogic) UserLogin(req *types.UserLoginReq, ip string) (resp *types.UserLoginResp, err error) {
	resp = &types.UserLoginResp{}
	if len(strings.TrimSpace(req.Mobile)) == 0 || len(strings.TrimSpace(req.Password)) == 0 {
		return nil, errorx.NewDefaultError("参数错误")
	}
	userInfo, err := l.svcCtx.UserModel.FindOneByMobile(l.ctx, req.Mobile)

	switch err {
	case nil:
	case sqlc.ErrNotFound:
		logx.WithContext(l.ctx).Errorf("用户不存在,参数:%s,异常:%s", req.Mobile, err.Error())
		return nil, errorx.NewDefaultError("用户不存在")
	default:
		logx.WithContext(l.ctx).Errorf("用户登录失败,参数:%s,异常:%s", req.Mobile, err.Error())
		return nil, errorx.NewDefaultError(err.Error())
	}

	if userInfo.Password != req.Password {
		logx.WithContext(l.ctx).Errorf("用户密码不正确,参数:%s", req.Password)
		return nil, errorx.NewDefaultError("用户密码不正确")
	}

	now := time.Now().Unix()
	accessExpire := l.svcCtx.Config.Auth.AccessExpire
	accessSecret := l.svcCtx.Config.Auth.AccessSecret

	jwtToken, _ := l.getJwtToken(accessSecret, userInfo.UserName, now, accessExpire, userInfo.Id)

	//保存登录日志
	_, _ = l.svcCtx.LoginLogModel.Insert(l.ctx, &model.SysLoginLog{
		UserName:  userInfo.UserName,
		Status:    "1",
		Ip:        ip,
		LoginTime: time.Now(),
	})

	return &types.UserLoginResp{
		Code: 0,
		Msg:  "登陆成功",
		Data: types.UserLoginData{
			UserId: userInfo.Id,
			Mobile: userInfo.Mobile,
			Token:  jwtToken,
		},
	}, nil
}

func (l *UserLoginLogic) getJwtToken(secretKey, userName string, iat, seconds, userId int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["userId"] = userId
	claims["userName"] = userName
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}

package user

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/feihua/zero-admin-template/api/internal/common/errorx"
	"github.com/feihua/zero-admin-template/api/internal/model"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"strconv"
	"strings"

	"github.com/feihua/zero-admin-template/api/internal/svc"
	"github.com/feihua/zero-admin-template/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryUserMenuListLogic 查询用户菜单列表
type QueryUserMenuListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryUserMenuListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryUserMenuListLogic {
	return &QueryUserMenuListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// QueryUserMenuList 查询用户菜单列表
func (l *QueryUserMenuListLogic) QueryUserMenuList() (resp *types.QueryUserMenuListResp, err error) {
	resp = &types.QueryUserMenuListResp{}
	userId, _ := l.ctx.Value("userId").(json.Number).Int64()

	userInfo, err := l.svcCtx.UserModel.FindOne(l.ctx, userId)
	switch err {
	case nil:
	case sqlc.ErrNotFound:
		logx.WithContext(l.ctx).Infof("用户不存在userId: %s", userId)
		return nil, errorx.NewDefaultError(fmt.Sprintf("用户不存在userId: %s", strconv.FormatInt(userId, 10)))
	default:
		return nil, err
	}

	var menuList *[]model.SysMenu

	roleUser, _ := l.svcCtx.UserRoleModel.FindAllByUserId(l.ctx, userId)
	if userId == 1 || roleUser != nil { //默认管理员,或者后来分配的
		menuList, _ = l.svcCtx.MenuModel.FindAll(l.ctx, "")
	} else {
		menuList, err = l.svcCtx.UserModel.FindAllByUserId(l.ctx, userId)
		if err != nil {
			logx.WithContext(l.ctx).Errorf("用户: %s,还没有权限,请求联系管理员", strconv.FormatInt(userId, 10))
			return nil, errorx.NewDefaultError("还没有权限,请求联系管理员")

		}
	}

	var list []types.UserMenuList
	var btnList []string

	var apiUrls []string

	for _, menu := range *menuList {
		if menu.MenuType != 3 {
			list = append(list, types.UserMenuList{
				Id:       menu.Id,
				ParentId: menu.ParentId,
				Name:     menu.MenuName,
				Path:     menu.MenuUrl,
				ApiUrl:   menu.ApiUrl,
				MenuType: menu.MenuType,
				Icon:     menu.MenuIcon.String,
			})
		}

		if menu.MenuType == 3 {
			btnList = append(btnList, menu.ApiUrl)
		}

		if len(strings.TrimSpace(menu.ApiUrl)) != 0 {
			apiUrls = append(apiUrls, menu.ApiUrl)
		}
	}

	//把能访问的url存在在redis，在middleware中检验
	userIdStr := fmt.Sprintf("token:%s", strconv.FormatInt(userId, 10))
	err = l.svcCtx.Redis.Set(userIdStr, strings.Join(apiUrls, ","))

	if err != nil {
		logx.Errorf("设置用户：%s,权限到redis异常: %+v", userInfo.UserName, err)
	}

	resp.Code = 0
	resp.Msg = "查询用户菜单列表成功"
	resp.Data = types.QueryUserMenuListData{
		SysMenu: list,
		BtnMenu: btnList,
		Avatar:  "https://gw.alipayobjects.com/zos/antfincdn/XAosXuNZyF/BiazfanxmamNRoxxVxka.png",
		Name:    userInfo.UserName,
	}
	return

}

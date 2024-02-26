package user

import (
	"net/http"

	"github.com/feihua/zero-admin-template/api/internal/logic/system/user"
	"github.com/feihua/zero-admin-template/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func QueryUserMenuListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := user.NewQueryUserMenuListLogic(r.Context(), svcCtx)
		resp, err := l.QueryUserMenuList()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

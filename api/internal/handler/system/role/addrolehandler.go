package role

import (
	"net/http"

	"github.com/feihua/zero-admin-template/api/internal/logic/system/role"
	"github.com/feihua/zero-admin-template/api/internal/svc"
	"github.com/feihua/zero-admin-template/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func AddRoleHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddRoleReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := role.NewAddRoleLogic(r.Context(), svcCtx)
		resp, err := l.AddRole(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

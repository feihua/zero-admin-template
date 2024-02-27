package operate

import (
	"net/http"

	"github.com/feihua/zero-admin-template/api/internal/logic/log/operate"
	"github.com/feihua/zero-admin-template/api/internal/svc"
	"github.com/feihua/zero-admin-template/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func QueryOperateLogListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.QueryOperateLogListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := operate.NewQueryOperateLogListLogic(r.Context(), svcCtx)
		resp, err := l.QueryOperateLogList(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

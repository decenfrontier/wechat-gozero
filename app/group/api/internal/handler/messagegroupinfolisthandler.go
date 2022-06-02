package handler

import (
	"net/http"

	"ws_chat/common/xresp"

	"github.com/zeromicro/go-zero/rest/httpx"
	"ws_chat/app/group/api/internal/logic"
	"ws_chat/app/group/api/internal/svc"
	"ws_chat/app/group/api/internal/types"
)

func MessageGroupInfoListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.MessageGroupInfoListRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		if err := xresp.Validate.StructCtx(r.Context(), req); err != nil {
			xresp.Response(r, w, nil, err)
			return
		}

		l := logic.NewMessageGroupInfoListLogic(r.Context(), svcCtx)
		resp, err := l.MessageGroupInfoList(&req)
		xresp.Response(r, w, resp, err)
	}
}

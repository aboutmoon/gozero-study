package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"study/internal/logic"
	"study/internal/svc"
	"study/internal/types"
)

func StudyHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewStudyLogic(r.Context(), svcCtx)
		resp, err := l.Study(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

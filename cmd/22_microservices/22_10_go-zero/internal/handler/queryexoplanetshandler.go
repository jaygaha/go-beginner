package handler

import (
	"net/http"

	"github.com/jaygaha/go-beginner/cmd/22_microservices/22_10_go-zero/internal/logic"
	"github.com/jaygaha/go-beginner/cmd/22_microservices/22_10_go-zero/internal/svc"
	"github.com/jaygaha/go-beginner/cmd/22_microservices/22_10_go-zero/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func queryExoplanetsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ExoplanetQueryRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewQueryExoplanetsLogic(r.Context(), svcCtx)
		resp, err := l.QueryExoplanets(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

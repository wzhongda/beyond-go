package handler

import (
	"net/http"

	"beyond-go/application/article/api/internal/logic"
	"beyond-go/application/article/api/internal/svc"
	"beyond-go/application/article/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ArticleListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ArticleListRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewArticleListLogic(r.Context(), svcCtx)
		resp, err := l.ArticleList(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

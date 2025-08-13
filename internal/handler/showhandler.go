package handler

import (
	"net/http"

	"shortener/internal/logic"
	"shortener/internal/svc"
	"shortener/internal/types"

	"github.com/go-playground/validator/v10"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ShowHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ShowRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		// 参数合法校验
		if err := validator.New().StructCtx(r.Context(), &req); err != nil {
			logx.Errorw("validation check failed", logx.LogField{
				Key:   "err",
				Value: err.Error(),
			})
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewShowLogic(r.Context(), svcCtx)
		resp, err := l.Show(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			// httpx.OkJsonCtx(r.Context(), w, resp)
			http.Redirect(w, r, resp.LongUrl, http.StatusFound)
		}
	}
}

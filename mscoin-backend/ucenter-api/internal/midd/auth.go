package midd

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
	common "mscoin-common"
	"mscoin-common/tools"
	"net/http"
)

func Auth(secret string) func(next http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			result := common.NewResult()
			result.Fail(4000, "no login")
			token := r.Header.Get("x-auth-token")
			if token == "" {
				logx.Infof("[Auth] token is empty")
				httpx.WriteJson(w, 200, result)
				return
			}
			logx.Infof("[Auth] token received: %s...", token[:min(20, len(token))])
			userId, err := tools.ParseToken(token, secret)
			if err != nil {
				logx.Errorf("[Auth] ParseToken error: %v, token: %s...", err, token[:min(20, len(token))])
				httpx.WriteJson(w, 200, result)
				return
			}
			logx.Infof("[Auth] userId parsed: %d", userId)
			ctx := r.Context()
			ctx = context.WithValue(ctx, "userId", userId)
			r = r.WithContext(ctx)
			next(w, r)
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

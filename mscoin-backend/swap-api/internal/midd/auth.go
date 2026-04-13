package midd

import (
	"context"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	common "mscoin-common"
	"mscoin-common/tools"
)

const AuthUserIDKey = "userId"

func Auth(secret string) func(next http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			result := common.NewResult()
			result.Fail(4000, "no login")

			token := r.Header.Get("x-auth-token")
			if token == "" {
				httpx.WriteJson(w, http.StatusOK, result)
				return
			}

			userID, err := tools.ParseToken(token, secret)
			if err != nil {
				httpx.WriteJson(w, http.StatusOK, result)
				return
			}

			ctx := context.WithValue(r.Context(), AuthUserIDKey, userID)
			next(w, r.WithContext(ctx))
		}
	}
}

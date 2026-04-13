package handler

import (
	"net/http"

	"swap-api/internal/midd"
)

func authenticatedMemberID(r *http.Request, fallback int64) int64 {
	value := r.Context().Value(midd.AuthUserIDKey)
	userID, ok := value.(int64)
	if !ok || userID <= 0 {
		return fallback
	}
	return userID
}

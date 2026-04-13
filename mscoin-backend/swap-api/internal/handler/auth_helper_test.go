package handler

import (
	"context"
	"net/http"
	"testing"

	"swap-api/internal/midd"
)

func TestAuthenticatedMemberIDUsesContextUser(t *testing.T) {
	req, err := http.NewRequest(http.MethodPost, "/uc/contract/current", nil)
	if err != nil {
		t.Fatalf("new request failed: %v", err)
	}

	req = req.WithContext(context.WithValue(req.Context(), midd.AuthUserIDKey, int64(2002)))
	if got := authenticatedMemberID(req, 1001); got != 2002 {
		t.Fatalf("expected authenticated user id, got %d", got)
	}
}

func TestAuthenticatedMemberIDFallsBackToBodyMemberID(t *testing.T) {
	req, err := http.NewRequest(http.MethodPost, "/uc/contract/current", nil)
	if err != nil {
		t.Fatalf("new request failed: %v", err)
	}

	if got := authenticatedMemberID(req, 1001); got != 1001 {
		t.Fatalf("expected fallback member id, got %d", got)
	}
}

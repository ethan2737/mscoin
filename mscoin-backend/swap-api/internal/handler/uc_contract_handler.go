package handler

import (
	"net/http"
	"strconv"
	"strings"

	"swap-api/internal/logic"
	"swap-api/internal/svc"
	"swap-api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
	common "mscoin-common"
	"mscoin-common/tools"
)

type UcContractHandler struct {
	svcCtx *svc.ServiceContext
}

func NewUcContractHandler(svcCtx *svc.ServiceContext) *UcContractHandler {
	return &UcContractHandler{
		svcCtx: svcCtx,
	}
}

func (h *UcContractHandler) Current(w http.ResponseWriter, r *http.Request) {
	var req types.CurrentOrderReq
	if err := httpx.ParseJsonBody(r, &req); err != nil {
		httpx.ErrorCtx(r.Context(), w, err)
		return
	}
	req.MemberId = authenticatedMemberID(r, req.MemberId)
	req.Ip = tools.GetRemoteClientIp(r)
	l := logic.NewUcContractLogic(r.Context(), h.svcCtx)
	resp, err := l.Current(&req)
	result := common.NewResult().Deal(resp, err)
	httpx.OkJsonCtx(r.Context(), w, result)
}

func (h *UcContractHandler) History(w http.ResponseWriter, r *http.Request) {
	var req types.HistoryOrderReq
	if err := httpx.ParseJsonBody(r, &req); err != nil {
		httpx.ErrorCtx(r.Context(), w, err)
		return
	}
	req.MemberId = authenticatedMemberID(r, req.MemberId)
	req.Ip = tools.GetRemoteClientIp(r)
	l := logic.NewUcContractLogic(r.Context(), h.svcCtx)
	resp, err := l.History(&req)
	result := common.NewResult().Deal(resp, err)
	httpx.OkJsonCtx(r.Context(), w, result)
}

func (h *UcContractHandler) Thumb(w http.ResponseWriter, r *http.Request) {
	req := &types.GetContractListReq{}
	l := logic.NewUcContractLogic(r.Context(), h.svcCtx)
	resp, err := l.Thumb(req)
	result := common.NewResult().Deal(resp, err)
	httpx.OkJsonCtx(r.Context(), w, result)
}

func (h *UcContractHandler) Cancel(w http.ResponseWriter, r *http.Request) {
	var req types.CancelOrderReq
	if err := httpx.ParseJsonBody(r, &req); err != nil {
		httpx.ErrorCtx(r.Context(), w, err)
		return
	}
	req.MemberId = authenticatedMemberID(r, req.MemberId)

	parts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
	if req.OrderId == 0 && len(parts) > 0 {
		idStr := parts[len(parts)-1]
		idInt, err := strconv.ParseInt(idStr, 10, 64)
		if err == nil {
			req.OrderId = idInt
		}
	}

	req.Ip = tools.GetRemoteClientIp(r)
	l := logic.NewUcContractLogic(r.Context(), h.svcCtx)
	resp, err := l.Cancel(&req)
	result := common.NewResult().Deal(resp, err)
	httpx.OkJsonCtx(r.Context(), w, result)
}

func ucCurrentOrderHandler(svc *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		NewUcContractHandler(svc).Current(w, r)
	}
}

func ucHistoryOrderHandler(svc *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		NewUcContractHandler(svc).History(w, r)
	}
}

func ucContractCoinThumbHandler(svc *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		NewUcContractHandler(svc).Thumb(w, r)
	}
}

func ucCancelOrderHandler(svc *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		NewUcContractHandler(svc).Cancel(w, r)
	}
}

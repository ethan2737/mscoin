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

type OrderEntrustHandler struct {
	svcCtx *svc.ServiceContext
}

func NewOrderEntrustHandler(svcCtx *svc.ServiceContext) *OrderEntrustHandler {
	return &OrderEntrustHandler{
		svcCtx: svcCtx,
	}
}

func (h *OrderEntrustHandler) Add(w http.ResponseWriter, r *http.Request) {
	var req types.AddOrderReq
	if err := httpx.ParseJsonBody(r, &req); err != nil {
		httpx.ErrorCtx(r.Context(), w, err)
		return
	}
	req.MemberId = authenticatedMemberID(r, req.MemberId)
	req.Ip = tools.GetRemoteClientIp(r)
	l := logic.NewOrderEntrustLogic(r.Context(), h.svcCtx)
	resp, err := l.Add(&req)
	result := common.NewResult().Deal(resp, err)
	httpx.OkJsonCtx(r.Context(), w, result)
}

func (h *OrderEntrustHandler) Cancel(w http.ResponseWriter, r *http.Request) {
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
	l := logic.NewOrderEntrustLogic(r.Context(), h.svcCtx)
	resp, err := l.Cancel(&req)
	result := common.NewResult().Deal(resp, err)
	httpx.OkJsonCtx(r.Context(), w, result)
}

func (h *OrderEntrustHandler) Current(w http.ResponseWriter, r *http.Request) {
	var req types.CurrentOrderReq
	if err := httpx.ParseJsonBody(r, &req); err != nil {
		httpx.ErrorCtx(r.Context(), w, err)
		return
	}
	req.MemberId = authenticatedMemberID(r, req.MemberId)
	req.Ip = tools.GetRemoteClientIp(r)
	l := logic.NewOrderEntrustLogic(r.Context(), h.svcCtx)
	resp, err := l.Current(&req)
	result := common.NewResult().Deal(resp, err)
	httpx.OkJsonCtx(r.Context(), w, result)
}

func (h *OrderEntrustHandler) History(w http.ResponseWriter, r *http.Request) {
	var req types.HistoryOrderReq
	if err := httpx.ParseJsonBody(r, &req); err != nil {
		httpx.ErrorCtx(r.Context(), w, err)
		return
	}
	req.MemberId = authenticatedMemberID(r, req.MemberId)
	req.Ip = tools.GetRemoteClientIp(r)
	l := logic.NewOrderEntrustLogic(r.Context(), h.svcCtx)
	resp, err := l.History(&req)
	result := common.NewResult().Deal(resp, err)
	httpx.OkJsonCtx(r.Context(), w, result)
}

func (h *OrderEntrustHandler) Position(w http.ResponseWriter, r *http.Request) {
	var req types.PositionReq
	if err := httpx.ParseJsonBody(r, &req); err != nil {
		httpx.ErrorCtx(r.Context(), w, err)
		return
	}
	req.MemberId = authenticatedMemberID(r, req.MemberId)
	req.Ip = tools.GetRemoteClientIp(r)
	l := logic.NewOrderEntrustLogic(r.Context(), h.svcCtx)
	resp, err := l.Position(&req)
	result := common.NewResult().Deal(resp, err)
	httpx.OkJsonCtx(r.Context(), w, result)
}

func (h *OrderEntrustHandler) QuickClose(w http.ResponseWriter, r *http.Request) {
	var req types.QuickCloseReq
	if err := httpx.ParseJsonBody(r, &req); err != nil {
		httpx.ErrorCtx(r.Context(), w, err)
		return
	}
	req.MemberId = authenticatedMemberID(r, req.MemberId)

	parts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
	if req.ContractCoinId == 0 && len(parts) > 0 {
		idStr := parts[len(parts)-1]
		idInt, err := strconv.ParseInt(idStr, 10, 32)
		if err == nil {
			req.ContractCoinId = int32(idInt)
		}
	}

	req.Ip = tools.GetRemoteClientIp(r)
	l := logic.NewOrderEntrustLogic(r.Context(), h.svcCtx)
	resp, err := l.QuickClose(&req)
	result := common.NewResult().Deal(resp, err)
	httpx.OkJsonCtx(r.Context(), w, result)
}

func (h *OrderEntrustHandler) Leverage(w http.ResponseWriter, r *http.Request) {
	var req types.LeverageReq
	if err := httpx.ParseJsonBody(r, &req); err != nil {
		httpx.ErrorCtx(r.Context(), w, err)
		return
	}
	req.MemberId = authenticatedMemberID(r, req.MemberId)
	req.Ip = tools.GetRemoteClientIp(r)
	l := logic.NewOrderEntrustLogic(r.Context(), h.svcCtx)
	resp, err := l.Leverage(&req)
	result := common.NewResult().Deal(resp, err)
	httpx.OkJsonCtx(r.Context(), w, result)
}

func addOrderEntrustHandler(svc *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		NewOrderEntrustHandler(svc).Add(w, r)
	}
}

func cancelOrderEntrustHandler(svc *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		NewOrderEntrustHandler(svc).Cancel(w, r)
	}
}

func currentOrderEntrustHandler(svc *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		NewOrderEntrustHandler(svc).Current(w, r)
	}
}

func historyOrderEntrustHandler(svc *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		NewOrderEntrustHandler(svc).History(w, r)
	}
}

func positionHandler(svc *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		NewOrderEntrustHandler(svc).Position(w, r)
	}
}

func quickCloseHandler(svc *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		NewOrderEntrustHandler(svc).QuickClose(w, r)
	}
}

func leverageHandler(svc *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		NewOrderEntrustHandler(svc).Leverage(w, r)
	}
}

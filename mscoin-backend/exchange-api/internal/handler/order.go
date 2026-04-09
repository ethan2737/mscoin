package handler

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"exchange-api/internal/logic"
	"exchange-api/internal/svc"
	"exchange-api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
	common "mscoin-common"
	"mscoin-common/tools"
)

type OrderHandler struct {
	svcCtx *svc.ServiceContext
}

func NewOrderHandler(svcCtx *svc.ServiceContext) *OrderHandler {
	return &OrderHandler{
		svcCtx: svcCtx,
	}
}

func parseExchangeReq(r *http.Request, req *types.ExchangeReq) error {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	if len(bytes.TrimSpace(body)) > 0 {
		if err := json.Unmarshal(body, req); err == nil {
			r.Body = io.NopCloser(bytes.NewReader(body))
			return nil
		}
	}
	r.Body = io.NopCloser(bytes.NewReader(body))
	return httpx.ParseForm(r, req)
}

func (h *OrderHandler) History(w http.ResponseWriter, r *http.Request) {
	var req types.ExchangeReq
	if err := parseExchangeReq(r, &req); err != nil {
		httpx.ErrorCtx(r.Context(), w, err)
		return
	}
	ip := tools.GetRemoteClientIp(r)
	req.Ip = ip
	l := logic.NewOrderLogic(r.Context(), h.svcCtx)
	resp, err := l.History(&req)
	result := common.NewResult().Deal(resp, err)
	httpx.OkJsonCtx(r.Context(), w, result)
}

func (h *OrderHandler) Current(w http.ResponseWriter, r *http.Request) {
	var req types.ExchangeReq
	if err := parseExchangeReq(r, &req); err != nil {
		httpx.ErrorCtx(r.Context(), w, err)
		return
	}
	ip := tools.GetRemoteClientIp(r)
	req.Ip = ip
	l := logic.NewOrderLogic(r.Context(), h.svcCtx)
	resp, err := l.Current(&req)
	result := common.NewResult().Deal(resp, err)
	httpx.OkJsonCtx(r.Context(), w, result)
}

func (h *OrderHandler) Add(w http.ResponseWriter, r *http.Request) {
	var req types.ExchangeReq
	if err := parseExchangeReq(r, &req); err != nil {
		httpx.ErrorCtx(r.Context(), w, err)
		return
	}
	ip := tools.GetRemoteClientIp(r)
	req.Ip = ip
	l := logic.NewOrderLogic(r.Context(), h.svcCtx)
	resp, err := l.AddOrder(&req)
	result := common.NewResult().Deal(resp, err)
	httpx.OkJsonCtx(r.Context(), w, result)
}

func (h *OrderHandler) Cancel(w http.ResponseWriter, r *http.Request) {
	var req types.ExchangeReq
	if err := parseExchangeReq(r, &req); err != nil {
		httpx.ErrorCtx(r.Context(), w, err)
		return
	}
	if req.OrderId == "" {
		parts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
		if len(parts) > 0 {
			req.OrderId = parts[len(parts)-1]
		}
	}
	ip := tools.GetRemoteClientIp(r)
	req.Ip = ip
	l := logic.NewOrderLogic(r.Context(), h.svcCtx)
	resp, err := l.Cancel(&req)
	result := common.NewResult().Deal(resp, err)
	httpx.OkJsonCtx(r.Context(), w, result)
}

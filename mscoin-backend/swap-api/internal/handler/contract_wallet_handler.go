package handler

import (
	"encoding/json"
	"net/http"

	"swap-api/internal/logic"
	"swap-api/internal/svc"
	"swap-api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
	common "mscoin-common"
	"mscoin-common/tools"
)

type ContractWalletHandler struct {
	svcCtx *svc.ServiceContext
}

func NewContractWalletHandler(svcCtx *svc.ServiceContext) *ContractWalletHandler {
	return &ContractWalletHandler{
		svcCtx: svcCtx,
	}
}

func (h *ContractWalletHandler) Info(w http.ResponseWriter, r *http.Request) {
	var req types.ContractWalletReq
	if err := httpx.ParseJsonBody(r, &req); err != nil {
		httpx.ErrorCtx(r.Context(), w, err)
		return
	}
	req.MemberId = authenticatedMemberID(r, req.MemberId)
	req.Ip = tools.GetRemoteClientIp(r)
	l := logic.NewContractWalletLogic(r.Context(), h.svcCtx)
	resp, err := l.Info(&req)
	result := common.NewResult().Deal(resp, err)
	httpx.OkJsonCtx(r.Context(), w, result)
}

func (h *ContractWalletHandler) Transfer(w http.ResponseWriter, r *http.Request) {
	var req types.ContractTransferReq
	if err := httpx.ParseJsonBody(r, &req); err != nil {
		httpx.ErrorCtx(r.Context(), w, err)
		return
	}
	req.MemberId = authenticatedMemberID(r, req.MemberId)
	req.Ip = tools.GetRemoteClientIp(r)
	l := logic.NewContractWalletLogic(r.Context(), h.svcCtx)
	resp, err := l.Transfer(&req)
	result := common.NewResult().Deal(resp, err)
	httpx.OkJsonCtx(r.Context(), w, result)
}

func (h *ContractWalletHandler) Transaction(w http.ResponseWriter, r *http.Request) {
	var req types.ContractTransactionReq
	decoder := json.NewDecoder(r.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&req); err != nil {
		httpx.ErrorCtx(r.Context(), w, err)
		return
	}
	req.MemberId = authenticatedMemberID(r, req.MemberId)
	req.Ip = tools.GetRemoteClientIp(r)
	l := logic.NewContractWalletLogic(r.Context(), h.svcCtx)
	resp, err := l.Transaction(&req)
	result := common.NewResult().Deal(resp, err)
	httpx.OkJsonCtx(r.Context(), w, result)
}

func contractWalletHandler(svc *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		NewContractWalletHandler(svc).Info(w, r)
	}
}

func contractWalletTransferHandler(svc *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		NewContractWalletHandler(svc).Transfer(w, r)
	}
}

func contractTransactionHandler(svc *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		NewContractWalletHandler(svc).Transaction(w, r)
	}
}

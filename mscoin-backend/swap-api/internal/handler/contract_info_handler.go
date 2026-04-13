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
)

type ContractInfoHandler struct {
	svcCtx *svc.ServiceContext
}

func NewContractInfoHandler(svcCtx *svc.ServiceContext) *ContractInfoHandler {
	return &ContractInfoHandler{
		svcCtx: svcCtx,
	}
}

func (h *ContractInfoHandler) Info(w http.ResponseWriter, r *http.Request) {
	// 从 URL 路径获取 id: /swap/contract-coin/info/:id
	parts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
	var id int32
	if len(parts) > 0 {
		idStr := parts[len(parts)-1]
		idInt, err := strconv.ParseInt(idStr, 10, 32)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		id = int32(idInt)
	}

	req := &types.GetContractInfoReq{Id: id}
	l := logic.NewContractInfoLogic(r.Context(), h.svcCtx)
	resp, err := l.Info(req)
	result := common.NewResult().Deal(resp, err)
	httpx.OkJsonCtx(r.Context(), w, result)
}

func (h *ContractInfoHandler) List(w http.ResponseWriter, r *http.Request) {
	req := &types.GetContractListReq{}
	l := logic.NewContractInfoLogic(r.Context(), h.svcCtx)
	resp, err := l.List(req)
	result := common.NewResult().Deal(resp, err)
	httpx.OkJsonCtx(r.Context(), w, result)
}

func (h *ContractInfoHandler) SymbolThumb(w http.ResponseWriter, r *http.Request) {
	req := &types.GetContractListReq{}
	l := logic.NewContractInfoLogic(r.Context(), h.svcCtx)
	resp, err := l.List(req)
	result := common.NewResult().Deal(resp, err)
	httpx.OkJsonCtx(r.Context(), w, result)
}

// 兼容旧代码的包装函数
func contractInfoHandler(svc *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		NewContractInfoHandler(svc).Info(w, r)
	}
}

func symbolThumbHandler(svc *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		NewContractInfoHandler(svc).SymbolThumb(w, r)
	}
}

func exchangePlateMiniHandler(svc *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO: 实现获取简版盘口 handler
		w.WriteHeader(http.StatusNotImplemented)
	}
}

func exchangePlateFullHandler(svc *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO: 实现获取深度图 handler
		w.WriteHeader(http.StatusNotImplemented)
	}
}

func latestTradeHandler(svc *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO: 实现获取最新成交 handler
		w.WriteHeader(http.StatusNotImplemented)
	}
}

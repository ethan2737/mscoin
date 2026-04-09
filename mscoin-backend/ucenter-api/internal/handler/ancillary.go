package handler

import (
	"net/http"

	"ucenter-api/internal/logic"
	"ucenter-api/internal/svc"
	"ucenter-api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
	common "mscoin-common"
)

type AncillaryHandler struct {
	svcCtx *svc.ServiceContext
}

func NewAncillaryHandler(svcCtx *svc.ServiceContext) *AncillaryHandler {
	return &AncillaryHandler{svcCtx: svcCtx}
}

func parseContentReq(r *http.Request, req *types.ContentReq) error {
	if r.ContentLength == 0 {
		return nil
	}
	if err := httpx.ParseJsonBody(r, req); err == nil {
		return nil
	}
	return httpx.ParseForm(r, req)
}

func (h *AncillaryHandler) HelpOverview(w http.ResponseWriter, r *http.Request) {
	var req types.ContentReq
	if err := parseContentReq(r, &req); err != nil {
		httpx.ErrorCtx(r.Context(), w, err)
		return
	}
	resp, err := logic.NewAncillaryLogic(r.Context(), h.svcCtx).HelpOverview(&req)
	httpx.OkJsonCtx(r.Context(), w, common.NewResult().Deal(resp, err))
}

func (h *AncillaryHandler) HelpPage(w http.ResponseWriter, r *http.Request) {
	var req types.ContentReq
	if err := parseContentReq(r, &req); err != nil {
		httpx.ErrorCtx(r.Context(), w, err)
		return
	}
	resp, err := logic.NewAncillaryLogic(r.Context(), h.svcCtx).HelpPage(&req)
	httpx.OkJsonCtx(r.Context(), w, common.NewResult().Deal(resp, err))
}

func (h *AncillaryHandler) HelpTop(w http.ResponseWriter, r *http.Request) {
	var req types.ContentReq
	if err := parseContentReq(r, &req); err != nil {
		httpx.ErrorCtx(r.Context(), w, err)
		return
	}
	resp, err := logic.NewAncillaryLogic(r.Context(), h.svcCtx).HelpTop(&req)
	httpx.OkJsonCtx(r.Context(), w, common.NewResult().Deal(resp, err))
}

func (h *AncillaryHandler) HelpDetail(w http.ResponseWriter, r *http.Request) {
	var req types.ContentReq
	if err := parseContentReq(r, &req); err != nil {
		httpx.ErrorCtx(r.Context(), w, err)
		return
	}
	resp, err := logic.NewAncillaryLogic(r.Context(), h.svcCtx).HelpDetail(&req)
	httpx.OkJsonCtx(r.Context(), w, common.NewResult().Deal(resp, err))
}

func (h *AncillaryHandler) AnnouncementPage(w http.ResponseWriter, r *http.Request) {
	var req types.ContentReq
	if err := parseContentReq(r, &req); err != nil {
		httpx.ErrorCtx(r.Context(), w, err)
		return
	}
	resp, err := logic.NewAncillaryLogic(r.Context(), h.svcCtx).AnnouncementPage(&req)
	httpx.OkJsonCtx(r.Context(), w, common.NewResult().Deal(resp, err))
}

func (h *AncillaryHandler) AnnouncementDetail(w http.ResponseWriter, r *http.Request) {
	var req types.ContentReq
	if err := parseContentReq(r, &req); err != nil {
		httpx.ErrorCtx(r.Context(), w, err)
		return
	}
	resp, err := logic.NewAncillaryLogic(r.Context(), h.svcCtx).AnnouncementDetail(&req)
	httpx.OkJsonCtx(r.Context(), w, common.NewResult().Deal(resp, err))
}

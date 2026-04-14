package handler

import (
	"errors"
	"net/http"
	"strconv"

	"ucenter-api/internal/logic"
	"ucenter-api/internal/svc"
	"ucenter-api/internal/types"

	common "mscoin-common"
	"github.com/zeromicro/go-zero/rest/httpx"
)

type ActivityHandler struct {
	svcCtx *svc.ServiceContext
}

func NewActivityHandler(svcCtx *svc.ServiceContext) *ActivityHandler {
	return &ActivityHandler{
		svcCtx: svcCtx,
	}
}

// ActivityPageQuery 活动列表分页查询
func (h *ActivityHandler) ActivityPageQuery(w http.ResponseWriter, r *http.Request) {
	var req types.ActivityPageQueryReq
	if err := httpx.Parse(r, &req); err != nil {
		httpx.ErrorCtx(r.Context(), w, err)
		return
	}

	l := logic.NewActivityLogic(h.svcCtx)
	resp, err := l.ActivityPageQuery(&req)
	httpx.OkJsonCtx(r.Context(), w, common.NewResult().Deal(resp, err))
}

// ActivityDetail 活动详情
func (h *ActivityHandler) ActivityDetail(w http.ResponseWriter, r *http.Request) {
	var req types.ActivityDetailReq
	// 从 URL query 参数解析 id
	req.Id, _ = strconv.ParseInt(r.URL.Query().Get("id"), 10, 64)

	if req.Id == 0 {
		httpx.ErrorCtx(r.Context(), w, errors.New("id is required"))
		return
	}

	l := logic.NewActivityLogic(h.svcCtx)
	resp, err := l.ActivityDetail(&req)
	httpx.OkJsonCtx(r.Context(), w, common.NewResult().Deal(resp, err))
}

// ActivityAttend 参与活动
func (h *ActivityHandler) ActivityAttend(w http.ResponseWriter, r *http.Request) {
	var req types.ActivityAttendReq
	if err := httpx.Parse(r, &req); err != nil {
		httpx.ErrorCtx(r.Context(), w, err)
		return
	}

	l := logic.NewActivityLogic(h.svcCtx)
	resp, err := l.ActivityAttend(&req)
	httpx.OkJsonCtx(r.Context(), w, common.NewResult().Deal(resp, err))
}

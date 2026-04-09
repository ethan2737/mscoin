package logic

import (
	"context"
	"errors"
	"strings"

	"mscoin-common/pages"
	"ucenter-api/internal/svc"
	"ucenter-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AncillaryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAncillaryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AncillaryLogic {
	return &AncillaryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func normalizeLang(lang string) string {
	if strings.EqualFold(lang, "EN") {
		return "EN"
	}
	return "CN"
}

func clampPage(pageNo, pageSize int64) (int64, int64) {
	if pageNo <= 0 {
		pageNo = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}
	return pageNo, pageSize
}

func sliceForPage[T any](items []T, pageNo, pageSize int64) []T {
	start := int((pageNo - 1) * pageSize)
	if start >= len(items) {
		return []T{}
	}
	end := start + int(pageSize)
	if end > len(items) {
		end = len(items)
	}
	return items[start:end]
}

func helpSummary(seed contentSeed, lang string) *types.ContentSummary {
	title := seed.TitleCN
	if normalizeLang(lang) == "EN" {
		title = seed.TitleEN
	}
	return &types.ContentSummary{
		Id:         seed.ID,
		Title:      title,
		CreateTime: seed.CreateTime,
		IsTop:      seed.IsTop,
	}
}

func helpDetail(seed contentSeed, lang string) *types.ContentDetail {
	title := seed.TitleCN
	content := seed.ContentCN
	if normalizeLang(lang) == "EN" {
		title = seed.TitleEN
		content = seed.ContentEN
	}
	return &types.ContentDetail{
		Id:         seed.ID,
		Title:      title,
		Content:    content,
		CreateTime: seed.CreateTime,
	}
}

func announcementSummary(seed contentSeed, lang string) *types.ContentSummary {
	title := seed.TitleCN
	if normalizeLang(lang) == "EN" {
		title = seed.TitleEN
	}
	return &types.ContentSummary{
		Id:         seed.ID,
		Title:      title,
		CreateTime: seed.CreateTime,
		IsTop:      seed.IsTop,
	}
}

func announcementDetail(seed contentSeed, lang string) *types.ContentDetail {
	title := seed.TitleCN
	content := seed.ContentCN
	if normalizeLang(lang) == "EN" {
		title = seed.TitleEN
		content = seed.ContentEN
	}
	return &types.ContentDetail{
		Id:         seed.ID,
		Title:      title,
		Content:    content,
		CreateTime: seed.CreateTime,
	}
}

func (l *AncillaryLogic) HelpOverview(req *types.ContentReq) ([]*types.HelpSection, error) {
	lang := normalizeLang(req.Lang)
	result := make([]*types.HelpSection, 0, len(helpCategorySeeds))
	for _, category := range helpCategorySeeds {
		section := &types.HelpSection{
			Cate:    category.Cate,
			TitleCN: category.TitleCN,
			TitleEN: category.TitleEN,
			Content: make([]*types.ContentSummary, 0),
		}
		for _, article := range helpArticleSeeds {
			if article.Cate != category.Cate {
				continue
			}
			section.Content = append(section.Content, helpSummary(article, lang))
			if len(section.Content) >= 5 {
				break
			}
		}
		result = append(result, section)
	}
	return result, nil
}

func (l *AncillaryLogic) HelpPage(req *types.ContentReq) (*pages.PageResult, error) {
	lang := normalizeLang(req.Lang)
	pageNo, pageSize := clampPage(req.PageNo, req.PageSize)
	list := make([]*types.ContentSummary, 0)
	for _, article := range helpArticleSeeds {
		if req.Cate > 0 && article.Cate != req.Cate {
			continue
		}
		list = append(list, helpSummary(article, lang))
	}
	current := sliceForPage(list, pageNo, pageSize)
	content := make([]any, len(current))
	for i := range current {
		content[i] = current[i]
	}
	return pages.New(content, pageNo, pageSize, int64(len(list))), nil
}

func (l *AncillaryLogic) HelpTop(req *types.ContentReq) ([]*types.ContentSummary, error) {
	lang := normalizeLang(req.Lang)
	result := make([]*types.ContentSummary, 0)
	for _, article := range helpArticleSeeds {
		if req.Cate > 0 && article.Cate != req.Cate {
			continue
		}
		result = append(result, helpSummary(article, lang))
		if len(result) >= 5 {
			break
		}
	}
	return result, nil
}

func (l *AncillaryLogic) HelpDetail(req *types.ContentReq) (*types.ContentDetail, error) {
	for _, article := range helpArticleSeeds {
		if article.ID == req.Id {
			return helpDetail(article, req.Lang), nil
		}
	}
	return nil, errors.New("help article not found")
}

func (l *AncillaryLogic) AnnouncementPage(req *types.ContentReq) (*pages.PageResult, error) {
	lang := normalizeLang(req.Lang)
	pageNo, pageSize := clampPage(req.PageNo, req.PageSize)
	summaries := make([]*types.ContentSummary, 0, len(announcementSeeds))
	for _, seed := range announcementSeeds {
		summaries = append(summaries, announcementSummary(seed, lang))
	}
	current := sliceForPage(summaries, pageNo, pageSize)
	content := make([]any, len(current))
	for i := range current {
		content[i] = current[i]
	}
	return pages.New(content, pageNo, pageSize, int64(len(summaries))), nil
}

func (l *AncillaryLogic) AnnouncementDetail(req *types.ContentReq) (*types.AnnouncementMoreResp, error) {
	if len(announcementSeeds) == 0 {
		return nil, errors.New("announcement not found")
	}
	index := -1
	for i, seed := range announcementSeeds {
		if seed.ID == req.Id {
			index = i
			break
		}
	}
	if index == -1 {
		index = 0
	}
	resp := &types.AnnouncementMoreResp{
		Info: announcementDetail(announcementSeeds[index], req.Lang),
	}
	if index > 0 {
		resp.Back = announcementSummary(announcementSeeds[index-1], req.Lang)
	}
	if index+1 < len(announcementSeeds) {
		resp.Next = announcementSummary(announcementSeeds[index+1], req.Lang)
	}
	return resp, nil
}

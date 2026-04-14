package logic

import (
	"ucenter-api/internal/svc"
	"ucenter-api/internal/types"
)

// ActivityLogic 活动逻辑
type ActivityLogic struct {
	svcCtx *svc.ServiceContext
}

// NewActivityLogic 创建活动逻辑实例
func NewActivityLogic(svcCtx *svc.ServiceContext) *ActivityLogic {
	return &ActivityLogic{
		svcCtx: svcCtx,
	}
}

// activitySeed 活动种子数据
type activitySeed struct {
	ID             int64   `json:"id"`
	Title          string  `json:"title"`
	Detail         string  `json:"detail"`
	SmallImageUrl  string  `json:"smallImageUrl"`
	Type           int64   `json:"type"`           // 活动类型 1-5
	Step           int64   `json:"step"`           // 活动状态 0-3
	StartTime      int64   `json:"startTime"`      // 开始时间戳 (毫秒)
	EndTime        int64   `json:"endTime"`        // 结束时间戳 (毫秒)
	TotalSupply    float64 `json:"totalSupply"`    // 总量
	TradedAmount   float64 `json:"tradedAmount"`   // 已兑数量
	Progress       float64 `json:"progress"`       // 进度百分比
	Price          float64 `json:"price"`          // 价格
	PriceScale     int     `json:"priceScale"`     // 价格精度
	Unit           string  `json:"unit"`           // 活动货币单位
	AcceptUnit     string  `json:"acceptUnit"`     // 接受货币单位
	AmountScale    int     `json:"amountScale"`    // 数量精度
	MaxLimitAmout  float64 `json:"maxLimitAmout"`  // 最大限购量
	MinLimitAmout  float64 `json:"minLimitAmout"`  // 最小限购量
	LimitTimes     int     `json:"limitTimes"`     // 限购次数
}

// activitySeeds Mock 活动数据
var activitySeeds = []activitySeed{
	{
		ID:            1,
		Title:         "MSC 首次上线抢购",
		Detail:        "MSC 代币首次发行，限时抢购",
		SmallImageUrl: "/static/images/activity1.png",
		Type:          1, // 抢购发行
		Step:          1, // 进行中
		StartTime:     1712966400000, // 2024-04-13 00:00:00
		EndTime:       1713225600000, // 2024-04-16 00:00:00
		TotalSupply:   1000000,
		TradedAmount:  650000,
		Progress:      65,
		Price:         0.1,
		PriceScale:    2,
		Unit:          "MSC",
		AcceptUnit:    "USDT",
		AmountScale:   2,
		MaxLimitAmout: 10000,
		MinLimitAmout: 100,
		LimitTimes:    5,
	},
	{
		ID:            2,
		Title:         "BTC 分摊发行",
		Detail:        "BTC 代币分摊发行，人人有份",
		SmallImageUrl: "/static/images/activity2.png",
		Type:          2, // 分摊发行
		Step:          1, // 进行中
		StartTime:     1712880000000,
		EndTime:       1713139200000,
		TotalSupply:   500000,
		TradedAmount:  500000,
		Progress:      100,
		Price:         0.05,
		PriceScale:    2,
		Unit:          "BTC",
		AcceptUnit:    "USDT",
		AmountScale:   4,
		MaxLimitAmout: 5000,
		MinLimitAmout: 50,
		LimitTimes:    3,
	},
	{
		ID:            3,
		Title:         "持仓瓜分 ETH",
		Detail:        "持仓瓜分 ETH，按持仓比例分配",
		SmallImageUrl: "/static/images/activity3.png",
		Type:          3, // 持仓瓜分
		Step:          0, // 即将开始
		StartTime:     1713312000000,
		EndTime:       1713571200000,
		TotalSupply:   200000,
		TradedAmount:  0,
		Progress:      0,
		Price:         0.08,
		PriceScale:    2,
		Unit:          "ETH",
		AcceptUnit:    "USDT",
		AmountScale:   4,
		MaxLimitAmout: 0,
		MinLimitAmout: 0,
		LimitTimes:    0,
	},
	{
		ID:            4,
		Title:         "自由认购 DOT",
		Detail:        "DOT 自由认购，多买多送",
		SmallImageUrl: "/static/images/activity4.png",
		Type:          4, // 自由认购
		Step:          2, // 派发中
		StartTime:     1712620800000,
		EndTime:       1712880000000,
		TotalSupply:   300000,
		TradedAmount:  280000,
		Progress:      93.33,
		Price:         0.12,
		PriceScale:    2,
		Unit:          "DOT",
		AcceptUnit:    "USDT",
		AmountScale:   2,
		MaxLimitAmout: 20000,
		MinLimitAmout: 200,
		LimitTimes:    10,
	},
	{
		ID:            5,
		Title:         "云矿机认购",
		Detail:        "购买云矿机，每日产币收益",
		SmallImageUrl: "/static/images/activity5.png",
		Type:          5, // 云矿机认购
		Step:          3, // 已完成
		StartTime:     1712361600000,
		EndTime:       1712620800000,
		TotalSupply:   1000,
		TradedAmount:  1000,
		Progress:      100,
		Price:         100,
		PriceScale:    2,
		Unit:          "MINER",
		AcceptUnit:    "USDT",
		AmountScale:   0,
		MaxLimitAmout: 100,
		MinLimitAmout: 1,
		LimitTimes:    0,
	},
}

// ActivityPageQuery 活动列表分页查询
func (l *ActivityLogic) ActivityPageQuery(req *types.ActivityPageQueryReq) (*types.ActivityPageQueryResp, error) {
	pageNo, pageSize := req.PageNo, req.PageSize
	if pageNo <= 0 {
		pageNo = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	// 筛选活动
	filteredList := make([]activitySeed, 0)
	for _, seed := range activitySeeds {
		if req.Step == -1 || seed.Step == int64(req.Step) {
			filteredList = append(filteredList, seed)
		}
	}

	// 分页
	total := int64(len(filteredList))
	start := (pageNo - 1) * pageSize
	end := start + pageSize
	if start >= len(filteredList) {
		start = len(filteredList)
	}
	if end > len(filteredList) {
		end = len(filteredList)
	}

	pageItems := filteredList[start:end]

	// 转换为响应格式
	content := make([]*types.Activity, len(pageItems))
	for i, item := range pageItems {
		content[i] = &types.Activity{
			Id:            item.ID,
			Title:         item.Title,
			Detail:        item.Detail,
			SmallImageUrl: item.SmallImageUrl,
			Type:          item.Type,
			Step:          item.Step,
			StartTime:     item.StartTime,
			EndTime:       item.EndTime,
			TotalSupply:   item.TotalSupply,
			TradedAmount:  item.TradedAmount,
			Progress:      item.Progress,
			Price:         item.Price,
			PriceScale:    item.PriceScale,
			Unit:          item.Unit,
			AcceptUnit:    item.AcceptUnit,
			AmountScale:   item.AmountScale,
			MaxLimitAmout: item.MaxLimitAmout,
			MinLimitAmout: item.MinLimitAmout,
			LimitTimes:    item.LimitTimes,
		}
	}

	return &types.ActivityPageQueryResp{
		Content:       content,
		TotalElements: total,
	}, nil
}

// ActivityDetail 活动详情
func (l *ActivityLogic) ActivityDetail(req *types.ActivityDetailReq) (*types.Activity, error) {
	for _, seed := range activitySeeds {
		if seed.ID == req.Id {
			return &types.Activity{
				Id:            seed.ID,
				Title:         seed.Title,
				Detail:        seed.Detail,
				SmallImageUrl: seed.SmallImageUrl,
				Type:          seed.Type,
				Step:          seed.Step,
				StartTime:     seed.StartTime,
				EndTime:       seed.EndTime,
				TotalSupply:   seed.TotalSupply,
				TradedAmount:  seed.TradedAmount,
				Progress:      seed.Progress,
				Price:         seed.Price,
				PriceScale:    seed.PriceScale,
				Unit:          seed.Unit,
				AcceptUnit:    seed.AcceptUnit,
				AmountScale:   seed.AmountScale,
				MaxLimitAmout: seed.MaxLimitAmout,
				MinLimitAmout: seed.MinLimitAmout,
				LimitTimes:    seed.LimitTimes,
			}, nil
		}
	}
	return nil, nil
}

// ActivityAttend 参与活动
func (l *ActivityLogic) ActivityAttend(req *types.ActivityAttendReq) (*types.ActivityAttendResp, error) {
	// TODO: 实现参与活动逻辑
	// 目前返回成功
	return &types.ActivityAttendResp{
		Success: true,
	}, nil
}

package handler

import (
	"context"
	"errors"
	"fmt"
	"math"
	"net/http"
	"strconv"
	"strings"
	"time"

	"grpc-common/market/mclient"
	"swap-api/internal/logic"
	"swap-api/internal/marketdata"
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
	if err != nil {
		httpx.ErrorCtx(r.Context(), w, err)
		return
	}

	symbolParts := strings.Split(resp.Symbol, "/")
	coinSymbol := symbolParts[0]
	baseSymbol := symbolParts[1]

	data := map[string]any{
		"id":                resp.Id,
		"symbol":            resp.Symbol,
		"coinSymbol":        coinSymbol,
		"baseSymbol":        baseSymbol,
		"type":              "ALWAYS",
		"contractType":      resp.ContractType,
		"pricePrecision":    resp.PricePrecision,
		"quantityPrecision": resp.QuantityPrecision,
		"shareNumber":       resp.ShareNumber,
		"makerFee":          resp.MakerFee,
		"takerFee":          resp.TakerFee,
		"minLeverage":       resp.MinLeverage,
		"maxLeverage":       resp.MaxLeverage,
		"status":            resp.Status,
		"coinScale":         4,
		"baseCoinScale":     4,
		"enableMarketBuy":   true,
		"enableMarketSell":  true,
		"exchangeable":      true,
	}

	result := map[string]any{
		"code":    0,
		"message": "success",
		"data":    data,
	}
	httpx.OkJsonCtx(r.Context(), w, result)
}

func (h *ContractInfoHandler) List(w http.ResponseWriter, r *http.Request) {
	req := &types.GetContractListReq{}
	l := logic.NewContractInfoLogic(r.Context(), h.svcCtx)
	resp, err := l.List(req)
	if err != nil {
		httpx.ErrorCtx(r.Context(), w, err)
		return
	}

	thumbMap := h.loadSymbolThumbMap(r.Context())
	data := make([]map[string]any, 0)
	for _, coin := range *resp {
		thumb := thumbMap[coin.Symbol]
		basePrice := fallbackSymbolPrice(coin.Symbol)
		high := basePrice * 1.05
		low := basePrice * 0.95
		volume := 1234.56
		chg := 0.0234
		rose := "+2.34%"
		if thumb != nil {
			basePrice = thumb.Close
			high = thumb.High
			low = thumb.Low
			volume = thumb.Volume
			chg = thumb.Chg
			rose = formatRose(thumb.Chg)
		}
		coinSymbol, baseSymbol := splitSymbol(coin.Symbol)
		data = append(data, map[string]any{
			"id":      coin.Id,
			"symbol":  coin.Symbol,
			"coin":    coinSymbol,
			"base":    baseSymbol,
			"name":    coin.Symbol,
			"close":   basePrice,
			"price":   basePrice,
			"chg":     chg,
			"rose":    rose,
			"high":    high,
			"low":     low,
			"volume":  volume,
			"zone":    0,
			"href":    coin.Id,
			"isFavor": false,
		})
	}

	result := map[string]any{
		"code":    0,
		"message": "success",
		"data":    data,
	}
	httpx.OkJsonCtx(r.Context(), w, result)
}

func (h *ContractInfoHandler) SymbolThumb(w http.ResponseWriter, r *http.Request) {
	req := &types.GetContractListReq{}
	l := logic.NewContractInfoLogic(r.Context(), h.svcCtx)
	resp, err := l.List(req)
	result := common.NewResult().Deal(resp, err)
	httpx.OkJsonCtx(r.Context(), w, result)
}

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
		if svc.MarketDataClient != nil {
			plate, err := loadTradePlate(r.Context(), r, svc, false)
			if err == nil && plate != nil {
				httpx.OkJsonCtx(r.Context(), w, map[string]any{
					"code":    0,
					"message": "success",
					"data":    plate,
				})
				return
			}
		}
		result := map[string]any{
			"code":    0,
			"message": "success",
			"data": map[string]any{
				"asks": []map[string]any{
					{"price": 64200.5, "amount": 0.5, "totalAmount": 0.5},
					{"price": 64210.0, "amount": 1.2, "totalAmount": 1.7},
					{"price": 64220.5, "amount": 0.8, "totalAmount": 2.5},
					{"price": 64230.0, "amount": 2.0, "totalAmount": 4.5},
					{"price": 64250.0, "amount": 1.5, "totalAmount": 6.0},
				},
				"bids": []map[string]any{
					{"price": 64190.0, "amount": 0.6, "totalAmount": 0.6},
					{"price": 64180.5, "amount": 1.0, "totalAmount": 1.6},
					{"price": 64170.0, "amount": 0.9, "totalAmount": 2.5},
					{"price": 64160.5, "amount": 1.8, "totalAmount": 4.3},
					{"price": 64150.0, "amount": 2.2, "totalAmount": 6.5},
				},
			},
		}
		httpx.OkJsonCtx(r.Context(), w, result)
	}
}

func exchangePlateFullHandler(svc *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if svc.MarketDataClient != nil {
			plate, err := loadTradePlate(r.Context(), r, svc, true)
			if err == nil && plate != nil {
				httpx.OkJsonCtx(r.Context(), w, map[string]any{
					"code":    0,
					"message": "success",
					"data":    plate,
				})
				return
			}
		}
		result := map[string]any{
			"code":    0,
			"message": "success",
			"data": map[string]any{
				"asks": []map[string]any{
					{"price": 64200.5, "amount": 0.5, "totalAmount": 0.5},
					{"price": 64210.0, "amount": 1.2, "totalAmount": 1.7},
					{"price": 64220.5, "amount": 0.8, "totalAmount": 2.5},
					{"price": 64230.0, "amount": 2.0, "totalAmount": 4.5},
					{"price": 64250.0, "amount": 1.5, "totalAmount": 6.0},
				},
				"bids": []map[string]any{
					{"price": 64190.0, "amount": 0.6, "totalAmount": 0.6},
					{"price": 64180.5, "amount": 1.0, "totalAmount": 1.6},
					{"price": 64170.0, "amount": 0.9, "totalAmount": 2.5},
					{"price": 64160.5, "amount": 1.8, "totalAmount": 4.3},
					{"price": 64150.0, "amount": 2.2, "totalAmount": 6.5},
				},
			},
		}
		httpx.OkJsonCtx(r.Context(), w, result)
	}
}

func latestTradeHandler(svc *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if svc.MarketDataClient != nil {
			trades, err := loadLatestTrade(r.Context(), r, svc)
			if err == nil && len(trades) > 0 {
				httpx.OkJsonCtx(r.Context(), w, map[string]any{
					"code":    0,
					"message": "success",
					"data":    trades,
				})
				return
			}
		}
		now := time.Now().UnixMilli()
		result := map[string]any{
			"code":    0,
			"message": "success",
			"data": []map[string]any{
				{"price": 64195.0, "amount": 0.15, "time": now - 5000, "direction": "BUY"},
				{"price": 64198.5, "amount": 0.25, "time": now - 10000, "direction": "SELL"},
				{"price": 64200.0, "amount": 0.5, "time": now - 15000, "direction": "BUY"},
				{"price": 64195.5, "amount": 0.3, "time": now - 20000, "direction": "SELL"},
				{"price": 64190.0, "amount": 0.8, "time": now - 25000, "direction": "SELL"},
			},
		}
		httpx.OkJsonCtx(r.Context(), w, result)
	}
}

func (h *ContractInfoHandler) loadSymbolThumbMap(ctx context.Context) map[string]*marketdata.SymbolThumb {
	result := make(map[string]*marketdata.SymbolThumb)
	if h.svcCtx.MarketDataClient == nil {
		return result
	}
	thumbs, err := h.svcCtx.MarketDataClient.LoadSymbolThumb(ctx)
	if err != nil {
		return result
	}
	for _, thumb := range thumbs {
		if thumb == nil {
			continue
		}
		result[thumb.Symbol] = thumb
	}
	return result
}

func loadTradePlate(ctx context.Context, r *http.Request, svc *svc.ServiceContext, full bool) (map[string]any, error) {
	symbol := "BTC/USDT"
	var req struct {
		Symbol string `json:"symbol"`
	}
	if err := httpx.ParseJsonBody(r, &req); err == nil && req.Symbol != "" {
		symbol = req.Symbol
	}
	plate, err := svc.MarketDataClient.LoadTradePlate(ctx, symbol, full)
	if err != nil || plate == nil {
		return nil, err
	}
	return map[string]any{
		"ask":  plate.Ask,
		"bid":  plate.Bid,
		"asks": convertTradePlateItems(plate.Ask),
		"bids": convertTradePlateItems(plate.Bid),
	}, nil
}

func loadLatestTrade(ctx context.Context, r *http.Request, svc *svc.ServiceContext) ([]*marketdata.LatestTrade, error) {
	symbol := "BTC/USDT"
	size := 20
	var req struct {
		Symbol string `json:"symbol"`
		Size   int    `json:"size"`
	}
	if err := httpx.ParseJsonBody(r, &req); err == nil {
		if req.Symbol != "" {
			symbol = req.Symbol
		}
		if req.Size > 0 {
			size = req.Size
		}
	}
	return svc.MarketDataClient.LoadLatestTrade(ctx, symbol, size)
}

func convertTradePlateItems(side *marketdata.TradePlateSide) []map[string]any {
	if side == nil || len(side.Items) == 0 {
		return []map[string]any{}
	}
	total := 0.0
	items := make([]map[string]any, 0, len(side.Items))
	for _, item := range side.Items {
		if item == nil {
			continue
		}
		total += item.Amount
		items = append(items, map[string]any{
			"price":       item.Price,
			"amount":      item.Amount,
			"totalAmount": total,
		})
	}
	return items
}

func klineHistoryHandler(svc *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		symbol := r.URL.Query().Get("symbol")
		fromStr := r.URL.Query().Get("from")
		toStr := r.URL.Query().Get("to")
		resolution := r.URL.Query().Get("resolution")

		if symbol == "" || fromStr == "" || toStr == "" || resolution == "" {
			httpx.ErrorCtx(r.Context(), w, errors.New("missing required parameters"))
			return
		}

		from, err := strconv.ParseInt(fromStr, 10, 64)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, errors.New("invalid from parameter"))
			return
		}

		to, err := strconv.ParseInt(toStr, 10, 64)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, errors.New("invalid to parameter"))
			return
		}

		ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
		defer cancel()

		list := make([][]any, 0)
		if svc.MarketRpc != nil {
			resp, err := svc.MarketRpc.HistoryKline(ctx, &mclient.MarketReq{
				Symbol:     symbol,
				From:       from,
				To:         to,
				Resolution: resolution,
			})
			if err == nil && resp != nil {
				for _, item := range resp.List {
					list = append(list, []any{
						item.Time,
						item.Open,
						item.High,
						item.Low,
						item.Close,
						item.Volume,
					})
				}
			}
		}

		if len(list) == 0 {
			list = buildFallbackSwapKline(symbol, from, to, resolution)
		}

		result := map[string]any{
			"code":    0,
			"message": "success",
			"data":    list,
		}
		httpx.OkJsonCtx(r.Context(), w, result)
	}
}

func buildFallbackSwapKline(symbol string, from, to int64, resolution string) [][]any {
	step := swapResolutionMillis(resolution)
	if step <= 0 {
		step = int64(time.Hour / time.Millisecond)
	}
	if to <= from {
		to = from + step*200
	}

	count := int((to-from)/step) + 1
	if count < 1 {
		count = 1
	}
	if count > 300 {
		count = 300
		from = to - int64(count-1)*step
	}

	base := fallbackSymbolPrice(symbol)
	list := make([][]any, 0, count)
	for i := 0; i < count; i++ {
		ts := from + int64(i)*step
		wave := math.Sin(float64(i) / 6)
		drift := math.Cos(float64(i) / 9)
		open := base * (1 + wave*0.002)
		close := open * (1 + drift*0.0015)
		high := math.Max(open, close) * 1.0012
		low := math.Min(open, close) * 0.9988
		volume := 20 + math.Abs(wave)*15 + float64(i%7)
		list = append(list, []any{
			ts,
			roundFloat(open, 4),
			roundFloat(high, 4),
			roundFloat(low, 4),
			roundFloat(close, 4),
			roundFloat(volume, 4),
		})
	}
	return list
}

func swapResolutionMillis(resolution string) int64 {
	switch resolution {
	case "1":
		return int64(time.Minute / time.Millisecond)
	case "5":
		return int64((5 * time.Minute) / time.Millisecond)
	case "15":
		return int64((15 * time.Minute) / time.Millisecond)
	case "30":
		return int64((30 * time.Minute) / time.Millisecond)
	case "60":
		return int64(time.Hour / time.Millisecond)
	case "240":
		return int64((4 * time.Hour) / time.Millisecond)
	case "1D":
		return int64((24 * time.Hour) / time.Millisecond)
	case "1W":
		return int64((7 * 24 * time.Hour) / time.Millisecond)
	case "1M":
		return int64((30 * 24 * time.Hour) / time.Millisecond)
	default:
		return 0
	}
}

func fallbackSymbolPrice(symbol string) float64 {
	switch symbol {
	case "BTC/USDT":
		return 64200
	case "ETH/USDT":
		return 3450
	default:
		return 100
	}
}

func roundFloat(value float64, scale int) float64 {
	pow := math.Pow10(scale)
	return math.Round(value*pow) / pow
}

func splitSymbol(symbol string) (string, string) {
	parts := strings.Split(symbol, "/")
	if len(parts) == 2 {
		return parts[0], parts[1]
	}
	return symbol, ""
}

func formatRose(chg float64) string {
	return fmt.Sprintf("%+.2f%%", chg*100)
}

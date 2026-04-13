package logic

import (
	"fmt"
	"strconv"

	"mscoin-common/pages"
	"swap-api/internal/model"
	"swap-api/internal/types"
)

func resolveTransferType(req *types.ContractTransferReq) int32 {
	if req.Type != 0 {
		return req.Type
	}
	return req.Direction
}

func parsePageValue(raw any, fallback int64) int64 {
	switch value := raw.(type) {
	case nil:
		return fallback
	case float64:
		return int64(value)
	case float32:
		return int64(value)
	case int:
		return int64(value)
	case int32:
		return int64(value)
	case int64:
		return value
	case string:
		if value == "" {
			return fallback
		}
		parsed, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return fallback
		}
		return parsed
	default:
		return fallback
	}
}

func buildUcOrderPage(orders []*model.ContractOrder, coinMap map[int64]*model.ContractCoin, pageNo, pageSize int64) *pages.PageResult {
	content := make([]any, 0, len(orders))
	for _, order := range orders {
		content = append(content, buildUcOrderItem(order, coinMap[order.ContractCoinId]))
	}
	return pages.New(content, pageNo, pageSize, int64(len(content)))
}

func buildSwapPositionItems(positions []*model.ContractPosition, coinMap map[int64]*model.ContractCoin) []*types.SwapPositionItem {
	items := make([]*types.SwapPositionItem, 0, len(positions))
	for _, position := range positions {
		items = append(items, buildSwapPositionItem(position, coinMap[position.ContractCoinId]))
	}
	return items
}

func buildUcOrderItem(order *model.ContractOrder, coin *model.ContractCoin) *types.UcOrderItem {
	coinName := ""
	contractCoinType := "ALWAYS"
	if coin != nil {
		coinName = coin.Symbol
		if coin.ContractType == 2 {
			contractCoinType = "SECOND"
		}
	}

	return &types.UcOrderItem{
		OrderId:          order.Id,
		ContractCoinId:   order.ContractCoinId,
		ContractCoinName: coinName,
		ContractCoinType: contractCoinType,
		HoldTime:         0,
		Direction:        mapOrderDirection(order.Direction),
		EntrustType:      mapOrderEntrustType(order.Type),
		Type:             mapPriceType(order.EntrustType),
		Leverage:         order.Leverage,
		EntrustPrice:     order.Price,
		StrikePrice:      0,
		TriggerPrice:     0,
		Share:            order.Amount,
		PrincipalAmount:  calculatePrincipal(order),
		Fee:              order.Fee,
		Profit:           order.Profit,
		Status:           mapOrderStatus(order.Status),
		CreateTime:       order.CreatedAt.Format("2006-01-02 15:04:05"),
	}
}

func buildSwapPositionItem(position *model.ContractPosition, coin *model.ContractCoin) *types.SwapPositionItem {
	coinName := ""
	contractCoinType := "ALWAYS"
	if coin != nil {
		coinName = coin.Symbol
		if coin.ContractType == 2 {
			contractCoinType = "SECOND"
		}
	}

	return &types.SwapPositionItem{
		Id:                position.Id,
		ContractCoinId:    position.ContractCoinId,
		ContractCoinName:  coinName,
		ContractCoinType:  contractCoinType,
		HoldTime:          0,
		Direction:         mapOrderDirection(position.Direction),
		Multiple:          position.Leverage,
		OpenPrice:         position.EntryPrice,
		PrincipalAmount:   position.Margin,
		Position:          position.Size,
		AvailablePosition: position.Size,
		LiquidationPrice:  position.LiquidationPrice,
		Profit:            position.UnrealizedPnl,
	}
}

func buildWalletItems(wallets []*model.ContractWallet) []*types.ContractWalletItem {
	items := make([]*types.ContractWalletItem, 0, len(wallets))
	for _, wallet := range wallets {
		items = append(items, &types.ContractWalletItem{
			Coin: &types.WalletCoinItem{
				Unit:    wallet.Unit,
				Name:    wallet.Unit,
				UsdRate: 1,
				CnyRate: 7.2,
			},
			Balance:       wallet.Balance,
			FrozenBalance: wallet.Frozen,
			ToReleased:    0,
			MainBalance:   wallet.MainBalance,
		})
	}
	return items
}

func buildTransactionPage(transactions []*model.ContractTransaction, pageNo, pageSize int64) *pages.PageResult {
	content := make([]any, 0, len(transactions))
	for _, tx := range transactions {
		content = append(content, &types.ContractTransactionItem{
			CreateTime: tx.CreatedAt.Format("2006-01-02 15:04:05"),
			Type:       tx.Type,
			Symbol:     tx.Unit,
			Amount:     tx.Amount,
		})
	}
	return pages.New(content, pageNo, pageSize, int64(len(content)))
}

func buildThumbItems(coins []*model.ContractCoin) []*types.ContractCoinThumbItem {
	items := make([]*types.ContractCoinThumbItem, 0, len(coins))
	for _, coin := range coins {
		items = append(items, &types.ContractCoinThumbItem{
			Id:   coin.Id,
			Name: coin.Symbol,
		})
	}
	return items
}

func mapOrderDirection(direction int32) string {
	if direction == 1 {
		return "BUY"
	}
	return "SELL"
}

func mapOrderEntrustType(orderType int32) string {
	if orderType == 1 {
		return "OPEN"
	}
	return "CLOSE"
}

func mapPriceType(entrustType int32) string {
	if entrustType == 1 {
		return "LIMIT_PRICE"
	}
	return "TRIGGER_PRICE"
}

func mapOrderStatus(status int32) string {
	switch status {
	case 1:
		return "ENTRUST_ING"
	case 2:
		return "ENTRUST_SUCCESS"
	case 3:
		return "ENTRUST_CANCEL"
	default:
		return fmt.Sprintf("UNKNOWN_%d", status)
	}
}

func calculatePrincipal(order *model.ContractOrder) float64 {
	if order.Leverage <= 0 {
		return order.Price * order.Amount
	}
	return order.Price * order.Amount / float64(order.Leverage)
}

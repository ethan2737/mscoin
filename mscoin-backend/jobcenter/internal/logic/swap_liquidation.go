package logic

import (
	"context"
	"log"
	"sync"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"grpc-common/market/mclient"
	"jobcenter/internal/model"
)

type SwapLiquidationConfig struct {
	LiquidationRatio float64
}

type SwapLiquidation struct {
	wg             sync.WaitGroup
	liquidationDao *SwapLiquidationDao
	config         SwapLiquidationConfig
	marketRpc      mclient.Market
	cache          cache.Cache
}

type SwapLiquidationDao struct {
	positionModel model.ContractPositionModel
	walletModel   model.ContractWalletModel
	coinModel     model.ContractCoinModel
}

func NewSwapLiquidationDao(positionModel model.ContractPositionModel, walletModel model.ContractWalletModel, coinModel model.ContractCoinModel) *SwapLiquidationDao {
	return &SwapLiquidationDao{
		positionModel: positionModel,
		walletModel:   walletModel,
		coinModel:     coinModel,
	}
}

func (d *SwapLiquidationDao) UpdatePosition(ctx context.Context, position *model.ContractPosition) error {
	return d.positionModel.Update(ctx, position)
}

func (d *SwapLiquidationDao) GetWallet(ctx context.Context, memberId int64, unit string) (*model.ContractWallet, error) {
	return d.walletModel.FindByMemberAndUnit(ctx, memberId, unit)
}

func (d *SwapLiquidationDao) UpdateWalletBalance(ctx context.Context, wallet *model.ContractWallet) error {
	return d.walletModel.UpdateBalance(ctx, wallet.Id, wallet.Balance, wallet.Frozen, wallet.MainBalance)
}

func NewSwapLiquidation(config SwapLiquidationConfig, marketRpc mclient.Market, positionModel model.ContractPositionModel, walletModel model.ContractWalletModel, coinModel model.ContractCoinModel, cache2 cache.Cache) *SwapLiquidation {
	return &SwapLiquidation{
		config:         config,
		marketRpc:      marketRpc,
		cache:          cache2,
		liquidationDao: NewSwapLiquidationDao(positionModel, walletModel, coinModel),
	}
}

func (l *SwapLiquidation) Scan() {
	ctx := context.Background()

	positions, err := l.getAllPositionsWithRisk(ctx)
	if err != nil {
		log.Printf("get all positions failed: %v", err)
		return
	}

	log.Printf("scan liquidation: found %d positions", len(positions))

	l.wg.Add(len(positions))
	for _, position := range positions {
		current := position
		go func() {
			defer l.wg.Done()
			l.checkAndLiquidate(ctx, current)
		}()
	}
	l.wg.Wait()
}

func (l *SwapLiquidation) getAllPositionsWithRisk(ctx context.Context) ([]*model.ContractPosition, error) {
	coins := l.getAllSwapCoins(ctx)

	var allPositions []*model.ContractPosition
	for _, coin := range coins {
		positions, err := l.liquidationDao.positionModel.FindByContractCoinId(ctx, coin.Id)
		if err != nil {
			log.Printf("get positions for coin %d failed: %v", coin.Id, err)
			continue
		}
		allPositions = append(allPositions, positions...)
	}

	riskPositions := make([]*model.ContractPosition, 0)
	for _, pos := range allPositions {
		if pos.Size > 0 {
			riskPositions = append(riskPositions, pos)
		}
	}

	return riskPositions, nil
}

func (l *SwapLiquidation) getAllSwapCoins(ctx context.Context) []*model.ContractCoin {
	coins, err := l.liquidationDao.coinModel.FindAll(ctx)
	if err != nil {
		log.Printf("get all contract coins failed: %v", err)
		return []*model.ContractCoin{}
	}

	var result []*model.ContractCoin
	for _, coin := range coins {
		if coin.Status == 1 {
			result = append(result, coin)
		}
	}

	return result
}

func (l *SwapLiquidation) checkAndLiquidate(ctx context.Context, position *model.ContractPosition) {
	currentPrice := l.getCurrentPrice(ctx, position.ContractCoinId)
	if currentPrice <= 0 {
		return
	}

	unrealizedPnl := l.calculateUnrealizedPnl(position, currentPrice)
	positionValue := currentPrice * position.Size
	requiredMargin := positionValue / float64(position.Leverage)
	marginRatio := (position.Margin + unrealizedPnl) / requiredMargin

	log.Printf("position %d margin ratio: %.4f, liquidation price: %.2f, current price: %.2f",
		position.Id, marginRatio, position.LiquidationPrice, currentPrice)

	if marginRatio < l.config.LiquidationRatio || currentPrice <= position.LiquidationPrice {
		log.Printf("trigger liquidation for position %d", position.Id)
		l.executeLiquidation(ctx, position, currentPrice)
	}
}

func (l *SwapLiquidation) calculateUnrealizedPnl(position *model.ContractPosition, currentPrice float64) float64 {
	if position.Direction == 1 {
		return (currentPrice - position.EntryPrice) * position.Size
	}

	return (position.EntryPrice - currentPrice) * position.Size
}

func (l *SwapLiquidation) getCurrentPrice(ctx context.Context, contractCoinId int64) float64 {
	coin, err := l.liquidationDao.coinModel.FindOne(ctx, contractCoinId)
	if err != nil || coin == nil {
		return 0
	}

	key := buildSwapRateCacheKey(coin.Symbol)

	var priceVal float64
	if err = l.cache.Get(key, &priceVal); err == nil {
		return priceVal
	}

	var priceText string
	if err = l.cache.Get(key, &priceText); err == nil {
		return parseSwapPriceString(priceText)
	}

	return 0
}

func (l *SwapLiquidation) executeLiquidation(ctx context.Context, position *model.ContractPosition, currentPrice float64) {
	unrealizedPnl := l.calculateUnrealizedPnl(position, currentPrice)

	wallet, err := l.liquidationDao.GetWallet(ctx, position.MemberId, "USDT")
	if err != nil {
		log.Printf("get wallet failed: %v", err)
		return
	}

	if wallet != nil {
		newBalance := wallet.Balance + unrealizedPnl
		if newBalance < 0 {
			newBalance = 0
		}

		err = l.liquidationDao.UpdateWalletBalance(ctx, &model.ContractWallet{
			Id:      wallet.Id,
			Balance: newBalance,
			Frozen:  wallet.Frozen,
		})
		if err != nil {
			log.Printf("update wallet failed: %v", err)
			return
		}
	}

	position.Size = 0
	position.UnrealizedPnl = 0
	position.Margin = 0
	err = l.liquidationDao.UpdatePosition(ctx, position)
	if err != nil {
		log.Printf("update position failed: %v", err)
		return
	}

	l.recordLiquidationTransaction(ctx, position, unrealizedPnl)
	log.Printf("liquidation executed for position %d, pnl: %.2f", position.Id, unrealizedPnl)
}

func (l *SwapLiquidation) recordLiquidationTransaction(ctx context.Context, position *model.ContractPosition, pnl float64) {
}

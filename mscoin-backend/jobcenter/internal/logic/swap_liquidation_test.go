package logic

import (
	"context"
	"errors"
	"testing"
	"time"

	"jobcenter/internal/model"
)

var errTestCacheNotFound = errors.New("cache not found")

type testContractCoinModel struct {
	coin *model.ContractCoin
}

func (m *testContractCoinModel) Insert(ctx context.Context, coin *model.ContractCoin) (int64, error) {
	return 0, nil
}

func (m *testContractCoinModel) FindOne(ctx context.Context, id int64) (*model.ContractCoin, error) {
	return m.coin, nil
}

func (m *testContractCoinModel) FindBySymbol(ctx context.Context, symbol string) (*model.ContractCoin, error) {
	return m.coin, nil
}

func (m *testContractCoinModel) FindAll(ctx context.Context) ([]*model.ContractCoin, error) {
	if m.coin == nil {
		return nil, nil
	}
	return []*model.ContractCoin{m.coin}, nil
}

func (m *testContractCoinModel) Update(ctx context.Context, coin *model.ContractCoin) error {
	return nil
}

func (m *testContractCoinModel) Delete(ctx context.Context, id int64) error {
	return nil
}

type testContractPositionModel struct{}

func (m *testContractPositionModel) Insert(ctx context.Context, position *model.ContractPosition) (int64, error) {
	return 0, nil
}

func (m *testContractPositionModel) FindOne(ctx context.Context, id int64) (*model.ContractPosition, error) {
	return nil, nil
}

func (m *testContractPositionModel) FindByMember(ctx context.Context, memberId int64, contractCoinId int64, direction int32) (*model.ContractPosition, error) {
	return nil, nil
}

func (m *testContractPositionModel) FindByMemberId(ctx context.Context, memberId int64) ([]*model.ContractPosition, error) {
	return nil, nil
}

func (m *testContractPositionModel) FindByContractCoinId(ctx context.Context, contractCoinId int64) ([]*model.ContractPosition, error) {
	return nil, nil
}

func (m *testContractPositionModel) Update(ctx context.Context, position *model.ContractPosition) error {
	return nil
}

func (m *testContractPositionModel) Upsert(ctx context.Context, position *model.ContractPosition) error {
	return nil
}

func (m *testContractPositionModel) Delete(ctx context.Context, id int64) error {
	return nil
}

type testContractWalletModel struct{}

func (m *testContractWalletModel) Insert(ctx context.Context, wallet *model.ContractWallet) (int64, error) {
	return 0, nil
}

func (m *testContractWalletModel) FindOne(ctx context.Context, id int64) (*model.ContractWallet, error) {
	return nil, nil
}

func (m *testContractWalletModel) FindByMemberAndUnit(ctx context.Context, memberId int64, unit string) (*model.ContractWallet, error) {
	return nil, nil
}

func (m *testContractWalletModel) UpdateBalance(ctx context.Context, id int64, balance, frozen, mainBalance float64) error {
	return nil
}

type testCache struct {
	values map[string]any
}

func (c *testCache) DelCtx(ctx context.Context, keys ...string) error {
	return nil
}

func (c *testCache) Del(keys ...string) error {
	return nil
}

func (c *testCache) Get(key string, v any) error {
	value, ok := c.values[key]
	if !ok {
		return errTestCacheNotFound
	}

	switch target := v.(type) {
	case *float64:
		number, ok := value.(float64)
		if !ok {
			return errTestCacheNotFound
		}
		*target = number
	case *string:
		text, ok := value.(string)
		if !ok {
			return errTestCacheNotFound
		}
		*target = text
	default:
		return errTestCacheNotFound
	}

	return nil
}

func (c *testCache) GetCtx(ctx context.Context, key string, v any) error {
	return c.Get(key, v)
}

func (c *testCache) IsNotFound(err error) bool {
	return errors.Is(err, errTestCacheNotFound)
}

func (c *testCache) Set(key string, v any) error {
	if c.values == nil {
		c.values = make(map[string]any)
	}
	c.values[key] = v
	return nil
}

func (c *testCache) SetCtx(ctx context.Context, key string, v any) error {
	return c.Set(key, v)
}

func (c *testCache) SetWithExpire(key string, v any, expire time.Duration) error {
	return c.Set(key, v)
}

func (c *testCache) SetWithExpireCtx(ctx context.Context, key string, v any, expire time.Duration) error {
	return c.Set(key, v)
}

func (c *testCache) Take(v any, key string, query func(v any) error) error {
	return query(v)
}

func (c *testCache) TakeCtx(ctx context.Context, v any, key string, query func(v any) error) error {
	return query(v)
}

func (c *testCache) TakeWithExpire(v any, key string, query func(v any, expire time.Duration) error) error {
	return query(v, time.Minute)
}

func (c *testCache) TakeWithExpireCtx(ctx context.Context, v any, key string, query func(v any, expire time.Duration) error) error {
	return query(v, time.Minute)
}

func TestGetCurrentPriceReadsRateCacheStringValue(t *testing.T) {
	const contractCoinID int64 = 7

	liquidation := NewSwapLiquidation(
		SwapLiquidationConfig{},
		nil,
		&testContractPositionModel{},
		&testContractWalletModel{},
		&testContractCoinModel{coin: &model.ContractCoin{
			Id:     contractCoinID,
			Symbol: "BTC/USDT",
			Status: 1,
		}},
		&testCache{
			values: map[string]any{
				buildSwapRateCacheKey("BTC/USDT"): "64123.45",
			},
		},
	)

	got := liquidation.getCurrentPrice(context.Background(), contractCoinID)
	if got != 64123.45 {
		t.Fatalf("expected cached swap rate, got %v", got)
	}
}

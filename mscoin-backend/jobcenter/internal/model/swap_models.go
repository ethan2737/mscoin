package model

import (
	"context"
	"time"
)

// ContractPosition 永续合约持仓
type ContractPosition struct {
	Id             int64     `db:"id"`
	MemberId       int64     `db:"member_id"`
	ContractCoinId int64     `db:"contract_coin_id"`
	Direction      int32     `db:"direction"`
	Leverage       int32     `db:"leverage"`
	Size           float64   `db:"size"`
	EntryPrice     float64   `db:"entry_price"`
	Margin         float64   `db:"margin"`
	UnrealizedPnl  float64   `db:"unrealized_pnl"`
	LiquidationPrice float64 `db:"liquidation_price"`
	CreatedAt      time.Time `db:"created_at"`
	UpdatedAt      time.Time `db:"updated_at"`
}

// ContractWallet 合约钱包
type ContractWallet struct {
	Id        int64     `db:"id"`
	MemberId  int64     `db:"member_id"`
	Unit      string    `db:"unit"`
	Balance   float64   `db:"balance"`
	Frozen    float64   `db:"frozen"`
	MainBalance float64 `db:"main_balance"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

// ContractCoin 合约配置
type ContractCoin struct {
	Id              int64     `db:"id"`
	Symbol          string    `db:"symbol"`
	ContractType    int32     `db:"contract_type"`
	PricePrecision  int32     `db:"price_precision"`
	QuantityPrecision int32   `db:"quantity_precision"`
	ShareNumber     float64   `db:"share_number"`
	MakerFee        float64   `db:"maker_fee"`
	TakerFee        float64   `db:"taker_fee"`
	MinLeverage     int32     `db:"min_leverage"`
	MaxLeverage     int32     `db:"max_leverage"`
	Status          int32     `db:"status"`
	CreatedAt       time.Time `db:"created_at"`
	UpdatedAt       time.Time `db:"updated_at"`
}

// ContractPositionModel 持仓模型接口
type ContractPositionModel interface {
	Insert(ctx context.Context, position *ContractPosition) (int64, error)
	FindOne(ctx context.Context, id int64) (*ContractPosition, error)
	FindByMember(ctx context.Context, memberId int64, contractCoinId int64, direction int32) (*ContractPosition, error)
	FindByMemberId(ctx context.Context, memberId int64) ([]*ContractPosition, error)
	FindByContractCoinId(ctx context.Context, contractCoinId int64) ([]*ContractPosition, error)
	Update(ctx context.Context, position *ContractPosition) error
	Upsert(ctx context.Context, position *ContractPosition) error
	Delete(ctx context.Context, id int64) error
}

// ContractWalletModel 钱包模型接口
type ContractWalletModel interface {
	Insert(ctx context.Context, wallet *ContractWallet) (int64, error)
	FindOne(ctx context.Context, id int64) (*ContractWallet, error)
	FindByMemberAndUnit(ctx context.Context, memberId int64, unit string) (*ContractWallet, error)
	UpdateBalance(ctx context.Context, id int64, balance, frozen, mainBalance float64) error
}

// ContractCoinModel 合约配置模型接口
type ContractCoinModel interface {
	Insert(ctx context.Context, coin *ContractCoin) (int64, error)
	FindOne(ctx context.Context, id int64) (*ContractCoin, error)
	FindBySymbol(ctx context.Context, symbol string) (*ContractCoin, error)
	FindAll(ctx context.Context) ([]*ContractCoin, error)
	Update(ctx context.Context, coin *ContractCoin) error
	Delete(ctx context.Context, id int64) error
}

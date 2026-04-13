package model

import (
	"context"
	"time"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ContractCoin struct {
	Id              int64   `db:"id"`
	Symbol          string  `db:"symbol"`
	ContractType    int32   `db:"contract_type"`
	PricePrecision  int32   `db:"price_precision"`
	QuantityPrecision int32 `db:"quantity_precision"`
	ShareNumber     float64 `db:"share_number"`
	MakerFee        float64 `db:"maker_fee"`
	TakerFee        float64 `db:"taker_fee"`
	MinLeverage     int32   `db:"min_leverage"`
	MaxLeverage     int32   `db:"max_leverage"`
	Status          int32   `db:"status"`
	CreatedAt       time.Time `db:"created_at"`
	UpdatedAt       time.Time `db:"updated_at"`
}

type ContractCoinModel interface {
	Insert(ctx context.Context, coin *ContractCoin) (int64, error)
	FindOne(ctx context.Context, id int64) (*ContractCoin, error)
	FindBySymbol(ctx context.Context, symbol string) (*ContractCoin, error)
	FindAll(ctx context.Context) ([]*ContractCoin, error)
	Update(ctx context.Context, coin *ContractCoin) error
	Delete(ctx context.Context, id int64) error
}

type defaultContractCoinModel struct {
	conn sqlx.SqlConn
	table string
}

func NewContractCoinModel(conn sqlx.SqlConn) ContractCoinModel {
	return &defaultContractCoinModel{
		conn: conn,
		table: "contract_coins",
	}
}

func (m *defaultContractCoinModel) Insert(ctx context.Context, coin *ContractCoin) (int64, error) {
	query := `INSERT INTO contract_coins (symbol, contract_type, price_precision, quantity_precision, share_number, maker_fee, taker_fee, min_leverage, max_leverage, status) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	res, err := m.conn.ExecCtx(ctx, query, coin.Symbol, coin.ContractType, coin.PricePrecision, coin.QuantityPrecision, coin.ShareNumber, coin.MakerFee, coin.TakerFee, coin.MinLeverage, coin.MaxLeverage, coin.Status)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	coin.Id = id
	return id, nil
}

func (m *defaultContractCoinModel) FindOne(ctx context.Context, id int64) (*ContractCoin, error) {
	query := `SELECT id, symbol, contract_type, price_precision, quantity_precision, share_number, maker_fee, taker_fee, min_leverage, max_leverage, status, created_at, updated_at FROM contract_coins WHERE id = ?`
	var resp ContractCoin
	err := m.conn.QueryRowCtx(ctx, &resp, query, id)
	if err != nil {
		if err == sqlx.ErrNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &resp, nil
}

func (m *defaultContractCoinModel) FindBySymbol(ctx context.Context, symbol string) (*ContractCoin, error) {
	query := `SELECT id, symbol, contract_type, price_precision, quantity_precision, share_number, maker_fee, taker_fee, min_leverage, max_leverage, status, created_at, updated_at FROM contract_coins WHERE symbol = ?`
	var resp ContractCoin
	err := m.conn.QueryRowCtx(ctx, &resp, query, symbol)
	if err != nil {
		if err == sqlx.ErrNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &resp, nil
}

func (m *defaultContractCoinModel) FindAll(ctx context.Context) ([]*ContractCoin, error) {
	query := `SELECT id, symbol, contract_type, price_precision, quantity_precision, share_number, maker_fee, taker_fee, min_leverage, max_leverage, status, created_at, updated_at FROM contract_coins`
	var resp []*ContractCoin
	err := m.conn.QueryRowsCtx(ctx, &resp, query)
	if err != nil {
		if err == sqlx.ErrNotFound {
			return []*ContractCoin{}, nil
		}
		return nil, err
	}
	return resp, nil
}

func (m *defaultContractCoinModel) Update(ctx context.Context, coin *ContractCoin) error {
	query := `UPDATE contract_coins SET symbol=?, contract_type=?, price_precision=?, quantity_precision=?, share_number=?, maker_fee=?, taker_fee=?, min_leverage=?, max_leverage=?, status=? WHERE id=?`
	_, err := m.conn.ExecCtx(ctx, query, coin.Symbol, coin.ContractType, coin.PricePrecision, coin.QuantityPrecision, coin.ShareNumber, coin.MakerFee, coin.TakerFee, coin.MinLeverage, coin.MaxLeverage, coin.Status, coin.Id)
	return err
}

func (m *defaultContractCoinModel) Delete(ctx context.Context, id int64) error {
	query := `DELETE FROM contract_coins WHERE id=?`
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

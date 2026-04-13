package model

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type defaultContractPositionModel struct {
	conn  sqlx.SqlConn
	table string
}

func NewContractPositionModel(conn sqlx.SqlConn) ContractPositionModel {
	return &defaultContractPositionModel{
		conn:  conn,
		table: "contract_positions",
	}
}

func (m *defaultContractPositionModel) Insert(ctx context.Context, position *ContractPosition) (int64, error) {
	query := `INSERT INTO contract_positions (member_id, contract_coin_id, direction, leverage, size, entry_price, margin, unrealized_pnl, liquidation_price) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`
	res, err := m.conn.ExecCtx(ctx, query, position.MemberId, position.ContractCoinId, position.Direction, position.Leverage, position.Size, position.EntryPrice, position.Margin, position.UnrealizedPnl, position.LiquidationPrice)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	position.Id = id
	return id, nil
}

func (m *defaultContractPositionModel) FindOne(ctx context.Context, id int64) (*ContractPosition, error) {
	query := `SELECT id, member_id, contract_coin_id, direction, leverage, size, entry_price, margin, unrealized_pnl, liquidation_price, created_at, updated_at FROM contract_positions WHERE id = ?`
	var resp ContractPosition
	err := m.conn.QueryRowCtx(ctx, &resp, query, id)
	if err != nil {
		if err == sqlx.ErrNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &resp, nil
}

func (m *defaultContractPositionModel) FindByMember(ctx context.Context, memberId int64, contractCoinId int64, direction int32) (*ContractPosition, error) {
	query := `SELECT id, member_id, contract_coin_id, direction, leverage, size, entry_price, margin, unrealized_pnl, liquidation_price, created_at, updated_at FROM contract_positions WHERE member_id = ? AND contract_coin_id = ? AND direction = ?`
	var resp ContractPosition
	err := m.conn.QueryRowCtx(ctx, &resp, query, memberId, contractCoinId, direction)
	if err != nil {
		if err == sqlx.ErrNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &resp, nil
}

func (m *defaultContractPositionModel) FindByMemberId(ctx context.Context, memberId int64) ([]*ContractPosition, error) {
	query := `SELECT id, member_id, contract_coin_id, direction, leverage, size, entry_price, margin, unrealized_pnl, liquidation_price, created_at, updated_at FROM contract_positions WHERE member_id = ?`
	var resp []*ContractPosition
	err := m.conn.QueryRowsCtx(ctx, &resp, query, memberId)
	if err != nil {
		if err == sqlx.ErrNotFound {
			return []*ContractPosition{}, nil
		}
		return nil, err
	}
	return resp, nil
}

func (m *defaultContractPositionModel) FindByContractCoinId(ctx context.Context, contractCoinId int64) ([]*ContractPosition, error) {
	query := `SELECT id, member_id, contract_coin_id, direction, leverage, size, entry_price, margin, unrealized_pnl, liquidation_price, created_at, updated_at FROM contract_positions WHERE contract_coin_id = ?`
	var resp []*ContractPosition
	err := m.conn.QueryRowsCtx(ctx, &resp, query, contractCoinId)
	if err != nil {
		if err == sqlx.ErrNotFound {
			return []*ContractPosition{}, nil
		}
		return nil, err
	}
	return resp, nil
}

func (m *defaultContractPositionModel) Update(ctx context.Context, position *ContractPosition) error {
	query := `UPDATE contract_positions SET size=?, entry_price=?, margin=?, unrealized_pnl=?, liquidation_price=? WHERE id=?`
	_, err := m.conn.ExecCtx(ctx, query, position.Size, position.EntryPrice, position.Margin, position.UnrealizedPnl, position.LiquidationPrice, position.Id)
	return err
}

func (m *defaultContractPositionModel) Upsert(ctx context.Context, position *ContractPosition) error {
	query := `INSERT INTO contract_positions (member_id, contract_coin_id, direction, leverage, size, entry_price, margin, unrealized_pnl, liquidation_price) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?) ON DUPLICATE KEY UPDATE size=VALUES(size), entry_price=VALUES(entry_price), margin=VALUES(margin), unrealized_pnl=VALUES(unrealized_pnl), liquidation_price=VALUES(liquidation_price)`
	_, err := m.conn.ExecCtx(ctx, query, position.MemberId, position.ContractCoinId, position.Direction, position.Leverage, position.Size, position.EntryPrice, position.Margin, position.UnrealizedPnl, position.LiquidationPrice)
	return err
}

func (m *defaultContractPositionModel) Delete(ctx context.Context, id int64) error {
	query := `DELETE FROM contract_positions WHERE id=?`
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

type defaultContractWalletModel struct {
	conn  sqlx.SqlConn
	table string
}

func NewContractWalletModel(conn sqlx.SqlConn) ContractWalletModel {
	return &defaultContractWalletModel{
		conn:  conn,
		table: "contract_wallets",
	}
}

func (m *defaultContractWalletModel) Insert(ctx context.Context, wallet *ContractWallet) (int64, error) {
	query := `INSERT INTO contract_wallets (member_id, unit, balance, frozen, main_balance) VALUES (?, ?, ?, ?, ?)`
	res, err := m.conn.ExecCtx(ctx, query, wallet.MemberId, wallet.Unit, wallet.Balance, wallet.Frozen, wallet.MainBalance)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	wallet.Id = id
	return id, nil
}

func (m *defaultContractWalletModel) FindOne(ctx context.Context, id int64) (*ContractWallet, error) {
	query := `SELECT id, member_id, unit, balance, frozen, main_balance, created_at, updated_at FROM contract_wallets WHERE id = ?`
	var resp ContractWallet
	err := m.conn.QueryRowCtx(ctx, &resp, query, id)
	if err != nil {
		if err == sqlx.ErrNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &resp, nil
}

func (m *defaultContractWalletModel) FindByMemberAndUnit(ctx context.Context, memberId int64, unit string) (*ContractWallet, error) {
	query := `SELECT id, member_id, unit, balance, frozen, main_balance, created_at, updated_at FROM contract_wallets WHERE member_id = ? AND unit = ?`
	var resp ContractWallet
	err := m.conn.QueryRowCtx(ctx, &resp, query, memberId, unit)
	if err != nil {
		if err == sqlx.ErrNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &resp, nil
}

func (m *defaultContractWalletModel) UpdateBalance(ctx context.Context, id int64, balance, frozen, mainBalance float64) error {
	query := `UPDATE contract_wallets SET balance=?, frozen=?, main_balance=? WHERE id=?`
	_, err := m.conn.ExecCtx(ctx, query, balance, frozen, mainBalance, id)
	return err
}

type defaultContractCoinModel struct {
	conn  sqlx.SqlConn
	table string
}

func NewContractCoinModel(conn sqlx.SqlConn) ContractCoinModel {
	return &defaultContractCoinModel{
		conn:  conn,
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

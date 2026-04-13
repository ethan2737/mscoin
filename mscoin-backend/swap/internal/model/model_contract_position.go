package model

import (
	"context"
	"time"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ContractPosition struct {
	Id              int64     `db:"id"`
	MemberId        int64     `db:"member_id"`
	ContractCoinId  int64     `db:"contract_coin_id"`
	Direction       int32     `db:"direction"`
	Leverage        int32     `db:"leverage"`
	Size            float64   `db:"size"`
	EntryPrice      float64   `db:"entry_price"`
	Margin          float64   `db:"margin"`
	UnrealizedPnl   float64   `db:"unrealized_pnl"`
	LiquidationPrice float64  `db:"liquidation_price"`
	CreatedAt       time.Time `db:"created_at"`
	UpdatedAt       time.Time `db:"updated_at"`
}

type ContractPositionModel interface {
	Insert(ctx context.Context, position *ContractPosition) (int64, error)
	FindOne(ctx context.Context, id int64) (*ContractPosition, error)
	FindByMember(ctx context.Context, memberId int64, contractCoinId int64, direction int32) (*ContractPosition, error)
	FindByMemberId(ctx context.Context, memberId int64) ([]*ContractPosition, error)
	Update(ctx context.Context, position *ContractPosition) error
	Upsert(ctx context.Context, position *ContractPosition) error
	Delete(ctx context.Context, id int64) error
}

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

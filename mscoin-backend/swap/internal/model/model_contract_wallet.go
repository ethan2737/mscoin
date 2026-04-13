package model

import (
	"context"
	"time"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ContractWallet struct {
	Id          int64     `db:"id"`
	MemberId    int64     `db:"member_id"`
	Unit        string    `db:"unit"`
	Balance     float64   `db:"balance"`
	Frozen      float64   `db:"frozen"`
	MainBalance float64   `db:"main_balance"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}

type ContractWalletModel interface {
	Insert(ctx context.Context, wallet *ContractWallet) (int64, error)
	FindOne(ctx context.Context, id int64) (*ContractWallet, error)
	FindByMemberId(ctx context.Context, memberId int64) ([]*ContractWallet, error)
	FindByMemberAndUnit(ctx context.Context, memberId int64, unit string) (*ContractWallet, error)
	UpdateBalance(ctx context.Context, id int64, balance float64, frozen float64, mainBalance float64) error
	Upsert(ctx context.Context, wallet *ContractWallet) error
	Delete(ctx context.Context, id int64) error
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

func (m *defaultContractWalletModel) FindByMemberId(ctx context.Context, memberId int64) ([]*ContractWallet, error) {
	query := `SELECT id, member_id, unit, balance, frozen, main_balance, created_at, updated_at FROM contract_wallets WHERE member_id = ?`
	var resp []*ContractWallet
	err := m.conn.QueryRowsCtx(ctx, &resp, query, memberId)
	if err != nil {
		if err == sqlx.ErrNotFound {
			return []*ContractWallet{}, nil
		}
		return nil, err
	}
	return resp, nil
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

func (m *defaultContractWalletModel) UpdateBalance(ctx context.Context, id int64, balance float64, frozen float64, mainBalance float64) error {
	query := `UPDATE contract_wallets SET balance=?, frozen=?, main_balance=? WHERE id=?`
	_, err := m.conn.ExecCtx(ctx, query, balance, frozen, mainBalance, id)
	return err
}

func (m *defaultContractWalletModel) Upsert(ctx context.Context, wallet *ContractWallet) error {
	query := `INSERT INTO contract_wallets (member_id, unit, balance, frozen, main_balance) VALUES (?, ?, ?, ?, ?) ON DUPLICATE KEY UPDATE balance=VALUES(balance), frozen=VALUES(frozen), main_balance=VALUES(main_balance)`
	_, err := m.conn.ExecCtx(ctx, query, wallet.MemberId, wallet.Unit, wallet.Balance, wallet.Frozen, wallet.MainBalance)
	return err
}

func (m *defaultContractWalletModel) Delete(ctx context.Context, id int64) error {
	query := `DELETE FROM contract_wallets WHERE id=?`
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

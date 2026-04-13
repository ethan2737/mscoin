package logic

import (
	"context"
	"errors"
	"strconv"
	"time"

	"mscoin-common/pages"
	"swap-api/internal/dao"
	"swap-api/internal/svc"
	"swap-api/internal/types"
)

type ContractWalletLogic struct {
	ctx            context.Context
	svcCtx         *svc.ServiceContext
	walletDao      dao.ContractWalletDao
	transactionDao dao.ContractTransactionDao
}

func NewContractWalletLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ContractWalletLogic {
	return &ContractWalletLogic{
		ctx:            ctx,
		svcCtx:         svcCtx,
		walletDao:      dao.NewContractWalletDao(svcCtx.ContractWalletModel),
		transactionDao: dao.NewContractTransactionDao(svcCtx.ContractTransactionModel),
	}
}

func (l *ContractWalletLogic) Info(req *types.ContractWalletReq) ([]*types.ContractWalletItem, error) {
	wallets, err := l.walletDao.GetByMemberId(l.ctx, req.MemberId)
	if err != nil {
		return nil, err
	}
	return buildWalletItems(wallets), nil
}

func (l *ContractWalletLogic) Transfer(req *types.ContractTransferReq) (*types.ContractTransferResp, error) {
	wallet, err := l.walletDao.GetByMemberAndUnit(l.ctx, req.MemberId, req.Unit)
	if err != nil {
		return nil, err
	}
	if wallet == nil {
		return nil, errors.New("合约钱包不存在")
	}

	transferType := resolveTransferType(req)
	switch transferType {
	case 1:
		newBalance := wallet.Balance + req.Amount
		if err := l.walletDao.UpdateBalance(l.ctx, wallet.Id, newBalance, wallet.Frozen, wallet.MainBalance); err != nil {
			return nil, err
		}
	case 2:
		if wallet.Balance < req.Amount {
			return nil, errors.New("余额不足")
		}
		newBalance := wallet.Balance - req.Amount
		if err := l.walletDao.UpdateBalance(l.ctx, wallet.Id, newBalance, wallet.Frozen, wallet.MainBalance); err != nil {
			return nil, err
		}
	default:
		return nil, errors.New("划转方向无效")
	}

	return &types.ContractTransferResp{Success: true}, nil
}

func (l *ContractWalletLogic) Transaction(req *types.ContractTransactionReq) (*pages.PageResult, error) {
	pageNo := req.PageNo
	if pageNo < 1 {
		pageNo = 1
	}
	pageSize := req.PageSize
	if pageSize < 1 {
		pageSize = 20
	}

	now := time.Now()
	startTime := now.AddDate(0, -1, 0)
	endTime := now
	if req.StartTime != "" {
		startTime = parseRequestTime(req.StartTime, startTime)
	}
	if req.EndTime != "" {
		endTime = parseRequestTime(req.EndTime, endTime)
	}

	txType := int32(-1)
	switch value := req.Type.(type) {
	case int:
		txType = int32(value)
	case int32:
		txType = value
	case int64:
		txType = int32(value)
	case float64:
		txType = int32(value)
	case string:
		if value != "" {
			if parsed, err := strconv.ParseInt(value, 10, 32); err == nil {
				txType = int32(parsed)
			}
		}
	}

	offset := int64((pageNo - 1) * pageSize)
	transactions, err := l.transactionDao.GetByMemberId(l.ctx, req.MemberId, startTime, endTime, txType, int64(pageSize), offset)
	if err != nil {
		return nil, err
	}

	return buildTransactionPage(transactions, int64(pageNo), int64(pageSize)), nil
}

func parseRequestTime(raw string, fallback time.Time) time.Time {
	if raw == "" {
		return fallback
	}
	if ts, err := strconv.ParseInt(raw, 10, 64); err == nil {
		if ts > 1_000_000_000_000 {
			return time.UnixMilli(ts)
		}
		return time.Unix(ts, 0)
	}
	if parsed, err := time.ParseInLocation("2006-01-02 15:04:05", raw, time.Local); err == nil {
		return parsed
	}
	return fallback
}

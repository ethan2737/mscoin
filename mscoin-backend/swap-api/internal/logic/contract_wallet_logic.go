package logic

import (
	"context"
	"encoding/json"
	"errors"
	"strconv"
	"time"

	"mscoin-common/pages"
	"swap-api/internal/dao"
	"swap-api/internal/model"
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
	// 如果钱包不存在，自动初始化 USDT 合约钱包
	if len(wallets) == 0 {
		initWallet := &model.ContractWallet{
			MemberId:    req.MemberId,
			Unit:        "USDT",
			Balance:     0,
			Frozen:      0,
			MainBalance: 0,
		}
		if err := l.walletDao.Upsert(l.ctx, initWallet); err != nil {
			return nil, err
		}
		wallets = append(wallets, initWallet)
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

	txType := parseTransactionType(req.Type)

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

func parseTransactionType(raw json.RawMessage) int32 {
	if len(raw) == 0 || string(raw) == `""` || string(raw) == "null" {
		return -1
	}

	var number int32
	if err := json.Unmarshal(raw, &number); err == nil {
		return number
	}

	var text string
	if err := json.Unmarshal(raw, &text); err == nil && text != "" {
		if parsed, err := strconv.ParseInt(text, 10, 32); err == nil {
			return int32(parsed)
		}
	}

	return -1
}

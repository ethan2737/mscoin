package logic

import (
	"context"
	"swap/internal/dao"
	"swap/internal/svc"
	"swap/types/swap"
	"time"
)

type ContractWalletLogic struct {
	ctx          context.Context
	svc          *svc.ServiceContext
	walletDao    dao.ContractWalletDao
	transactionDao dao.ContractTransactionDao
}

func NewContractWalletLogic(ctx context.Context, svc *svc.ServiceContext) *ContractWalletLogic {
	return &ContractWalletLogic{
		ctx:            ctx,
		svc:            svc,
		walletDao:      dao.NewContractWalletDao(svc.ContractWalletModel),
		transactionDao: dao.NewContractTransactionDao(svc.ContractTransactionModel),
	}
}

func (l *ContractWalletLogic) GetWallet(req *swap.GetWalletRequest) (*swap.GetWalletResponse, error) {
	wallet, err := l.walletDao.GetByMemberAndUnit(l.ctx, req.MemberId, req.Unit)
	if err != nil {
		return nil, err
	}
	if wallet == nil {
		// 钱包不存在，返回空响应
		return &swap.GetWalletResponse{}, nil
	}
	return &swap.GetWalletResponse{
		Wallet: &swap.WalletInfo{
			Id:          wallet.Id,
			MemberId:    wallet.MemberId,
			Unit:        wallet.Unit,
			Balance:     wallet.Balance,
			Frozen:      wallet.Frozen,
			MainBalance: wallet.MainBalance,
		},
	}, nil
}

func (l *ContractWalletLogic) Transfer(req *swap.TransferRequest) (*swap.TransferResponse, error) {
	// 1. 获取钱包
	wallet, err := l.walletDao.GetByMemberAndUnit(l.ctx, req.MemberId, req.Unit)
	if err != nil {
		return nil, err
	}

	// 2. 根据方向处理划转
	if req.Direction == 1 {
		// 从主账户划入合约账户
		if err := l.transferIn(req, wallet); err != nil {
			return nil, err
		}
	} else {
		// 从合约账户划出到主账户
		if err := l.transferOut(req, wallet); err != nil {
			return nil, err
		}
	}

	return &swap.TransferResponse{Success: true}, nil
}

func (l *ContractWalletLogic) transferIn(req *swap.TransferRequest, wallet interface{}) error {
	// TODO: 实现从主账户划入逻辑
	// 需要调用 ucenter RPC 扣减主账户余额
	return nil
}

func (l *ContractWalletLogic) transferOut(req *swap.TransferRequest, wallet interface{}) error {
	// 检查余额是否足够
	// TODO: 实现划出逻辑
	return nil
}

func (l *ContractWalletLogic) recordTransaction(memberId int64, unit string, txType int32, amount float64, balance float64, orderId int64, remark string) error {
	tx := &struct {
		MemberId       int64
		Unit           string
		Type           int32
		Amount         float64
		Balance        float64
		RelatedOrderId int64
		Remark         string
	}{
		MemberId:       memberId,
		Unit:           unit,
		Type:           txType,
		Amount:         amount,
		Balance:        balance,
		RelatedOrderId: orderId,
		Remark:         remark,
	}
	_ = tx
	// TODO: 实现账单记录
	return nil
}

func (l *ContractWalletLogic) GetTransactions(req *swap.GetTransactionsRequest) (*swap.GetTransactionsResponse, error) {
	startTime := time.Unix(req.StartTime, 0)
	endTime := time.Unix(req.EndTime, 0)

	// 默认分页参数：limit=10, offset=0
	limit := int64(10)
	offset := int64(0)

	txs, err := l.transactionDao.GetByMemberId(l.ctx, req.MemberId, startTime, endTime, req.Type, limit, offset)
	if err != nil {
		return nil, err
	}

	list := make([]*swap.TransactionInfo, 0, len(txs))
	for _, tx := range txs {
		list = append(list, &swap.TransactionInfo{
			Id:             tx.Id,
			MemberId:       tx.MemberId,
			Unit:           tx.Unit,
			Type:           tx.Type,
			Amount:         tx.Amount,
			Balance:        tx.Balance,
			RelatedOrderId: tx.RelatedOrderId,
			Remark:         tx.Remark,
		})
	}
	return &swap.GetTransactionsResponse{List: list}, nil
}

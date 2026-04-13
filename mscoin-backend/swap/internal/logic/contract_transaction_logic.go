package logic

import (
	"context"
	"swap/internal/dao"
	"swap/internal/svc"
	"swap/types/swap"
	"time"
)

type ContractTransactionLogic struct {
	ctx            context.Context
	svc            *svc.ServiceContext
	transactionDao dao.ContractTransactionDao
}

func NewContractTransactionLogic(ctx context.Context, svc *svc.ServiceContext) *ContractTransactionLogic {
	return &ContractTransactionLogic{
		ctx:            ctx,
		svc:            svc,
		transactionDao: dao.NewContractTransactionDao(svc.ContractTransactionModel),
	}
}

func (l *ContractTransactionLogic) GetTransactions(req *swap.GetTransactionsRequest) (*swap.GetTransactionsResponse, error) {
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

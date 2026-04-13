package repo

import (
	"context"
	"jobcenter/internal/model"
)

type KlineRepo interface {
	SaveBatch(ctx context.Context, data []*model.Kline, symbol, period string) error
	DeleteGtTime(background context.Context, time int64, symbol string, period string) error
	SaveBatchSwap(ctx context.Context, data []*model.SwapKline, symbol, period string) error
	DeleteGtTimeSwap(background context.Context, time int64, symbol string, period string) error
}

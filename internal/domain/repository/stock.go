package repository

import (
	"context"
	"live-trading/internal/domain/entity"
)

type StockRepo interface {
	WatchPickStock(ctx context.Context, codes entity.StockCodes, rec chan []entity.PickStock) error
	StartWatch()
	StopWatch()
}

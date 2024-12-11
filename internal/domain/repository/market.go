package repository

import (
	"context"
	"live-trading/internal/domain/entity"
)

type MarketRepo interface {
	WatchMarket(ctx context.Context, code entity.StockCodes, res chan<- []entity.PickStock) error
	ListMarket(ctx context.Context) (entity.Markets, error)
}

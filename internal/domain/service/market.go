package service

import (
	"context"
	"fmt"
	"live-trading/internal/domain/entity"
	"live-trading/internal/domain/repository"
	"live-trading/internal/infrastructure/dongfang"
	"live-trading/internal/infrastructure/xueqiu"
)

type IMarket interface {
	WatchMarket() error
	ListMarket() (entity.Markets, error)
}

type Market struct {
	ctx                  context.Context
	cancel               context.CancelFunc
	marketRepo           repository.MarketRepo
	xueqiuMarketRepoImpl repository.MarketRepo
}

func NewMarket() *Market {

	return &Market{
		ctx:                  context.Background(),
		marketRepo:           dongfang.NewMarketRepoImpl(),
		xueqiuMarketRepoImpl: xueqiu.NewMarketRepoImpl(),
	}
}

func NewMarketWithContext(ctx context.Context) *Market {
	market := &Market{
		marketRepo: dongfang.NewMarketRepoImpl(),
	}
	market.ctx, market.cancel = context.WithCancel(ctx)
	return market
}

func (m *Market) WatchMarket() error {
	rec := make(chan []entity.PickStock, 100)
	go func() {
		err := m.marketRepo.WatchMarket(m.ctx, entity.MarketCode, rec)
		if err != nil {
			fmt.Println(err)
		}
	}()

	for stocks := range rec {
		copyStocks := stocks
		entity.RefreshGlobalMarketStock(copyStocks)
	}

	return nil
}

func (m *Market) ListMarket() (entity.Markets, error) {
	return m.xueqiuMarketRepoImpl.ListMarket(m.ctx)
}

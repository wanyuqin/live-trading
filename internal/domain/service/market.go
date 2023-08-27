package service

import (
	"context"
	"fmt"
	"live-trading/internal/domain/entity"
	"live-trading/internal/domain/repository"
	"live-trading/internal/infrastructure/dongfang"
)

type IMarket interface {
	WatchMarket() error
}

type Market struct {
	ctx        context.Context
	cancel     context.CancelFunc
	marketRepo repository.MarketRepo
}

func NewMarket() *Market {

	return &Market{
		ctx:        context.Background(),
		marketRepo: dongfang.NewDongFangMarketRepoImpl(),
	}
}

func NewMarketWithContext(ctx context.Context) *Market {
	market := &Market{
		marketRepo: dongfang.NewDongFangMarketRepoImpl(),
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
			return
		}
	}()

	for stocks := range rec {
		copyStocks := stocks
		entity.RefreshGlobalMarketStock(copyStocks)
	}

	return nil
}

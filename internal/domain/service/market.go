package service

import (
	"context"
	"fmt"
	"live-trading/internal/domain/entity"
	"live-trading/internal/domain/repository"
	"live-trading/internal/infrastructure/dongfang"
)

type IMarket interface {
	ListMarket(ctx context.Context) (entity.Markets, error)
}

type Market struct {
	marketRepo repository.MarketRepo
}

func NewMarket() *Market {
	return &Market{
		marketRepo: dongfang.NewDongFangMarketRepoImpl(),
	}
}

func (m Market) ListMarket(ctx context.Context) (entity.Markets, error) {
	res := make(chan []byte)
	go m.marketRepo.ListMarket(res)
	for re := range res {
		fmt.Println(re)
	}
	return nil, nil
}

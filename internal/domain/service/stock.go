package service

import (
	"context"
	"errors"
	"fmt"
	"live-trading/internal/configs"
	"live-trading/internal/domain/entity"
	"live-trading/internal/domain/repository"
	"live-trading/internal/infrastructure/dongfang"
	"regexp"
)

type IStock interface {
	WatchPickStocks() error
	GetPickStocks(ctx context.Context) entity.StockCodes
	AddPickStockCode(ctx context.Context, code string) error
	RestartWatchPickStocks(ctx context.Context) error
	DeletePickStockCode(code string) error
}

var stockCodeRegex = regexp.MustCompile("^[0-9a-zA-Z]{6}$")

type Stock struct {
	ctx       context.Context
	cancel    context.CancelFunc
	stockRepo repository.StockRepo
}

func NewStock() *Stock {
	return &Stock{
		ctx:       context.Background(),
		stockRepo: dongfang.NewDongFangStockRepoImpl(),
	}
}

func NewStockWithContext(ctx context.Context) *Stock {
	ctx, cancel := context.WithCancel(ctx)
	return &Stock{
		ctx:       ctx,
		cancel:    cancel,
		stockRepo: dongfang.NewDongFangStockRepoImpl(),
	}
}

func (s *Stock) WatchPickStocks() error {
	pickStocks := s.GetPickStocks(s.ctx)
	rec := make(chan []entity.PickStock, 100)
	go func() {
		err := s.stockRepo.WatchPickStock(s.ctx, pickStocks, rec)
		if err != nil {
			fmt.Println(err)
			return
		}
	}()

	for stocks := range rec {
		copyStocks := stocks
		entity.RefreshGlobalPickStock(copyStocks)
	}

	return nil

}

func (s *Stock) GetPickStocks(ctx context.Context) entity.StockCodes {
	var pickStock entity.StockCodes
	config := configs.GetConfig()
	if config == nil {
		return pickStock
	}

	return entity.NewStockCodes(config.WatchList.Stock)
}

func (s *Stock) AddPickStockCode(ctx context.Context, code string) error {
	if !stockCodeRegex.MatchString(code) {
		return errors.New("stock code invalid")
	}
	err := configs.GetConfig().AddStockCode(code)
	if err != nil {
		return err
	}

	return nil
}

func (s *Stock) RestartWatchPickStocks(ctx context.Context) error {
	s.cancel()
	s.ctx, s.cancel = context.WithCancel(ctx)
	s.WatchPickStocks()
	return nil
}

func (s *Stock) DeletePickStockCode(code string) error {
	return configs.GetConfig().DeleteStockCode(code)

}

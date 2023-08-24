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
	WatchPickStocks(ctx context.Context) error
	GetPickStocks(ctx context.Context) entity.StockCodes
	AddPickStockCode(ctx context.Context, code string) error
	RestartWatchPickStocks() error
}

var stockCodeRegex = regexp.MustCompile("^[0-9a-zA-Z]{6}$")

type Stock struct {
	stockRepo repository.StockRepo
}

func NewStock() *Stock {
	return &Stock{
		stockRepo: dongfang.NewDongFangStockRepoImpl(),
	}
}

func (s *Stock) WatchPickStocks(ctx context.Context) error {
	pickStocks := s.GetPickStocks(ctx)
	entity.NewGlobalPickStock()
	rec := make(chan []entity.PickStock, 100)
	go func() {
		err := s.stockRepo.WatchPickStock(ctx, pickStocks, rec)
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

func (s *Stock) RestartWatchPickStocks() error {
	s.stockRepo.StopWatch()

	return nil
}

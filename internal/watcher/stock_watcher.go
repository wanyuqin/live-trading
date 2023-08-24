package watcher

import (
	"context"
	"live-trading/internal/domain/service"
)

type StockWatcher struct {
	Ctx context.Context

	Codes        []string
	StockService service.IStock
}

func NewStockWatcher(ctx context.Context, codes []string) *StockWatcher {
	return &StockWatcher{
		Ctx:          ctx,
		Codes:        codes,
		StockService: service.NewStock(),
	}
}

func (watcher *StockWatcher) Start() error {
	err := watcher.StockService.WatchPickStocks(watcher.Ctx)
	return err
}

func (watcher *StockWatcher) Reload() {
	watcher.StockService.RestartWatchPickStocks()
}

func (watcher *StockWatcher) Stop() {

}

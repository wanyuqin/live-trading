package index

import "fmt"

func (m *Model) deleteStockCode() {
	code := m.stock.GetRowCode(m.stock.Table.GetHighlightedRowIndex())
	err := m.stockService.DeletePickStockCode(code)
	if err != nil {
		fmt.Println(err)
	}
	go m.stockService.RestartWatchPickStocks(m.ctx)
}

func (m *Model) addStockCode(code string) {
	err := m.stockService.AddPickStockCode(m.ctx, code)
	if err != nil {
		fmt.Println(err)
	}

	go m.stockService.RestartWatchPickStocks(m.ctx)
}

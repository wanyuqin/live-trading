package stock

import (
	"fmt"
	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/lipgloss"
	"live-trading/internal/domain/entity"
)

func defaultPickStockTableColumn() []table.Column {
	columns := []table.Column{
		{Title: "代码", Width: 10},
		{Title: "名称", Width: 10},
		{Title: "最新价", Width: 10},
		{Title: "涨跌额", Width: 10},
	}
	return columns
}

func initPickStocksTable() []table.Row {
	pickStocks := entity.GetGlobalPickStock()
	return transformTableRows(pickStocks)
}

func transformTableRows(pickStocks []entity.PickStock) []table.Row {
	rows := make([]table.Row, 0, len(pickStocks))
	for i := range pickStocks {
		pickStock := pickStocks[i]
		row := table.Row{
			pickStock.Code,
			pickStock.Name,
			fmt.Sprintf("%.2f", pickStock.Now),
			fmt.Sprintf("%.2f", pickStock.Diff),
		}
		rows = append(rows, row)
	}

	return rows
}

func defaultTableStyle() table.Styles {
	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(false).
		Align(lipgloss.Center)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Bold(false)

	return s
}

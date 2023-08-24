package stock

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/lipgloss"
	"live-trading/internal/domain/entity"
)

type Model struct {
	List  list.Model
	Table table.Model
	Keys  *tableKeyMap
}

func NewStockModel() Model {
	t := table.New(
		table.WithColumns(defaultPickStockTableColumn()),
		table.WithRows(initPickStocksTable()),
		table.WithFocused(true),
		table.WithHeight(7),
		table.WithWidth(100))
	t.SetStyles(defaultTableStyle())
	return Model{
		Table: t,
		Keys:  newTableKeyMap(),
	}

}

func (m *Model) RefreshTable() {
	pickStocks := entity.GetGlobalPickStock()
	rows := transformTableRows(pickStocks)
	t := table.New(
		table.WithColumns(defaultPickStockTableColumn()),
		table.WithRows(rows),
		table.WithHeight(7),
		table.WithWidth(100),
	)

	style := defaultTableStyle()
	if len(rows) > 0 {
		style.Cell.Align(lipgloss.Center)
	}
	t.SetStyles(style)
	m.Table = t

	return

}

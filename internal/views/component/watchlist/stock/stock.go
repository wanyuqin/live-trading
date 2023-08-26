package stock

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/evertras/bubble-table/table"
	"live-trading/internal/domain/entity"
)

type Model struct {
	List  list.Model
	Table table.Model
	Keys  *tableKeyMap
}

func NewStockModel() Model {

	t := table.New(defaultPickStockTableColumn()).
		WithRows(initPickStocksTable()).
		Focused(true).
		WithKeyMap(initKeyMap()).
		Filtered(true).
		Border(table.Border{
			Left:   "",
			Right:  "",
			Top:    "",
			Bottom: "",
		})

	return Model{
		Table: t,
		Keys:  newTableKeyMap(),
	}

}

func (m *Model) RefreshTable() {
	pickStocks := entity.GetGlobalPickStock()
	rows := transformTableRows(pickStocks)
	m.Table = m.Table.WithRows(rows)
	return

}

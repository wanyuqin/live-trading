package stock

import (
	"fmt"
	"github.com/charmbracelet/lipgloss"
	"github.com/evertras/bubble-table/table"
	"live-trading/internal/domain/entity"
	"live-trading/internal/views/component"
)

const (
	defaultTableHeight = 10
	defaultTableWidth  = 100
)

const (
	columnKeyCode          = "code"
	columnKeyName          = "name"
	columnKeyNow           = "trade"
	columnKeyDiff          = "diff"
	columnKeyChangePercent = "changePercent"
)

func defaultPickStockTableColumn() []table.Column {
	columns := []table.Column{
		table.NewColumn(columnKeyCode, "代码", 10).WithFiltered(true),
		table.NewColumn(columnKeyName, "名称", 10),
		table.NewColumn(columnKeyNow, "最新价", 10),
		table.NewColumn(columnKeyDiff, "涨跌额", 10),
		table.NewColumn(columnKeyChangePercent, "涨跌幅", 10),
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
		row := makeRow(pickStock.Code, pickStock.Name, pickStock.Trade, pickStock.Diff, pickStock.ChangePercent)
		rows = append(rows, row)
	}

	return rows
}

func makeRow(code string, name string, trade float64, diff float64, changePercent float64) table.Row {
	var (
		tradeColumn         interface{} = trade
		diffColumn          interface{} = diff
		changePercentColumn interface{} = changePercent
	)

	if diff < 0 {
		diffColumn = table.NewStyledCell(fmt.Sprintf("%.2f", diff), lipgloss.NewStyle().Foreground(lipgloss.Color(component.ColorGreen)))
		tradeColumn = table.NewStyledCell(fmt.Sprintf("%.2f", trade), lipgloss.NewStyle().Foreground(lipgloss.Color(component.ColorGreen)))
		changePercentColumn = table.NewStyledCell(fmt.Sprintf("%.2f%%", changePercent), lipgloss.NewStyle().Foreground(lipgloss.Color(component.ColorGreen)))
	}

	if diff > 0 {
		diffColumn = table.NewStyledCell(fmt.Sprintf("%.2f", diff), lipgloss.NewStyle().Foreground(lipgloss.Color(component.ColorFire)))
		tradeColumn = table.NewStyledCell(fmt.Sprintf("%.2f", trade), lipgloss.NewStyle().Foreground(lipgloss.Color(component.ColorFire)))
		changePercentColumn = table.NewStyledCell(fmt.Sprintf("%.2f%s", changePercent, "%"), lipgloss.NewStyle().Foreground(lipgloss.Color(component.ColorFire)))

	}

	return table.NewRow(table.RowData{
		columnKeyCode:          code,
		columnKeyName:          name,
		columnKeyNow:           tradeColumn,
		columnKeyDiff:          diffColumn,
		columnKeyChangePercent: changePercentColumn,
	})
}

func initKeyMap() table.KeyMap {
	keys := table.DefaultKeyMap()
	return keys
}

func (m *Model) GetRowCode() string {
	rows := m.Table.GetVisibleRows()
	if code, ok := rows[m.Table.GetHighlightedRowIndex()].Data[columnKeyCode].(string); ok {
		return code
	}
	return ""
}

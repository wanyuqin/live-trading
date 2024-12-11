package stock

import (
	"fmt"
	"live-trading/internal/domain/entity"

	"github.com/charmbracelet/lipgloss"
	"github.com/evertras/bubble-table/table"
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

	colorNormal   = "#fa0"
	colorFire     = "#f64"
	colorElectric = "#ff0"
	colorWater    = "#44f"
	colorPlant    = "#8b8"
)

func defaultPickStockTableColumn() []table.Column {
	columns := []table.Column{
		table.NewColumn(columnKeyCode, "代码", 15).WithFiltered(true),
		table.NewColumn(columnKeyName, "名称", 20),
		table.NewColumn(columnKeyNow, "最新价", 15),
		table.NewColumn(columnKeyDiff, "涨跌额", 15),
		table.NewColumn(columnKeyChangePercent, "涨跌幅", 15),
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
		diffColumn = table.NewStyledCell(fmt.Sprintf("%.2f", diff), lipgloss.NewStyle().Foreground(lipgloss.Color(colorWater)))
		tradeColumn = table.NewStyledCell(fmt.Sprintf("%.2f", trade), lipgloss.NewStyle().Foreground(lipgloss.Color(colorWater)))
		changePercentColumn = table.NewStyledCell(fmt.Sprintf("%.2f%%", changePercent), lipgloss.NewStyle().Foreground(lipgloss.Color(colorWater)))
	}

	if diff > 0 {
		diffColumn = table.NewStyledCell(fmt.Sprintf("%.2f", diff), lipgloss.NewStyle().Foreground(lipgloss.Color(colorFire)))
		tradeColumn = table.NewStyledCell(fmt.Sprintf("%.2f", trade), lipgloss.NewStyle().Foreground(lipgloss.Color(colorFire)))
		changePercentColumn = table.NewStyledCell(fmt.Sprintf("%.2f%s", changePercent, "%"), lipgloss.NewStyle().Foreground(lipgloss.Color(colorFire)))

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

func (m *Model) GetRowCode(index int) string {
	rows := m.Table.GetVisibleRows()
	if code, ok := rows[index].Data[columnKeyCode].(string); ok {
		return code
	}

	return ""
}

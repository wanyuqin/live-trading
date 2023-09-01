package fund

import (
	"fmt"
	"github.com/charmbracelet/lipgloss"
	"github.com/evertras/bubble-table/table"
	"live-trading/internal/domain/entity"
	"live-trading/internal/views/component"
	"strconv"
)

const (
	columnKeyCode             = "code"
	columnKeyName             = "name"
	columnKeyFNAV             = "fnav"
	columnKeyTRF              = "trf"
	columnKeyDailyInc         = "dailyInc"
	columnKeyDailyIncRate     = "dailyIncRate"
	columnKeyInceptionIncRate = "dailyInceptionIncRate"
	columnKeyType             = "dailyType"
	columnKeyDataTime         = "dailyDataTime"
)

func defaultFundTableColumn() []table.Column {
	columns := []table.Column{
		table.NewColumn(columnKeyCode, "代码", 10).WithFiltered(true).WithStyle(lipgloss.NewStyle().Align(lipgloss.Center)),
		table.NewColumn(columnKeyName, "名称", 30).WithStyle(lipgloss.NewStyle().Align(lipgloss.Left)),
		table.NewColumn(columnKeyDataTime, "日期", 15).WithStyle(lipgloss.NewStyle().Align(lipgloss.Center)),
		table.NewColumn(columnKeyFNAV, "单位净值", 10).WithStyle(lipgloss.NewStyle().Align(lipgloss.Left)),
		table.NewColumn(columnKeyTRF, "累计净值", 10).WithStyle(lipgloss.NewStyle().Align(lipgloss.Left)),
		table.NewColumn(columnKeyDailyInc, "日增长值", 10).WithStyle(lipgloss.NewStyle().Align(lipgloss.Center)),
		table.NewColumn(columnKeyDailyIncRate, "日增长率", 10).WithStyle(lipgloss.NewStyle().Align(lipgloss.Center)),
		table.NewColumn(columnKeyInceptionIncRate, "成立以来", 10).WithStyle(lipgloss.NewStyle().Align(lipgloss.Left)),
		table.NewColumn(columnKeyType, "类型", 20).WithStyle(lipgloss.NewStyle().Align(lipgloss.Center)),
	}
	return columns
}

func transformTableRows(list entity.FundList) []table.Row {
	rows := make([]table.Row, 0, len(list))
	for i := range list {
		fund := list[i]
		row := makeRow(fund)

		rows = append(rows, row)

	}
	return rows
}

func makeRow(found entity.Fund) table.Row {

	var (
		dailyInc         interface{} = found.DailyInc
		dailyIncRate     interface{} = found.DailyIncRate
		inceptionIncRate interface{} = found.InceptionIncRate
	)

	dailyIncf, err := strconv.ParseFloat(found.DailyInc, 64)
	if err == nil {
		if dailyIncf < 0 {
			dailyInc = table.NewStyledCell(found.DailyInc, lipgloss.NewStyle().Foreground(lipgloss.Color(component.ColorGreen)))
			dailyIncRate = table.NewStyledCell(fmt.Sprintf("%s%%", found.DailyInc), lipgloss.NewStyle().Foreground(lipgloss.Color(component.ColorGreen)))
		}

		if dailyIncf > 0 {
			dailyInc = table.NewStyledCell(found.DailyInc, lipgloss.NewStyle().Foreground(lipgloss.Color(component.ColorFire)))
			dailyIncRate = table.NewStyledCell(fmt.Sprintf("%s%%", found.DailyInc), lipgloss.NewStyle().Foreground(lipgloss.Color(component.ColorFire)))
		}
	}

	inceptionIncRatef, err := strconv.ParseFloat(found.InceptionIncRate, 64)
	if err == nil {
		if inceptionIncRatef < 0 {
			inceptionIncRate = table.NewStyledCell(fmt.Sprintf("%s%%", found.InceptionIncRate), lipgloss.NewStyle().Foreground(lipgloss.Color(component.ColorGreen)))
		}

		if inceptionIncRatef > 0 {
			inceptionIncRate = table.NewStyledCell(fmt.Sprintf("%s%%", found.InceptionIncRate), lipgloss.NewStyle().Foreground(lipgloss.Color(component.ColorFire)))
		}
	}

	return table.NewRow(table.RowData{
		columnKeyCode:             found.Code,
		columnKeyName:             found.Name,
		columnKeyDataTime:         found.DataTime,
		columnKeyFNAV:             found.FNAV,
		columnKeyTRF:              found.TRF,
		columnKeyDailyInc:         dailyInc,
		columnKeyDailyIncRate:     dailyIncRate,
		columnKeyInceptionIncRate: inceptionIncRate,
		columnKeyType:             found.Type,
	})
}

func initFundTable() []table.Row {
	return []table.Row{}
}

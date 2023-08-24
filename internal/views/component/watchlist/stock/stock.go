package stock

import (
	"fmt"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/lipgloss"
	"live-trading/internal/domain/service"
)

var (
	appStyle = lipgloss.NewStyle().Padding(1, 2)

	titleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFFDF5")).
			Background(lipgloss.Color("#25A065")).
			Padding(0, 1)

	statusMessageStyle = lipgloss.NewStyle().
				Foreground(lipgloss.AdaptiveColor{Light: "#04B575", Dark: "#04B575"}).
				Render
)

type Model struct {
	List         list.Model
	Table        table.Model
	Keys         *tableKeyMap
	DelegateKeys *delegateKeyMap
	StockService service.IStock
}

type PickStock struct {
	DataId int64   `json:"data_id"` // 每次的ID是一致的
	Name   string  `json:"name"`    // f14
	Code   string  `json:"code"`    // code f12
	Now    float64 `json:"now"`     // 当前价f31
	Diff   float64 `json:"diff"`    // 差值 f4
}

func (s PickStock) FilterValue() string { return s.Code }

func (s PickStock) Title() string { return s.Code }

func (s PickStock) Description() string {
	return fmt.Sprintf("%s              %.2f          %.2f", s.Name, s.Now, s.Diff)
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
		Table:        t,
		Keys:         newTableKeyMap(),
		StockService: service.NewStock(),
	}

}

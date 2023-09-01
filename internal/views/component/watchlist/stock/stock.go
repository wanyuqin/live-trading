package stock

import (
	"context"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/lipgloss"
	"github.com/evertras/bubble-table/table"
	"live-trading/internal/domain/entity"
	"live-trading/internal/domain/service"
)

type Model struct {
	ctx          context.Context
	List         list.Model
	Table        table.Model
	Keys         *tableKeyMap
	stockService service.IStock
}

func NewStockModel(ctx context.Context) *Model {
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

	return &Model{
		ctx:          ctx,
		Table:        t,
		Keys:         newTableKeyMap(),
		stockService: service.NewStockWithContext(ctx),
	}

}

func (m *Model) RefreshTable() {
	pickStocks := entity.GetGlobalPickStock()
	rows := transformTableRows(pickStocks)
	m.Table = m.Table.WithRows(rows)
	return
}

func (m *Model) View() string {
	return lipgloss.NewStyle().MarginLeft(1).Render(m.Table.View())

}

func (m *Model) AddItem(ctx context.Context, code string) error {
	err := m.stockService.AddPickStockCode(ctx, code)
	if err != nil {

		return err
	}

	go m.stockService.RestartWatchPickStocks(ctx)
	return nil
}

func (m *Model) DeleteItem(ctx context.Context, code string) error {
	return m.stockService.DeletePickStockCode(ctx, code)
}

func (m *Model) Detail(ctx context.Context, code string) (string, error) {
	return "", nil
}

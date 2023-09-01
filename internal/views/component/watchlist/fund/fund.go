package fund

import (
	"context"
	"fmt"
	"github.com/charmbracelet/lipgloss"
	"github.com/evertras/bubble-table/table"
	"live-trading/internal/domain/service"
)

type Model struct {
	ctx         context.Context
	Table       table.Model
	fundService service.IFund
}

func NewFundModel(ctx context.Context) *Model {
	t := table.New(defaultFundTableColumn()).
		WithRows(initFundTable()).
		WithKeyMap(table.DefaultKeyMap()).
		Focused(true).
		Filtered(true).
		Border(table.Border{
			Left:   "",
			Right:  "",
			Top:    "",
			Bottom: "",
		})

	return &Model{
		ctx:         ctx,
		Table:       t,
		fundService: service.NewFund(),
	}
}

func (m *Model) RefreshTable() {
	fundList, err := m.fundService.ListFundsInfo(context.Background())
	if err != nil {
		fmt.Println(err)
	}
	rows := transformTableRows(fundList)
	m.Table = m.Table.WithRows(rows)
	return
}

func (m *Model) View() string {
	return lipgloss.NewStyle().MarginLeft(1).Render(m.Table.View())
}

func (m *Model) AddItem(ctx context.Context, code string) error {
	err := m.fundService.AddFundCode(ctx, code)
	if err != nil {
		return err
	}
	return nil
}

func (m *Model) DeleteItem(ctx context.Context, code string) error {
	return m.fundService.DeleteFundCode(ctx, code)
}

func (m *Model) GetRowCode() string {
	rows := m.Table.GetVisibleRows()
	if code, ok := rows[m.Table.GetHighlightedRowIndex()].Data[columnKeyCode].(string); ok {
		return code
	}
	return ""
}

func (m *Model) Detail(ctx context.Context, code string) (string, error) {
	fundDetail, err := m.fundService.GetFundDetail(ctx, code)
	if err != nil {
		return "", err
	}
	return fundDetail.ParseTemplate()
}

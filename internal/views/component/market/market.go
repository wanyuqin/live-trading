package market

import (
	"context"
	"fmt"
	"github.com/charmbracelet/lipgloss"
	"live-trading/internal/domain/entity"
	"live-trading/internal/views/component"
)

type Model struct {
	ctx context.Context

	market string
}

var (
	subtle    = lipgloss.AdaptiveColor{Light: "#D9DCCF", Dark: "#383838"}
	highlight = lipgloss.AdaptiveColor{Light: "#874BFD", Dark: "#7D56F4"}
	special   = lipgloss.AdaptiveColor{Light: "#43BF6D", Dark: "#73F59F"}

	listStyle = lipgloss.NewStyle().
			Border(lipgloss.Border{
			Left:  "|",
			Right: "|",
		}, false, true, false, false).
		BorderForeground(subtle).
		Align(lipgloss.Center).
		MarginRight(1).
		Height(2).
		Width(columnWidth + 1)

	listHeader = lipgloss.NewStyle().
		//BorderStyle(lipgloss.NormalBorder()).
		BorderBottom(true).
		Align(lipgloss.Center).
		BorderForeground(subtle).
		Render

	listItem = lipgloss.NewStyle().Align(lipgloss.Center).Render

	downStyle = lipgloss.NewStyle().Foreground(lipgloss.Color(component.ColorGreen))
	upStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color(component.ColorRed))
)

const (
	columnWidth = 30
)

func NewModel(ctx context.Context) *Model {
	return &Model{
		ctx: ctx,
	}
}

func GetMarket() string {
	lists := lipgloss.JoinHorizontal(lipgloss.Top,
		listStyle.Copy().Render(
			lipgloss.JoinVertical(lipgloss.Left,
				listHeader("上证指数"),
				listItem("")),
		),

		listStyle.Copy().Render(
			lipgloss.JoinVertical(lipgloss.Left,
				listHeader("深证成指"),
				listItem("")),
		),

		listStyle.Copy().Render(
			lipgloss.JoinVertical(lipgloss.Left,
				listHeader("创业板指"),
				listItem("")),
		),
	)
	return lipgloss.JoinHorizontal(lipgloss.Top, lists)
}

func (m *Model) RefreshTable() {
	marketStock := entity.GetGlobalMarketStock()
	m.market = transformMarket(marketStock)
}

func (m *Model) View() string {
	return m.market
}

func transformMarket(markets []entity.PickStock) string {

	marketRenders := make([]string, 0, len(markets))
	for i := range markets {
		market := markets[i]
		changePercent := upStyle.Render(fmt.Sprintf("%.2f%%", market.ChangePercent))
		trade := upStyle.Render(fmt.Sprintf("%.2f", market.Trade))
		if market.ChangePercent < 0 {
			changePercent = downStyle.Render(fmt.Sprintf("%.2f%%", market.ChangePercent))
		}

		if market.Diff < 0 {
			trade = downStyle.Render(fmt.Sprintf("%.2f", market.Trade))
		}
		render := listStyle.Copy().Render(
			lipgloss.JoinVertical(lipgloss.Left,
				listHeader(fmt.Sprintf("%s    %s", market.Name, changePercent)),
				listItem(trade)),
		)

		marketRenders = append(marketRenders, render)

	}
	lists := lipgloss.JoinHorizontal(lipgloss.Top, marketRenders...)

	return lipgloss.JoinHorizontal(lipgloss.Top, lists)
}

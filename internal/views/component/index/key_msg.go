package index

import (
	tea "github.com/charmbracelet/bubbletea"
	"live-trading/internal/views/component/watchlist/stock"
)

func (m *Model) toggleInsertItem() tea.Cmd {
	m.input.Placeholder = inputPlaceholder
	m.input.Focus()
	m.input.PromptStyle = focusedStyle
	m.input.TextStyle = focusedStyle
	m.openInput = true
	inputModel, cmd := m.input.Update(nil)
	m.input = inputModel
	return cmd
}

func (m *Model) deleteItem() {
	component := m.components[m.selectedComponentIndex]
	code := component.GetRowCode()
	err := component.DeleteItem(m.ctx, code)
	if err != nil {
		m.errorModel.HandleError(err)
		return
	}

	switch component.(type) {
	case *stock.Model:
		go func() {
			err := m.stockService.RestartWatchPickStocks(m.ctx)
			m.errorModel.HandleError(err)
		}()
	}
}

func (m *Model) changeSelectedView() {
	maxSelected := len(m.components) - 1
	if m.selectedComponentIndex >= maxSelected {
		m.selectedComponentIndex = 0
		return
	}
	m.selectedComponentIndex++
}

func (m *Model) selectedDetail() {
	component := m.components[m.selectedComponentIndex]
	code := component.GetRowCode()
	detail, err := component.Detail(m.ctx, code)
	if err != nil {
		m.errorModel.HandleError(err)
	}
	if detail != "" {
		m.detail.InsertString(detail)
		m.detailShow = true
	}

}

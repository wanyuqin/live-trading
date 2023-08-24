package views

import (
	"context"
	"fmt"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"live-trading/internal/domain/service"
	"live-trading/internal/views/component/watchlist/stock"
	"time"
)

var (
	appStyle     = lipgloss.NewStyle().Padding(1, 2)
	focusedStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	windowWidth  int
	windowHeight int
)

type quoteMsg struct {
}

type Model struct {
	ctx          context.Context
	openInput    bool
	stock        stock.Model
	input        textinput.Model
	stockService service.IStock
}

func NewModel() *Model {
	input := textinput.New()
	input.Cursor.Style = focusedStyle.Copy()
	return &Model{
		stock:        stock.NewStockModel(),
		input:        input,
		stockService: service.NewStock(),
	}
}

func (m *Model) Init() tea.Cmd {

	// 初始化一些IO
	//return tea.Batch(tea.EnterAltScreen)
	return tea.Batch(quoteTick(), tea.EnterAltScreen)
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	if m.openInput {
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.String() {
			case "enter":
				code := m.input.Value()
				m.addStockCode(code)
				m.input = textinput.New()
				m.openInput = false
			}

		}

		cmd := m.updateInput(msg)
		cmds = append(cmds, cmd)
		return m, tea.Batch(cmds...)
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.stock.Keys.InsertItem):
			//m.stock.DelegateKeys.Remove.SetEnabled(true)
			m.input.Placeholder = "please input stock code"
			m.input.Focus()
			m.input.PromptStyle = focusedStyle
			m.input.TextStyle = focusedStyle
			m.openInput = true
			inputModel, cmd := m.input.Update(nil)
			m.input = inputModel
			cmds = append(cmds, cmd)
		}
	case quoteMsg:
		// 定期获取
		m.stock.RefreshTable()
		newTable, cmd := m.stock.Table.Update(msg)
		cmds = append(cmds, cmd)
		m.stock.Table = newTable
	}

	cmds = append(cmds, quoteTick())
	return m, tea.Batch(cmds...)
}

func (m *Model) updateInput(msg tea.Msg) tea.Cmd {
	if m.openInput {
		im, cmd := m.input.Update(msg)
		m.input = im
		return cmd
	}

	return nil

}

func (m *Model) addStockCode(code string) {
	err := m.stockService.AddPickStockCode(m.ctx, code)
	if err != nil {
		fmt.Println(err)
	}
}

func (m *Model) View() string {
	if m.openInput {
		return appStyle.Render(m.input.View())
	}
	return appStyle.Render(m.stock.Table.View())
}

func getTime() string {
	t := time.Now()
	return fmt.Sprintf("%s %02d:%02d:%02d", t.Weekday().String(), t.Hour(), t.Minute(), t.Second())
}

func quoteTick() tea.Cmd {
	return tea.Tick(3*time.Second, func(t time.Time) tea.Msg {
		return quoteMsg{}
	})
}

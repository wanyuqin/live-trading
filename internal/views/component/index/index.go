package index

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
	cancel       context.CancelCauseFunc
	openInput    bool
	stock        stock.Model
	input        textinput.Model
	stockService service.IStock
}

func NewModel() *Model {
	//ctx, cancel := context.WithCancelCause(context.Background())
	ctx := context.Background()
	stockService := service.NewStockContext(ctx)
	input := textinput.New()
	input.Cursor.Style = focusedStyle.Copy()
	return &Model{
		stock:        stock.NewStockModel(),
		input:        input,
		stockService: stockService,
		ctx:          ctx,
	}
}

func (m *Model) Init() tea.Cmd {
	m.startWatchPickStock()
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
			case "esc":
				m.input = textinput.New()
				m.openInput = false
			}
		}

		cmd := m.updateInput(msg)
		cmds = append(cmds, cmd)
		return m, tea.Batch(cmds...)
	}
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		width, height := appStyle.GetFrameSize()
		windowWidth = width
		windowHeight = height
		m.stock.Table.WithMaxTotalWidth(width)
	case tea.KeyMsg:
		switch msg.String() {
		case "x":
			m.deleteStockCode()
		case "ctrl+c":
			cmds = append(cmds, tea.Quit)
		}
		switch {
		case key.Matches(msg, m.stock.Keys.InsertItem):
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
	}
	newTable, cmd := m.stock.Table.Update(msg)
	m.stock.Table = newTable
	cmds = append(cmds, cmd)
	cmds = append(cmds, quoteTick())

	return m, tea.Batch(cmds...)
}

func (m *Model) View() string {
	now := getTime()
	if m.openInput {
		return now + "\n" + appStyle.Render(m.input.View())
	}

	return now + "\n" + lipgloss.NewStyle().MarginLeft(1).Render(m.stock.Table.View())
}

func (m *Model) updateInput(msg tea.Msg) tea.Cmd {
	if m.openInput {
		im, cmd := m.input.Update(msg)
		m.input = im
		return cmd
	}

	return nil

}

func (m *Model) startWatchPickStock() {
	go func() {
		err := m.stockService.WatchPickStocks()
		if err != nil {
			fmt.Println(err)
		}
	}()
}

func getTime() string {
	t := time.Now()
	return fmt.Sprintf("%d-%02d-%02d %s  %02d:%02d:%02d", t.Year(), t.Month(), t.Day(), t.Weekday().String(), t.Hour(), t.Minute(), t.Second())
}

func quoteTick() tea.Cmd {
	return tea.Tick(1*time.Second, func(t time.Time) tea.Msg {
		return quoteMsg{}
	})
}

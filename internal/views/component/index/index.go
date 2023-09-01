package index

import (
	"context"
	"fmt"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"live-trading/internal/domain/service"
	"live-trading/internal/views/component"
	"live-trading/internal/views/component/market"
	"live-trading/internal/views/component/watchlist/fund"
	"live-trading/internal/views/component/watchlist/stock"
	"strings"
	"time"
)

var (
	appStyle     = lipgloss.NewStyle().Padding(1, 2)
	focusedStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	windowWidth  int
	windowHeight int

	inputPlaceholder = "please input code"
)

type quoteMsg struct {
}

type Model struct {
	ctx           context.Context
	cancel        context.CancelCauseFunc
	stock         *stock.Model
	market        *market.Model
	fund          *fund.Model
	errorModel    *ErrorModel
	keys          *indexKeyMap
	input         textinput.Model
	help          help.Model
	stockService  service.IStock
	marketService service.IMarket
	fundService   service.IFund
	components    []component.Component
	detail        *FundDetailModel

	detailShow             bool
	selectedComponentIndex int
	openInput              bool
	width                  int
	height                 int
}

func NewModel() *Model {
	ctx := context.Background()
	stockService := service.NewStockWithContext(ctx)
	marketService := service.NewMarketWithContext(ctx)
	input := textinput.New()
	input.Cursor.Style = focusedStyle.Copy()
	stockModel := stock.NewStockModel(ctx)
	marketModel := market.NewModel(ctx)
	fundModel := fund.NewFundModel(ctx)
	errorModel := NewErrorModel()
	detailModel := NewFundDetailModel(ctx)
	model := Model{
		input:         input,
		stockService:  stockService,
		marketService: marketService,
		ctx:           ctx,
		market:        marketModel,
		fund:          fundModel,
		stock:         stockModel,
		errorModel:    errorModel,
		detail:        detailModel,
		keys:          newIndexKeyMap(),
	}
	model.addComponent(stockModel)
	model.addComponent(fundModel)

	return &model
}

func (m *Model) addComponent(c component.Component) {
	m.components = append(m.components, c)
}

func (m *Model) Init() tea.Cmd {

	m.startWatchPickStock()
	m.startWatchMarketStock()
	// 初始化一些IO
	return tea.Batch(tea.EnterAltScreen)
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	if m.openInput {
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.String() {
			case "enter":
				code := m.input.Value()
				err := m.components[m.selectedComponentIndex].AddItem(m.ctx, code)
				m.errorModel.HandleError(err)
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

	if m.detailShow {
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.String() {
			case "esc":
				m.detailShow = false
				return m, tea.Batch(cmds...)
			}
		}

	}

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.stock.Table.WithMaxTotalWidth(m.width)
	case tea.KeyMsg:
		switch msg.String() {
		case "x":
			m.deleteItem()
		case "ctrl+c":
			cmds = append(cmds, tea.Quit)
		}
		switch {
		//case key.Matches(msg, m.stock.Keys.InsertItem):
		case key.Matches(msg, m.keys.ToggleInsertItem):
			if !m.detailShow {
				cmd := m.toggleInsertItem()
				cmds = append(cmds, cmd)
			}
		case key.Matches(msg, m.keys.ToggleChangeView):
			m.changeSelectedView()
		case key.Matches(msg, m.keys.ToggleDetail):
			m.selectedDetail()

		}
	case quoteMsg:
		// 定期获取
		m.stock.RefreshTable()
		m.market.RefreshTable()
		m.fund.RefreshTable()
		m.errorModel.Restore()

	}

	newFundTable, cmd := m.fund.Table.Update(msg)
	cmds = append(cmds, cmd)
	m.fund.Table = newFundTable
	newTable, cmd := m.stock.Table.Update(msg)
	cmds = append(cmds, cmd)
	m.stock.Table = newTable
	m.detail.area, cmd = m.detail.area.Update(msg)
	cmds = append(cmds, cmd)
	//cmds = append(cmds, quoteTick())

	m.resetSize()
	return m, tea.Batch(cmds...)
}

func (m *Model) View() string {
	doc := strings.Builder{}
	doc.WriteString(m.market.View() + "\n\n")
	if m.openInput {
		doc.WriteString(appStyle.Render(m.input.View()))
		return doc.String()
	}
	if m.detailShow {
		doc.WriteString(m.detail.View())
		return doc.String()
	}
	doc.WriteString(m.components[m.selectedComponentIndex].View())
	doc.WriteString("\n\n\n\n")
	doc.WriteString(component.HelpStyle.Render(m.help.View(m.keys)))

	doc.WriteString("\n\n")
	doc.WriteString(m.errorModel.View())
	return lipgloss.JoinHorizontal(lipgloss.Top, doc.String())
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
			m.errorModel.HandleError(err)
		}
	}()
}

func (m *Model) startWatchMarketStock() {
	go func() {
		err := m.marketService.WatchMarket()
		if err != nil {
			m.errorModel.HandleError(err)
		}
	}()
}

func (m *Model) resetSize() {
	m.detail.area.Width = m.width
	m.detail.area.Height = 20
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

package index

import (
	"context"
	"github.com/charmbracelet/bubbles/viewport"
	"github.com/charmbracelet/glamour"
	"live-trading/internal/domain/service"
)

type FundDetailModel struct {
	ctx         context.Context
	fundService service.IFund
	area        viewport.Model
}

func NewFundDetailModel(ctx context.Context) *FundDetailModel {
	m := &FundDetailModel{
		ctx:         ctx,
		fundService: service.NewFund(),
		area:        viewport.New(windowWidth, windowHeight),
	}

	return m
}

func (m *FundDetailModel) View() string {
	renderer, err := glamour.NewTermRenderer(glamour.WithAutoStyle(), glamour.WithWordWrap(windowWidth))
	if err != nil {

	}
	detail, err := renderer.Render(m.area.View())
	return detail
}

func (m *FundDetailModel) InsertString(s string) {
	m.area.SetContent(s)
}

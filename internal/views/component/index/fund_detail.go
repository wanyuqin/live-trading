package fund

import (
	"context"
	"github.com/charmbracelet/bubbles/textarea"
	"live-trading/internal/domain/service"
)

type FundDetailModel struct {
	ctx         context.Context
	fundService service.IFund
	area        textarea.Model
}

func NewFundDetailModel(ctx context.Context) *FundDetailModel {
	return &FundDetailModel{
		ctx:         ctx,
		fundService: service.NewFund(),
	}
}

func (m *FundDetailModel) View() string {

	return ""
}

func (m *FundDetailModel) InsertString(s string) {
	m.area.InsertString(s)
}

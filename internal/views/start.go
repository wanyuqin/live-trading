package views

import (
	"context"
	"live-trading/internal/views/component/index"

	tea "github.com/charmbracelet/bubbletea"
)

func Start() error {
	ctx := context.Background()
	p := tea.NewProgram(index.NewModel(ctx), tea.WithAltScreen(), tea.WithContext(ctx))

	_, err := p.Run()
	if err != nil {
		return err
	}
	return nil
}

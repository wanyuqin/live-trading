package views

import (
	"context"
	tea "github.com/charmbracelet/bubbletea"
	"live-trading/internal/views/component/index"
)

func Start() error {
	ctx := context.Background()
	p := tea.NewProgram(index.NewModel(), tea.WithAltScreen(), tea.WithContext(ctx))

	_, err := p.Run()
	if err != nil {
		return err
	}
	return nil
}

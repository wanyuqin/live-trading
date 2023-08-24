package views

import (
	"context"
	tea "github.com/charmbracelet/bubbletea"
)

func Start() error {
	ctx := context.Background()
	p := tea.NewProgram(NewModel(), tea.WithAltScreen(), tea.WithContext(ctx))

	_, err := p.Run()
	if err != nil {
		return err
	}
	return nil
}

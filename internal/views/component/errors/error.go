package errors

import (
	"fmt"
	"github.com/charmbracelet/lipgloss"
)

const (
	initErrorShowTimes int = 0
	maxErrorShowTimes  int = 3
)

type Model struct {
	message   string
	showTimes int
}

var (
	errorMessageStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FFFDF5")).
		Background(lipgloss.Color("#FF5F87")).
		Padding(0, 1).
		MarginRight(1)
)

func NewModel() *Model {
	return &Model{}
}

func (m *Model) HandleError(err error) {
	if err != nil {
		if m.showTimes < maxErrorShowTimes {
			mes := fmt.Sprintf(":( ️Error Message: %s", err.Error())
			if m.message != "" {
				m.message = fmt.Sprintf("%s\n%s", mes, m.message)
			} else {
				m.message = mes
			}
			m.showTimes = initErrorShowTimes
		}
	}

}

func (m *Model) View() string {
	if m.message == "" {
		return ""
	}
	return errorMessageStyle.Render(m.message)
}

func (m *Model) Restore() {
	if m.message != "" {
		if m.showTimes > maxErrorShowTimes {
			m.showTimes = initErrorShowTimes
			m.message = ""
			return
		}
		m.showTimes++
	}

}

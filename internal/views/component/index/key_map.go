package index

import (
	"fmt"
	"github.com/charmbracelet/bubbles/key"
	"strings"
)

type indexKeyMap struct {
	ToggleChangeView key.Binding
	ToggleInsertItem key.Binding
	ToggleDeleteItem key.Binding
	ToggleDetail     key.Binding
}

func newIndexKeyMap() *indexKeyMap {
	return &indexKeyMap{
		ToggleChangeView: key.NewBinding(
			key.WithKeys("tab"),
			key.WithHelp(" [tab]", "change view ")),
		ToggleInsertItem: key.NewBinding(
			key.WithKeys("a"),
			key.WithHelp(" [a]", "add item ")),
		ToggleDeleteItem: key.NewBinding(
			key.WithKeys("x"),
			key.WithHelp(" [x]", "delete item ")),
		ToggleDetail: key.NewBinding(
			key.WithKeys("enter"),
			key.WithHelp(" [enter]", "show fund detail")),
	}
}

func (i *indexKeyMap) Help() string {
	doc := strings.Builder{}

	doc.WriteString(fmt.Sprintf("[%s] ", i.ToggleDetail.Help().Key))
	doc.WriteString(i.ToggleDetail.Help().Desc)
	doc.WriteString(fmt.Sprintf("[%s] ", i.ToggleDeleteItem.Help().Key))
	doc.WriteString(i.ToggleDeleteItem.Help().Desc)
	doc.WriteString(fmt.Sprintf("[%s] ", i.ToggleChangeView.Help().Key))
	doc.WriteString(i.ToggleChangeView.Help().Desc)
	doc.WriteString(fmt.Sprintf("[%s] ", i.ToggleInsertItem.Help().Key))
	doc.WriteString(i.ToggleInsertItem.Help().Desc)
	return doc.String()
}

func (i *indexKeyMap) ShortHelp() []key.Binding {
	return []key.Binding{
		i.ToggleDetail,
		i.ToggleInsertItem,
		i.ToggleDeleteItem,
		i.ToggleChangeView,
	}
}

func (i *indexKeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{
			i.ToggleDetail,
			i.ToggleInsertItem,
			i.ToggleDeleteItem,
			i.ToggleChangeView,
		},
	}
}

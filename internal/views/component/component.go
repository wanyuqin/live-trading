package component

import "context"

type Component interface {
	View() string

	AddItem(ctx context.Context, code string) error
	DeleteItem(ctx context.Context, code string) error
	GetRowCode() string
	Detail(ctx context.Context, code string) (string, error)
}

package gox

import (
	"context"
	"fmt"
)

func RunSafe(ctx context.Context, fn func(ctx context.Context)) {
	defer func() {

	}()
	go func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Printf("panic: %v\n", err)
			}
		}()
		fn(ctx)
	}()

}

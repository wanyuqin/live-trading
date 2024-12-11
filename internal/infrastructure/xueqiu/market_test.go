package xueqiu

import (
	"context"
	"fmt"
	"log"
	"testing"
)

func TestMarketRepoImpl_ListMarket(t *testing.T) {
	err := GetCookie()
	if err != nil {
		log.Fatal(err)
	}
	marketRepoImpl := NewMarketRepoImpl()

	markets, err := marketRepoImpl.ListMarket(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%#v", markets)
}

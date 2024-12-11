package xueqiu

import (
	"context"
	"fmt"
	"log"
	"testing"
)

func TestFundRepoImpl_GetPositionStocks(t *testing.T) {
	repoImpl := NewFundRepoImpl()
	fundPosition, err := repoImpl.GetFundPosition(context.Background(), "003095")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v", fundPosition)
}

func TestFundRepoImpl_GetFundManager(t *testing.T) {
	repoImpl := NewFundRepoImpl()
	fundManagers, err := repoImpl.GetFundManager(context.Background(), "003095")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v", fundManagers)
}

func TestFundRepoImpl_GetFundDetail(t *testing.T) {
	repoImpl := NewFundRepoImpl()
	fundDetail, err := repoImpl.GetFundDetail(context.Background(), "003095")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%#v", fundDetail)
}

func TestFundRepoImpl_GetFundSummary(t *testing.T) {
	repoImpl := NewFundRepoImpl()
	summary, err := repoImpl.GetFundSummary(context.Background(), "003095")
	if err != nil {
		return
	}

	fmt.Printf("%#v", summary)
}

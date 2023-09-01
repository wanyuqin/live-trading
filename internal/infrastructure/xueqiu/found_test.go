package xueqiu

import (
	"context"
	"fmt"
	"log"
	"testing"
)

func TestXueQiuFundRepoImpl_GetPositionStocks(t *testing.T) {
	repoImpl := NewXueQiuFundRepoImpl()
	fundPosition, err := repoImpl.GetFundPosition(context.Background(), "003095")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v", fundPosition)
}

func TestXueQiuFundRepoImpl_GetFundManager(t *testing.T) {
	repoImpl := NewXueQiuFundRepoImpl()
	fundManagers, err := repoImpl.GetFundManager(context.Background(), "003095")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v", fundManagers)
}

func TestXueQiuFundRepoImpl_GetFundDetail(t *testing.T) {
	repoImpl := NewXueQiuFundRepoImpl()
	fundDetail, err := repoImpl.GetFundDetail(context.Background(), "003095")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%#v", fundDetail)
}

func TestXueQiuFundRepoImpl_GetFundSummary(t *testing.T) {
	repoImpl := NewXueQiuFundRepoImpl()
	summary, err := repoImpl.GetFundSummary(context.Background(), "003095")
	if err != nil {
		return
	}

	fmt.Printf("%#v", summary)
}

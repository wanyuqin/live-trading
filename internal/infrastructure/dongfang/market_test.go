package dongfang

import (
	"fmt"
	"testing"
)

func TestDongFangMarketRepoImpl_ListMarket(t *testing.T) {
	res := make(chan []byte, 10)

	go NewDongFangMarketRepoImpl().ListMarket(res)
	for re := range res {
		fmt.Println(string(re))
	}
}

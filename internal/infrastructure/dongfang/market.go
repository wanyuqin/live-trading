package dongfang

import (
	"context"
	"fmt"
	"live-trading/internal/domain/repository"
	"live-trading/internal/infrastructure/client"
	"sync"
)

var (
	// 上证指数
	szzs = "http://23.push2.eastmoney.com/api/qt/stock/sse?fields=f58,f43,f169,f170&secid=1.000001"
	// 深证成指
	sczs = "http://76.push2.eastmoney.com/api/qt/stock/sse?fields=f58,f43,f169,f170&secid=0.399001"
	//
	cybz = "http://99.push2.eastmoney.com/api/qt/stock/sse?secid=0.399006&fields=f58,f43,f169,f170"
)

type DongFangMarketRepoImpl struct {
	repository.MarketRepo
}

func NewDongFangMarketRepoImpl() *DongFangMarketRepoImpl {
	return &DongFangMarketRepoImpl{}
}

type MarketResponse struct {
}

// ListMarket TODO 重构
func (d DongFangMarketRepoImpl) ListMarket(res chan<- []byte) {
	wg := &sync.WaitGroup{}
	wg.Add(3)
	c := client.NewClient()
	go func() {
		defer wg.Done()
		request, err := c.NewRequest(context.Background(), "GET", szzs)
		if err != nil {
			fmt.Println(err)

		}
		stream, err := c.SendRequestStream(request)
		if err != nil {
			fmt.Println(err)
		}

		for {
			body, err := stream.ProcessLine()
			if err != nil {
				fmt.Println(err)
			}
			res <- body

		}
	}()

	go func() {
		defer wg.Done()
		request, err := c.NewRequest(context.Background(), "GET", sczs)
		if err != nil {
			fmt.Println(err)

		}
		stream, err := c.SendRequestStream(request)
		if err != nil {
			fmt.Println(err)
		}

		for {
			body, err := stream.ProcessLine()
			if err != nil {
				fmt.Println(err)
			}
			res <- body

		}
	}()

	go func() {
		defer wg.Done()
		request, err := c.NewRequest(context.Background(), "GET", cybz)
		if err != nil {
			fmt.Println(err)
		}
		stream, err := c.SendRequestStream(request)
		if err != nil {
			fmt.Println(err)
		}
		for {
			body, err := stream.ProcessLine()
			if err != nil {
				fmt.Println(err)
			}
			res <- body
		}
	}()
	wg.Wait()

}

package dongfang

import (
	"context"
	"live-trading/internal/domain/entity"
	"live-trading/internal/domain/repository"
	"live-trading/internal/infrastructure/client"
)

var (
	// 上证指数
	szzs = "http://23.push2.eastmoney.com/api/qt/stock/sse?fields=f58,f43,f169,f170&secid=1.000001,0.399001,0.399006"
	// 深证成指
	sczs = "http://76.push2.eastmoney.com/api/qt/stock/sse?fields=f58,f43,f169,f170&secid=0.399001"
	//
	cybz = "http://99.push2.eastmoney.com/api/qt/stock/sse?secid=0.399006&fields=f58,f43,f169,f170"
)

type MarketRepoImpl struct {
	repository.MarketRepo
}

func NewMarketRepoImpl() *MarketRepoImpl {
	return &MarketRepoImpl{}
}

type MarketResponse struct {
}

func (d *MarketRepoImpl) WatchMarket(ctx context.Context, codes entity.StockCodes, rec chan<- []entity.PickStock) error {
	c := client.NewClient()
	u, err := GetStockUrl(codes.RequestCodes())
	defer close(rec)
	if err != nil {
		return err
	}
	request, err := c.NewRequest(context.Background(), "GET", u)
	if err != nil {
		return err
	}

	stream, err := c.SendRequestStream(request)
	if err != nil {
		return err
	}

	for {
		select {
		case <-ctx.Done():
			return nil
		default:
			body, err := stream.ProcessLine()
			if err != nil {
				return err
			}
			pickStocks, err := ParseWatchPickStock(body)
			if err != nil {
				return err
			}

			if len(pickStocks) > 0 {
				rec <- pickStocks
			}
		}

	}
}

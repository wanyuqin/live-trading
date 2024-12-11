package dongfang

import (
	"context"
	"live-trading/internal/domain/entity"
	"live-trading/internal/domain/repository"
	"live-trading/internal/infrastructure/client"
)

type StockRepoImpl struct {
	repository.StockRepo
}

func NewStockRepoImpl() *StockRepoImpl {
	return &StockRepoImpl{}
}

func (d *StockRepoImpl) WatchPickStock(ctx context.Context, codes entity.StockCodes, rec chan []entity.PickStock) error {
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

package xueqiu

import (
	"context"
	"io"
	"live-trading/internal/domain/entity"
	"live-trading/internal/domain/repository"
	"live-trading/internal/infrastructure/client"
	"net/http"
)

type MarketRepoImpl struct {
	repository.MarketRepo
}

func NewMarketRepoImpl() *MarketRepoImpl {
	return &MarketRepoImpl{}
}

func (m MarketRepoImpl) ListMarket(ctx context.Context) (entity.Markets, error) {
	c := client.NewClient()
	request, err := c.NewRequest(ctx, http.MethodGet, listMarketUrl)
	if err != nil {
		return nil, nil
	}
	if len(cookies) == 0 {
		err = GetCookie()
		if err != nil {
			return nil, err
		}
	}
	for i := range cookies {
		cookie := cookies[i]
		request.AddCookie(cookie)
	}
	response, err := c.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, nil
	}
	markets, err := parseListMarket(body)
	return markets, err
}

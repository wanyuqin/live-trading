package xueqiu

import (
	"context"
	"fmt"
	"io"
	"live-trading/internal/domain/entity"
	"live-trading/internal/domain/repository"
	"live-trading/internal/infrastructure/client"
	"net/http"
	"net/url"
)

type FundRepoImpl struct {
	repository.FundRepo
}

func NewFundRepoImpl() *FundRepoImpl {
	return &FundRepoImpl{}
}

func (repo *FundRepoImpl) GetFundPosition(ctx context.Context, code string) (entity.FundPosition, error) {
	client := client.NewClient()
	values := url.Values{}
	values.Set("fund_code", code)
	request, err := client.NewRequest(ctx, http.MethodGet, positionStockUrl)
	if err != nil {
		return entity.FundPosition{}, err
	}
	request.URL.RawQuery = values.Encode()
	response, err := client.Do(request)
	defer response.Body.Close()
	if err != nil {
		return entity.FundPosition{}, err
	}
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return entity.FundPosition{}, err
	}

	return parseFundPosition(body)

}

func (repo *FundRepoImpl) GetFundManager(ctx context.Context, code string) (entity.FundManagerList, error) {
	values := url.Values{}
	values.Set("post_status", "1")
	values.Set("fund_code", code)

	client := client.NewClient()
	request, err := client.NewRequest(ctx, http.MethodGet, fundManagerUrl)
	if err != nil {
		return nil, err
	}
	request.URL.RawQuery = values.Encode()
	response, err := client.Do(request)
	defer response.Body.Close()
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return parseFundManager(body)
}

func (repo *FundRepoImpl) GetFundDetail(ctx context.Context, code string) (entity.FoundDetail, error) {
	u := fmt.Sprintf("%s/%s", fundDetailUrl, code)
	client := client.NewClient()
	request, err := client.NewRequest(ctx, http.MethodGet, u)
	if err != nil {
		return entity.FoundDetail{}, err
	}
	response, err := client.Do(request)
	defer response.Body.Close()
	if err != nil {
		return entity.FoundDetail{}, err
	}
	data, err := io.ReadAll(response.Body)

	if err != nil {
		return entity.FoundDetail{}, err
	}
	return parseFundDetail(data)
}

func (repo *FundRepoImpl) GetFundSummary(ctx context.Context, code string) (entity.FundSummary, error) {
	u := fmt.Sprintf("%s/%s", fundSummaryUrl, code)
	client := client.NewClient()
	request, err := client.NewRequest(ctx, http.MethodGet, u)
	if err != nil {
		return entity.FundSummary{}, err
	}
	response, err := client.Do(request)
	defer response.Body.Close()
	if err != nil {
		return entity.FundSummary{}, err
	}
	data, err := io.ReadAll(response.Body)
	if err != nil {
		return entity.FundSummary{}, err
	}
	return parseFundSummary(data)
}

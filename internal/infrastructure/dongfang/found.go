package dongfang

import (
	"context"
	"fmt"
	"io"
	"live-trading/internal/domain/entity"
	"live-trading/internal/domain/repository"
	"live-trading/internal/infrastructure/client"
	"net/http"
	"net/url"
	"strings"
)

type DongFangFundRepoImpl struct {
	repository.FundRepo
}

func NewDongFangFundRepoImpl() *DongFangFundRepoImpl {
	return &DongFangFundRepoImpl{}
}

func getFundListInfoUrl() {

}

func (d *DongFangFundRepoImpl) ListFundsInfo(ctx context.Context, fcodes []string) (entity.FundList, error) {
	client := client.NewClient()
	request, err := client.NewRequest(ctx, http.MethodPost, "https://api.fund.eastmoney.com/favor/GetFundsInfo")
	if err != nil {
		return nil, err
	}

	request.Header.Set("Origin", "https://favor.fund.eastmoney.com")
	request.Header.Set("Referer", "https://favor.fund.eastmoney.com")
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	values := url.Values{
		"fcodes": []string{strings.Join(fcodes, ",")},
	}

	request.Body = io.NopCloser(strings.NewReader(values.Encode()))

	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return ParseFundList(body), nil

}

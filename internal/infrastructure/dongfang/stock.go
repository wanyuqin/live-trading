package dongfang

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/duke-git/lancet/v2/strutil"
	"live-trading/internal/domain/entity"
	"live-trading/internal/domain/repository"
	"live-trading/internal/infrastructure/client"
	"live-trading/tools/xmath"
	"math/rand"
	"net/url"
	"strings"
)

type DongFangStockRepoImpl struct {
	repository.StockRepo

	WatchState int
}

func NewDongFangStockRepoImpl() *DongFangStockRepoImpl {
	return &DongFangStockRepoImpl{}
}

var schema = "https"
var quoteApi = "push2.eastmoney.com"
var WatchState int = 0

const (
	WatchStateStart = iota
	WatchStateStop
)

func GetStockHost() string {
	i := rand.Int31n(100) + 1
	return fmt.Sprintf("%s://%d.%s", schema, i, quoteApi)
}

func GetStockUrl(stockCode []string) (string, error) {
	host := GetStockHost()
	u := "/api/qt/ulist/sse"
	url, err := url.Parse(fmt.Sprintf("%s%s", host, u))
	if err != nil {
		return "", err
	}
	query := url.Query()
	query.Add("mpi", "2000")
	query.Add("fields", "f12,f13,f19,f14,f139,f148,f2,f4,f1,f125,f18,f3,f152,f5,f30,f31,f32,f6,f8,f7,f10,f22,f9,f112,f100")
	query.Add("pi", "0")
	query.Add("secids", strings.Join(stockCode, ","))

	url.RawQuery = query.Encode()
	return url.String(), nil
}

func (d *DongFangStockRepoImpl) WatchPickStock(ctx context.Context, codes entity.StockCodes, rec chan []entity.PickStock) error {
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
		if !d.IsStartWatch() {
			return nil
		}
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

func ParseWatchPickStock(data []byte) ([]entity.PickStock, error) {
	trim := strutil.Trim(string(data), "data:")
	data = []byte(trim)

	resp := struct {
		Rc     int    `json:"rc"`
		Rt     int    `json:"rt"`
		Svr    int64  `json:"svr"`
		Lt     int    `json:"lt"`
		Full   int    `json:"full"`
		Dlmkts string `json:"dlmkts"`
		Data   struct {
			Total int `json:"total"`
			Diff  map[string]struct {
				F1   int     `json:"f1,omitempty"`
				F2   int     `json:"f2,omitempty"`
				F3   int     `json:"f3,omitempty"`
				F4   int     `json:"f4,omitempty"`
				F5   int     `json:"f5,omitempty"`
				F6   float64 `json:"f6,omitempty"`
				F7   int     `json:"f7,omitempty"`
				F8   int     `json:"f8,omitempty"`
				F9   int     `json:"f9,omitempty"`
				F10  int     `json:"f10,omitempty"`
				F12  string  `json:"f12,omitempty"`
				F13  int     `json:"f13,omitempty"`
				F14  string  `json:"f14,omitempty"`
				F18  int     `json:"f18,omitempty"`
				F19  int     `json:"f19,omitempty"`
				F22  int     `json:"f22,omitempty"`
				F30  int     `json:"f30,omitempty"`
				F31  int     `json:"f31,omitempty"`
				F32  int     `json:"f32,omitempty"`
				F100 string  `json:"f100,omitempty"`
				F112 float64 `json:"f112,omitempty"`
				F125 int     `json:"f125,omitempty"`
				F139 int     `json:"f139,omitempty"`
				F148 int     `json:"f148,omitempty"`
				F152 int     `json:"f152,omitempty"`
			} `json:"diff"`
		}
	}{}
	r := resp
	err := json.Unmarshal(data, &r)
	if err != nil {
		return nil, err
	}

	pickStocks := make([]entity.PickStock, 0)

	for _, stock := range r.Data.Diff {
		ps := entity.PickStock{
			DataId: r.Svr,
			Name:   stock.F14,
			Code:   stock.F12,
			Now:    xmath.DivideByHundred(stock.F2),
			Diff:   xmath.DivideByHundred(stock.F4),
		}
		pickStocks = append(pickStocks, ps)
	}

	return pickStocks, err
}

func (d *DongFangStockRepoImpl) StartWatch() {
	d.WatchState = WatchStateStart
}

func (d *DongFangStockRepoImpl) StopWatch() {
	d.WatchState = WatchStateStop
}

func (d *DongFangStockRepoImpl) IsStartWatch() bool {
	return d.WatchState == WatchStateStart
}

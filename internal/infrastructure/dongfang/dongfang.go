package dongfang

import (
	"encoding/json"
	"fmt"
	"github.com/duke-git/lancet/v2/slice"
	"github.com/duke-git/lancet/v2/strutil"
	"live-trading/internal/configs"
	"live-trading/internal/domain/entity"
	"live-trading/tools/xmath"
	"math/rand"
	"net/url"
	"strings"
)

var (
	schema   = "https"
	quoteApi = "push2.eastmoney.com"
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
				F3   int     `json:"f3,omitempty"` // 涨跌幅
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
			DataId:        r.Svr,
			Name:          stock.F14,
			Code:          stock.F12,
			Trade:         xmath.DivideByHundred(stock.F2),
			Diff:          xmath.DivideByHundred(stock.F4),
			ChangePercent: xmath.DivideByHundred(stock.F3),
		}
		pickStocks = append(pickStocks, ps)
	}

	return pickStocks, err
}

func ParseFundList(src []byte) entity.FundList {
	resp := struct {
		Data struct {
			KFS []struct {
				FCODE        string      `json:"FCODE"`
				SHORTNAME    string      `json:"SHORTNAME"`
				ISHOT        string      `json:"ISHOT"`
				ISBUY        bool        `json:"ISBUY"`
				DTZT         string      `json:"DTZT"`
				DWJZ         string      `json:"DWJZ"`
				LJJZ         string      `json:"LJJZ"`
				FUNDTYPE     string      `json:"FUNDTYPE"`
				FTYPE        string      `json:"FTYPE"`
				RZDE         string      `json:"RZDE"`
				RZDF         string      `json:"RZDF"`
				FSRQ         string      `json:"FSRQ"`
				IPESTART1    string      `json:"IPESTART1"`
				IPEEND1      string      `json:"IPEEND1"`
				SHIPESTART1  string      `json:"SHIPESTART1"`
				SHIPEEND1    string      `json:"SHIPEEND1"`
				ISSALES      string      `json:"ISSALES"`
				TradeBuyType int         `json:"TradeBuyType"`
				Gsz          string      `json:"gsz"`
				Gszzl        string      `json:"gszzl"`
				Order        int         `json:"Order"`
				SGZT         string      `json:"SGZT"`
				ISSBDATE     interface{} `json:"ISSBDATE"`
				ISSEDATE     interface{} `json:"ISSEDATE"`
				ISNEW        interface{} `json:"ISNEW"`
				KFR          interface{} `json:"KFR"`
				SYLLN        string      `json:"SYL_LN"`
			} `json:"KFS"`
			HBX                []interface{} `json:"HBX"`
			LCX                []interface{} `json:"LCX"`
			CN                 []interface{} `json:"CN"`
			HK                 interface{}   `json:"HK"`
			GD                 interface{}   `json:"GD"`
			Fcodes             []string      `json:"Fcodes"`
			Orders             []string      `json:"Orders"`
			IsShowSetRecommend int           `json:"IsShowSetRecommend"`
		} `json:"Data"`
		ErrCode    int         `json:"ErrCode"`
		ErrMsg     interface{} `json:"ErrMsg"`
		TotalCount int         `json:"TotalCount"`
		Expansion  struct {
			BjTime     string `json:"bjTime"`
			UpdateTime string `json:"updateTime"`
		} `json:"Expansion"`
		PageSize  int `json:"PageSize"`
		PageIndex int `json:"PageIndex"`
	}{}

	err := json.Unmarshal(src, &resp)
	if err != nil {
		fmt.Println(err)
	}

	funds := make([]entity.Fund, 0, len(resp.Data.KFS))
	for i := range resp.Data.KFS {
		kfs := resp.Data.KFS[i]
		if slice.Contain(configs.GetConfig().WatchList.Fund, kfs.FCODE) {
			Fund := entity.Fund{
				Name:             kfs.SHORTNAME,
				Code:             kfs.FCODE,
				FNAV:             kfs.DWJZ,
				TRF:              kfs.LJJZ,
				DailyInc:         kfs.RZDE,
				DailyIncRate:     kfs.RZDF,
				InceptionIncRate: kfs.SYLLN,
				Type:             kfs.FTYPE,
				DataTime:         kfs.FSRQ,
			}
			funds = append(funds, Fund)
		}

	}
	return funds
}

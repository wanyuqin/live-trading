package entity

import "fmt"

type Stock struct {
}

type PickStock struct {
	DataId        int64   `json:"data_id"`        // 每次的ID是一致的
	Name          string  `json:"name"`           // f14
	Code          string  `json:"code"`           // code f12
	Trade         float64 `json:"trade"`          // 当前价f2
	Diff          float64 `json:"diff"`           // 差值 f4
	ChangePercent float64 `json:"change_percent"` // 涨跌幅 //
}

type StockCode string

type StockCodes []StockCode

func NewStockCodes(codes []string) StockCodes {
	stockCodes := make([]StockCode, 0, len(codes))
	for i := range codes {
		stockCodes = append(stockCodes, StockCode(codes[i]))
	}

	return stockCodes
}

func (s StockCodes) GetRequestCode() {

}

func (s StockCodes) RequestCodes() []string {
	codes := make([]string, 0, len(s))
	for i := range s {
		codes = append(codes, fmt.Sprintf("0.%s", s[i].String()))
	}
	return codes
}

func (s StockCode) String() string {
	return string(s)
}

var globalPickStock []PickStock

func GetGlobalPickStock() []PickStock {
	return globalPickStock
}

func NewGlobalPickStock() {
	globalPickStock = make([]PickStock, 0)
	return
}

func RefreshGlobalPickStock(picStock []PickStock) {
	if len(globalPickStock) == 0 {
		globalPickStock = picStock
		return
	}
	m := PickStockMap(picStock)

	for i := range globalPickStock {
		if newPickStock, ok := m[MapKey(globalPickStock[i].DataId, globalPickStock[i].Code)]; ok {
			if newPickStock.Trade > 0 {
				globalPickStock[i].Trade = newPickStock.Trade
				globalPickStock[i].Diff = newPickStock.Diff
				globalPickStock[i].ChangePercent = newPickStock.ChangePercent
			}

		}
	}
}

func PickStockMap(pickStocks []PickStock) map[string]PickStock {
	m := make(map[string]PickStock)
	for i := range pickStocks {
		pickStock := pickStocks[i]
		m[MapKey(pickStock.DataId, pickStock.Code)] = pickStock
	}

	return m
}

func MapKey(dataId int64, code string) string {
	return fmt.Sprintf("%d-%s", dataId, code)
}

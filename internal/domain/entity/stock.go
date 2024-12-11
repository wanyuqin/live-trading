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

var MarketCode = []StockCode{
	"000001", // 上证指数
	"399001", // 深圳成指
	"399006", // 创业板指
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
		codes = append(codes, s[i].RequestCode())
	}
	return codes
}

func (s StockCode) RequestCode() string {
	if s == "000001" {
		return fmt.Sprintf("1.%s", s.String())
	}
	first := s[:1]
	if first == "6" || first == "5" {
		return fmt.Sprintf("1.%s", s.String())
	}

	if first == "3" || first == "0" || first == "1" {
		return fmt.Sprintf("0.%s", s.String())
	}
	return ""
}

func (s StockCode) String() string {
	return string(s)
}

var (
	globalPickStock   = make([]PickStock, 0)
	globalMarketStock = make([]PickStock, 0)
)

func GetGlobalPickStock() []PickStock {
	return globalPickStock
}

func ClearGlobalPickStock() {
	globalPickStock = make([]PickStock, 0)
}

func GetGlobalMarketStock() []PickStock {
	return globalMarketStock
}

func NewGlobalPickStock() {
	globalPickStock = make([]PickStock, 0)
	return
}

func RefreshGlobalPickStock(picStock []PickStock) {
	refreshGlobalStock(picStock, &globalPickStock)
}

func RefreshGlobalMarketStock(picStock []PickStock) {
	refreshGlobalStock(picStock, &globalMarketStock)
}

func refreshGlobalStock(newStock []PickStock, globalStocks *[]PickStock) {
	if len(*globalStocks) == 0 {
		*globalStocks = newStock
		return
	}
	m := PickStockMap(newStock)
	copyStocks := *globalStocks
	for i := range copyStocks {
		if newPickStock, ok := m[MapKey(copyStocks[i].DataId, copyStocks[i].Code)]; ok {
			if newPickStock.Trade > 0 {
				copyStocks[i].Trade = newPickStock.Trade
				copyStocks[i].Diff = newPickStock.Diff
				if newPickStock.Diff != 0 && newPickStock.ChangePercent != 0 {
					copyStocks[i].ChangePercent = newPickStock.ChangePercent
				}
			}

		}
	}

	*globalStocks = copyStocks
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

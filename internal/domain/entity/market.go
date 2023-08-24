package entity

import "encoding/json"

type Market struct {
	Name         string `json:"name"`          // 名称
	Current      string `json:"current"`       //当前值
	Float        string `json:"float"`         // 浮动
	FloatPercent string `json:"float_percent"` // 浮动百分
}

type MarketRaw struct {
	F43  int64  `json:"f43"`
	F58  string `json:"f58"`
	F169 int64  `json:"f169"`
	F170 int64  `json:"f170"`
}

func NewMarketRawWithJsonRaw(raw []byte) (MarketRaw, error) {
	origin := struct {
		Data struct {
			Rc     int       `json:"rc"`
			Rt     int       `json:"rt"`
			Svr    int       `json:"svr"`
			Lt     int       `json:"lt"`
			Full   int       `json:"full"`
			Dlmkts string    `json:"dlmkts"`
			Data   MarketRaw `json:"data"`
		} `json:"data"`
	}{}

	err := json.Unmarshal(raw, &origin)
	if err != nil {
		return MarketRaw{}, err
	}

	mr := MarketRaw{}
	mr = origin.Data.Data
	return mr, nil
}

type Markets []Market

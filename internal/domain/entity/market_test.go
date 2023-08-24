package entity

import "testing"

func TestNewMarketRawWithJsonRaw(t *testing.T) {
	s := `data: {"rc":0,"rt":4,"svr":182482656,"lt":1,"full":1,"dlmkts":"","data":{"f43":309298,"f58":"上证指数","f169":-3897,"f170":-124}}`
	NewMarketRawWithJsonRaw([]byte(s))
}

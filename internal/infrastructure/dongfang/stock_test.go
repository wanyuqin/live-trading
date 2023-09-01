package dongfang

import (
	"fmt"
	"log"
	"testing"
)

func TestGetStockHost(t *testing.T) {
	host := GetStockHost()
	fmt.Println(host)
}

func TestGetStockUrl(t *testing.T) {
	//0.000882,0.002336,0.002176,0.300059
	url, err := GetStockUrl([]string{"0.000882", "0.002336", "0.002176", "0.300059"})
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println(url)
}

func TestDongFangStockRepoImpl_WatchPickStock(t *testing.T) {
	//stockRepoImpl := NewDongFangStockRepoImpl()
	//rec := make(chan []entity.PickStock, 1000)
	////0.000882,0.002336,0.002176,0.300059
	//go func() {
	//	err := stockRepoImpl.WatchPickStock(entity.StockCodes{
	//		entity.StockCode("000882"),
	//		entity.StockCode("002336"),
	//		entity.StockCode("002176"),
	//		entity.StockCode("300059"),
	//	}, rec)
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//}()
	//
	//for stock := range rec {
	//	fmt.Printf("%#v\n", stock)
	//}

}

func TestParseWatchPickStock(t *testing.T) {

	data := `data: {"rc":0,"rt":5,"svr":180606324,"lt":1,"full":1,"dlmkts":"","data":{"total":4,"diff":{"0":{"f1":2,"f2":187,"f3":-158,"f4":-3,"f5":305354,"f6":57473902.0,"f7":211,"f8":112,"f9":15168,"f10":59,"f12":"000882","f13":0,"f14":"华联股份","f18":190,"f19":6,"f22":0,"f30":5981,"f31":186,"f32":187,"f100":"商业百货","f112":0.003082134,"f125":0,"f139":2,"f148":1,"f152":2},"1":{"f1":2,"f2":1502,"f3":-228,"f4":-35,"f5":105524,"f6":160100475.2,"f7":423,"f8":283,"f9":-1676,"f10":74,"f12":"002336","f13":0,"f14":"人人乐","f18":1537,"f19":6,"f22":-20,"f30":-1196,"f31":1502,"f32":1503,"f100":"商业百货","f112":-0.224058246,"f125":0,"f139":2,"f148":1025,"f152":2},"2":{"f1":2,"f2":920,"f3":-397,"f4":-38,"f5":318933,"f6":297280622.78,"f7":365,"f8":187,"f9":7749,"f10":108,"f12":"002176","f13":0,"f14":"江特电机","f18":958,"f19":6,"f22":0,"f30":-5155,"f31":920,"f32":921,"f100":"电机","f112":0.029682578,"f125":0,"f139":2,"f148":1089,"f152":2},"3":{"f1":2,"f2":1534,"f3":-204,"f4":-32,"f5":2001232,"f6":3089386792.63,"f7":153,"f8":150,"f9":2879,"f10":54,"f12":"300059","f13":0,"f14":"东方财富","f18":1566,"f19":80,"f22":0,"f30":-40640,"f31":1534,"f32":1535,"f100":"互联网服务","f112":0.266421152,"f125":0,"f139":5,"f148":1089,"f152":2}}}}`
	fmt.Printf("%s\n", []byte(data))
	_, err := ParseWatchPickStock([]byte(data))
	if err != nil {
		log.Fatal(err)
	}
}

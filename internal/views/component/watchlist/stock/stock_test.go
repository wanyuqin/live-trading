package stock

import (
	"live-trading/internal/configs"
	"log"
	"testing"
)

func TestStartWatchPickStocks(t *testing.T) {
	err := configs.LoadConfig("")
	if err != nil {
		log.Fatal(err)
	}
	StartWatchPickStocks()
}

package cmd

import (
	"context"
	"github.com/spf13/cobra"
	"live-trading/internal/configs"
	"live-trading/internal/domain/service"
	"live-trading/internal/views"
	"log"
)

var rootCmd = &cobra.Command{
	Use: "trading",
	Run: tradingStart,
}

func init() {

}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}

func tradingStart(cmd *cobra.Command, args []string) {
	err := configs.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	go startWatchPickStocks()

	err = views.Start()
	if err != nil {
		log.Fatal(err)
	}

}

func startWatchPickStocks() {
	stockService := service.NewStock()

	go func() {
		err := stockService.WatchPickStocks(context.Background())
		if err != nil {
			log.Fatal(err)
		}
	}()
}

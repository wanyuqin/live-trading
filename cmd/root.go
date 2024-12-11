package cmd

import (
	"live-trading/internal/configs"
	"live-trading/internal/views"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var (
	defaultConfigFileName = "trading.yaml"
	defaultConfigDirName  = "live_trading"
	configPath            string

	rootCmd = &cobra.Command{
		Use: "trading",
		Run: tradingStart,
	}
)

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.Flags().StringVarP(&configPath, "config", "c", "", "use customer config")
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}

func tradingStart(cmd *cobra.Command, args []string) {
	err := views.Start()
	if err != nil {
		log.Fatal(err)
	}

}

func initConfig() {
	if configPath == "" {
		// init default config in home dir
		homeDir, err := os.UserHomeDir()
		if err != nil {
			panic(err)
		}

		dirPath := filepath.Join(homeDir, defaultConfigDirName)
		_, err = os.Stat(dirPath)
		if os.IsNotExist(err) {
			err = os.MkdirAll(dirPath, 0777)
			if err != nil {
				panic(err)
			}
		}
		configPath = filepath.Join(homeDir, defaultConfigDirName, defaultConfigFileName)

	}
	err := configs.LoadConfig(configPath)
	if err != nil {
		panic(err)
	}
}

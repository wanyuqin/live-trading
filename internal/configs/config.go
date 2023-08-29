package configs

import (
	"embed"
	_ "embed"
	"errors"
	"github.com/duke-git/lancet/v2/slice"
	"gopkg.in/yaml.v3"
	"os"
)

type ApplicationConfig struct {
	WatchList WatchList `json:"watchList" yaml:"watchList"`
}

type WatchList struct {
	Stock []string `json:"stock" yaml:"stock"`
}

//go:embed trading.yaml
var configFile embed.FS

var cfg *ApplicationConfig

var configName = ""

func LoadConfig(configPath string) error {
	if configPath == "" {
		return errors.New("config path is empty")
	}

	err := CheckOrCreate(configPath)
	if err != nil {
		return err
	}

	body, err := os.ReadFile(configPath)
	if err != nil {
		return err
	}

	var ac ApplicationConfig

	err = yaml.Unmarshal(body, &ac)
	if err != nil {
		return err
	}

	cfg = &ac
	cfg.UniqueWatchStock()
	configName = configPath
	return nil
}

func CheckOrCreate(configPath string) error {
	_, err := os.Stat(configPath)
	if os.IsNotExist(err) {
		err = createDefaultConfigFile(configPath)
		if err != nil {
			return err
		}
	}
	return err
}

func createDefaultConfigFile(configPath string) error {
	file, err := os.Create(configPath)
	defer file.Close()
	if err != nil {
		return err
	}
	defaultConfig := ApplicationConfig{}
	body, err := yaml.Marshal(defaultConfig)
	if err != nil {
		return err
	}
	_, err = file.Write(body)
	return err

}

func GetConfig() *ApplicationConfig {
	return cfg
}

func (config *ApplicationConfig) AddStockCode(code string) error {
	if slice.Contain(config.WatchList.Stock, code) {
		return nil
	}
	config.WatchList.Stock = append(config.WatchList.Stock, code)
	return config.refreshConfig()
}

func (config *ApplicationConfig) DeleteStockCode(code string) error {
	if !slice.Contain(config.WatchList.Stock, code) {
		return nil
	}
	index := slice.IndexOf(config.WatchList.Stock, code)

	config.WatchList.Stock = slice.DeleteAt(config.WatchList.Stock, index)
	return config.refreshConfig()
}

func (config *ApplicationConfig) UniqueWatchStock() {
	config.WatchList.Stock = slice.Unique(config.WatchList.Stock)
}

func (config *ApplicationConfig) refreshConfig() error {
	file, err := os.OpenFile(configName, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0666)
	defer file.Close()

	if err != nil {
		return err
	}

	body, err := config.Marshal()
	if err != nil {
		return err
	}

	_, err = file.Write(body)
	if err != nil {
		return err
	}
	return nil
}

func (config *ApplicationConfig) Marshal() ([]byte, error) {
	return yaml.Marshal(config)
}

func RefreshConfig() error {
	return cfg.refreshConfig()
}

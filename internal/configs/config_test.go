package configs

import (
	"log"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	err := LoadConfig("")
	if err != nil {
		log.Fatal(err)
	}
}

func TestApplicationConfig_RefreshConfig(t *testing.T) {
	config := ApplicationConfig{
		WatchList: WatchList{
			Stock: []string{"000217", "000312"},
		},
	}
	err := config.refreshConfig()
	if err != nil {
		log.Fatal(err)
		return
	}

}

package xueqiu

import (
	"log"
	"testing"
)

func TestGetCookie(t *testing.T) {
	err := GetCookie()
	if err != nil {
		log.Fatal(err)
	}
}

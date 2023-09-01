package service

import (
	"context"
	"fmt"
	"log"
	"testing"
)

func TestFundService_GetFundDetail(t *testing.T) {
	fund := NewFund()
	fundDetail, err := fund.GetFundDetail(context.Background(), "003095")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%#v", fundDetail)
}

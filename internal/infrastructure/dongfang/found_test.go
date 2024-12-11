package dongfang

import (
	"context"
	"testing"
)

func TestFundRepoImpl_ListfundsInfo(t *testing.T) {
	fundRepoImpl := NewFundRepoImpl()
	//,,006228,003095,005918,001552,400015,005062
	fundRepoImpl.ListFundsInfo(context.Background(), []string{"012414"})
}

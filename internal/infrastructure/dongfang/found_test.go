package dongfang

import (
	"context"
	"testing"
)

func TestDongFangfundRepoImpl_ListfundsInfo(t *testing.T) {
	fundRepoImpl := NewDongFangFundRepoImpl()
	//,,006228,003095,005918,001552,400015,005062
	fundRepoImpl.ListFundsInfo(context.Background(), []string{"012414"})
}

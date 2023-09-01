package repository

import (
	"context"
	"live-trading/internal/domain/entity"
)

type FundRepo interface {
	ListFundsInfo(ctx context.Context, fcodes []string) (entity.FundList, error)
}

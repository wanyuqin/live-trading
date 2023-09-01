package repository

import (
	"context"
	"live-trading/internal/domain/entity"
)

type FundRepo interface {
	ListFundsInfo(ctx context.Context, fcodes []string) (entity.FundList, error)
	FundDetail(ctx context.Context, code string) (entity.Fund, error)
	GetFundPosition(ctx context.Context, code string) (entity.FundPosition, error)
	GetFundManager(ctx context.Context, code string) (entity.FundManagerList, error)
	GetFundDetail(ctx context.Context, code string) (entity.FoundDetail, error)
	GetFundSummary(ctx context.Context, code string) (entity.FundSummary, error)
}

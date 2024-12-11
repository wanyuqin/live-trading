package service

import (
	"context"
	"live-trading/internal/configs"
	"live-trading/internal/domain/entity"
	"live-trading/internal/domain/repository"
	"live-trading/internal/infrastructure/dongfang"
	"live-trading/internal/infrastructure/xueqiu"
)

type IFund interface {
	ListFundsInfo(ctx context.Context) (entity.FundList, error)
	AddFundCode(ctx context.Context, code string) error
	DeleteFundCode(ctx context.Context, code string) error
	GetFundDetail(ctx context.Context, code string) (entity.FoundDetail, error)
}

type FundService struct {
	FundRepo   repository.FundRepo
	XueQiuRepo repository.FundRepo
}

func NewFund() *FundService {
	return &FundService{
		FundRepo:   dongfang.NewFundRepoImpl(),
		XueQiuRepo: xueqiu.NewFundRepoImpl(),
	}
}

func (f *FundService) ListFundsInfo(ctx context.Context) (entity.FundList, error) {
	FundList, err := f.FundRepo.ListFundsInfo(ctx, configs.GetConfig().WatchList.Fund)

	return FundList, err
}

func (f *FundService) AddFundCode(ctx context.Context, code string) error {
	return configs.GetConfig().AddFundCode(code)
}

func (f *FundService) DeleteFundCode(ctx context.Context, code string) error {
	return configs.GetConfig().DeleteFundCode(code)
}

func (f *FundService) GetFundDetail(ctx context.Context, code string) (entity.FoundDetail, error) {
	fundDetail, err := f.XueQiuRepo.GetFundDetail(ctx, code)
	if err != nil {
		return fundDetail, err
	}

	fundPosition, err := f.XueQiuRepo.GetFundPosition(ctx, code)
	if err != nil {
		return fundDetail, err
	}

	fundDetail.FundPosition = fundPosition

	fundSummary, err := f.XueQiuRepo.GetFundSummary(ctx, code)
	if err != nil {
		return fundDetail, err
	}

	fundDetail.FundSummary = fundSummary

	return fundDetail, nil
}

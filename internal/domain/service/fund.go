package service

import (
	"context"
	"live-trading/internal/configs"
	"live-trading/internal/domain/entity"
	"live-trading/internal/domain/repository"
	"live-trading/internal/infrastructure/dongfang"
)

type Ifund interface {
	ListfundsInfo(ctx context.Context) (entity.fundList, error)
}

type fundService struct {
	fundRepo repository.fundRepo
}

func Newfund() *fundService {

	return &fundService{
		fundRepo: dongfang.NewDongFangfundRepoImpl(),
	}
}

func (f *fundService) ListfundsInfo(ctx context.Context) (entity.fundList, error) {
	fundList, err := f.fundRepo.ListfundsInfo(ctx, configs.GetConfig().WatchList.fund)

	return fundList, err
}

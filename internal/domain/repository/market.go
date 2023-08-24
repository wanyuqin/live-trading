package repository

type MarketRepo interface {
	ListMarket(res chan<- []byte)
}

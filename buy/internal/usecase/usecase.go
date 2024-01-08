package usecase

import (
	"buy/internal/domain"
)

type BuyUsecase struct {
	buyRepository domain.BuyRepository
}

func NewBuyUsecase(buyRepository domain.BuyRepository) *BuyUsecase {
	return &BuyUsecase{
		buyRepository: buyRepository,
	}
}

func (u *BuyUsecase) CreateBuy(buy *domain.Buy) error {
	return u.buyRepository.CreateBuy(buy)
}

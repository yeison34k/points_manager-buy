package app

import (
	"buy/internal/domain"
	"buy/internal/usecase"
)

type MyApp struct {
	buyUsecase usecase.BuyUsecase
}

func NewMyApp(buyUsecase usecase.BuyUsecase) *MyApp {
	return &MyApp{
		buyUsecase: buyUsecase,
	}
}

func (a *MyApp) HandleRequest(req *domain.Buy) error {
	err := a.buyUsecase.CreateBuy(req)
	if err != nil {
		return err
	}

	return nil
}

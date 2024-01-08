package domain

type BuyRepository interface {
	CreateBuy(buy *Buy) error
}

package usecase_test

import (
	"buy/internal/domain"
	"buy/internal/usecase"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockBuyRepository struct {
}

func (m *MockBuyRepository) CreateBuy(buy *domain.Buy) error {
	return nil
}

func TestBuyUsecase_CreateBuy(t *testing.T) {
	mockRepository := &MockBuyRepository{}
	usecase := usecase.NewBuyUsecase(mockRepository)

	err := usecase.CreateBuy(&domain.Buy{ID: "1", ProductName: "Gasolina"})

	assert.NoError(t, err)
}

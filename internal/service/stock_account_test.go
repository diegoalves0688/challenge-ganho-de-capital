package service

import (
	"testing"

	"challenge/internal/domain"

	"github.com/stretchr/testify/assert"
)

func TestStockAccountService(t *testing.T) {

	t.Run("01_-_test_buy_operation_with_no_tax", func(t *testing.T) {
		stockAccountService := NewStockAccountService(&domain.StockAccount{})

		operation := domain.OperationInput{ Type: domain.BuyOperationType, UnitCost: 20.10, Quantity: 15200}

		output := stockAccountService.Buy(operation)
		assert.Equal(t, domain.OperationOutput{Tax: 0}, output)
	})

	t.Run("02_-_test_sell_operation_with_tax", func(t *testing.T) {
		stockAccountService := NewStockAccountService(&domain.StockAccount{})

		operation := domain.OperationInput{ Type: domain.SellOperationType, UnitCost: 20.11, Quantity: 15200}

		output := stockAccountService.Sell(operation)
		assert.Equal(t, domain.OperationOutput{Tax: 61134.39}, output)
	})

}

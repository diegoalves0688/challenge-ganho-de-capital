package service

import (
	"challenge/internal/domain"
	"fmt"
)

type StockAccount interface {
	Sell(domain.InputOperation) (output domain.OutputLine)
	Buy(domain.InputOperation) (output domain.OutputLine)
	GetBalance()
}

type StockAccountService struct {
	StockAccount *domain.StockAccount
}

func NewStockAccountService(stockAccount *domain.StockAccount) StockAccount {
	return &StockAccountService{StockAccount: stockAccount}
}

func (sa *StockAccountService) Sell(operation domain.InputOperation) (output domain.OutputLine) {
	sa.StockAccount.StockQuantity -= operation.Quantity

	if hasLoss(sa.StockAccount, operation) {
		sa.StockAccount.Loss = (sa.StockAccount.WeightedAveragePrice - operation.UnitCost) * operation.Quantity
	} else {
		currentProfit := (operation.UnitCost - sa.StockAccount.WeightedAveragePrice) * operation.Quantity
		if sa.StockAccount.Loss > 0 && sa.StockAccount.Loss <= currentProfit {
			currentProfit = currentProfit - sa.StockAccount.Loss
			sa.StockAccount.Loss = 0
		} else if sa.StockAccount.Loss > 0 && sa.StockAccount.Loss > currentProfit {
			sa.StockAccount.Loss -= currentProfit
			currentProfit = 0
		}

		if hasTax(sa.StockAccount, operation) {
			output.Tax = (currentProfit / 100) * 20
		}
	}

	return
}

func (sa *StockAccountService) Buy(operation domain.InputOperation) (output domain.OutputLine) {
	sa.StockAccount.WeightedAveragePrice = WeightedAveragePrice(sa.StockAccount.StockQuantity,
		sa.StockAccount.WeightedAveragePrice, operation.Quantity, operation.UnitCost)

	sa.StockAccount.StockQuantity += operation.Quantity
	return
}

func (sa *StockAccountService) GetBalance() {
	fmt.Println("StockAccount.StockQuantity: ", sa.StockAccount.StockQuantity)
	fmt.Println("StockAccount.WeightedAveragePrice: ", sa.StockAccount.WeightedAveragePrice)
	fmt.Println("StockAccount.Loss: ", sa.StockAccount.Loss)
}

func WeightedAveragePrice(currentStockQuantity float64, currentWeightedAverege float64, purchasedStockQuantity float64, purchaseAmount float64) float64 {
	return ((currentStockQuantity * currentWeightedAverege) + (purchasedStockQuantity * purchaseAmount)) / (currentStockQuantity + purchasedStockQuantity)
}

func hasTax(stockAccount *domain.StockAccount, operation domain.InputOperation) bool {
	operationTotalAmount := operation.UnitCost * operation.Quantity
	return operation.UnitCost > stockAccount.WeightedAveragePrice && operationTotalAmount > 20000
}

func hasLoss(stockAccount *domain.StockAccount, operation domain.InputOperation) bool {
	return operation.UnitCost < stockAccount.WeightedAveragePrice
}

package service

import (
	"challenge/internal/domain"
	"math"
)

type StockAccount interface {
	ProccessOperations(operations []domain.OperationInput) (outputLine []domain.OperationOutput)
	Sell(domain.OperationInput) (output domain.OperationOutput)
	Buy(domain.OperationInput) (output domain.OperationOutput)
}

type StockAccountService struct {
	StockAccount *domain.StockAccount
}

func NewStockAccountService(stockAccount *domain.StockAccount) StockAccount {
	return &StockAccountService{StockAccount: stockAccount}
}

func (sa *StockAccountService) ProccessOperations(operations []domain.OperationInput) (outputLine []domain.OperationOutput) {
	for _, operationInput := range operations {
		switch operationInput.Type {
		case domain.BuyOperationType:
			outputElement := sa.Buy(operationInput)
			outputLine = append(outputLine, outputElement)
		case domain.SellOperationType:
			outputElement := sa.Sell(operationInput)
			outputLine = append(outputLine, outputElement)
		default:
		}
	}
	return
}

func (sa *StockAccountService) Sell(operation domain.OperationInput) (output domain.OperationOutput) {
	sa.StockAccount.StockQuantity -= operation.Quantity

	if operationHasLoss(sa.StockAccount, operation) {
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

		if operationHasTax(sa.StockAccount, operation) {
			output.Tax = format((currentProfit / 100) * 20)
		}
	}

	return
}

func (sa *StockAccountService) Buy(operation domain.OperationInput) (output domain.OperationOutput) {
	sa.StockAccount.WeightedAveragePrice = WeightedAveragePrice(sa.StockAccount.StockQuantity,
		sa.StockAccount.WeightedAveragePrice, operation.Quantity, operation.UnitCost)

	sa.StockAccount.StockQuantity += operation.Quantity
	return
}

func WeightedAveragePrice(currentStockQuantity float64, currentWeightedAverege float64,
	purchasedStockQuantity float64, purchaseAmount float64) float64 {
	return format(((currentStockQuantity * currentWeightedAverege) + (purchasedStockQuantity * purchaseAmount)) / (currentStockQuantity + purchasedStockQuantity))
}

func operationHasTax(stockAccount *domain.StockAccount, operation domain.OperationInput) bool {
	operationTotalAmount := operation.UnitCost * operation.Quantity
	return operation.UnitCost > stockAccount.WeightedAveragePrice && operationTotalAmount > 20000
}

func operationHasLoss(stockAccount *domain.StockAccount, operation domain.OperationInput) bool {
	return operation.UnitCost < stockAccount.WeightedAveragePrice
}

func format(tax float64) float64 {
	return math.Floor(tax*100) / 100
}

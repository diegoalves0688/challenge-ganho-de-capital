package service

import (
	"challenge/internal/domain"
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

func (sa *StockAccountService) ProccessOperations(operations []domain.OperationInput) (outputLine []domain.OperationOutput){
	for _, operation := range operations {
		switch operation.Type {
		case domain.BuyOperationType:
			outputElement := sa.Buy(operation)
			outputLine = append(outputLine, outputElement)
		case domain.SellOperationType:
			outputElement := sa.Sell(operation)
			outputLine = append(outputLine, outputElement)
		default:
		}
	}
	return
}

func (sa *StockAccountService) Sell(operation domain.OperationInput) (output domain.OperationOutput) {
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

func (sa *StockAccountService) Buy(operation domain.OperationInput) (output domain.OperationOutput) {
	sa.StockAccount.WeightedAveragePrice = WeightedAveragePrice(sa.StockAccount.StockQuantity,
		sa.StockAccount.WeightedAveragePrice, operation.Quantity, operation.UnitCost)

	sa.StockAccount.StockQuantity += operation.Quantity
	return
}

func WeightedAveragePrice(currentStockQuantity float64, currentWeightedAverege float64,
	purchasedStockQuantity float64, purchaseAmount float64) float64 {
	return ((currentStockQuantity * currentWeightedAverege) + (purchasedStockQuantity * purchaseAmount)) / (currentStockQuantity + purchasedStockQuantity)
}

func hasTax(stockAccount *domain.StockAccount, operation domain.OperationInput) bool {
	operationTotalAmount := operation.UnitCost * operation.Quantity
	return operation.UnitCost > stockAccount.WeightedAveragePrice && operationTotalAmount > 20000
}

func hasLoss(stockAccount *domain.StockAccount, operation domain.OperationInput) bool {
	return operation.UnitCost < stockAccount.WeightedAveragePrice
}

package handler

import (
	"challenge/internal/domain"
	"challenge/internal/service"
	"challenge/pkg"
)

type StockAccountHandler interface {
	ProccessOperationsList([]string) string
	ProccessOperations([]domain.OperationInput) []domain.OperationOutput
}

type StockAccountHandlerImpl struct {
}

func NewStockAccountHandlerImpl() StockAccountHandler {
	return &StockAccountHandlerImpl{}
}

func (sah *StockAccountHandlerImpl) ProccessOperationsList(serializedLines []string) string {
	operationsList, err := pkg.ParseInputLines(serializedLines)
	if err != nil {
		panic(err)
	}

	output := [][]domain.OperationOutput{}
	for _, operations := range operationsList {
		output = append(output, sah.ProccessOperations(operations))
	}

	return pkg.ParseOutputLines(output)
}

func (sah *StockAccountHandlerImpl) ProccessOperations(operations []domain.OperationInput) []domain.OperationOutput {
	return service.NewStockAccountService(&domain.StockAccount{}).ProccessOperations(operations)
}

package handler

import (
	"challenge/internal/domain"
	"challenge/internal/service"
	"challenge/pkg"
	"fmt"
)

type StockAccountHandler interface {
	ProccessOperations()
}

type StockAccountHandlerImpl struct {
}

func NewStockAccountHandlerImpl() StockAccountHandler {
	return &StockAccountHandlerImpl{}
}

func (sah *StockAccountHandlerImpl) ProccessOperations() {
	operationsList, err := pkg.ParseLines(pkg.ReadLines())
	if err != nil {
		panic(err)
	}

	output := [][]domain.OperationOutput{}
	for _, operations := range operationsList {
		stockAccountService := service.NewStockAccountService(&domain.StockAccount{})
		outputLine := stockAccountService.ProccessOperations(operations)
		output = append(output, outputLine)
	}

	fmt.Println("output stdout:")
	fmt.Println(pkg.ParseOutputLines(output))
}

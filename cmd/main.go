package main

import (
	"challenge/internal/domain"
	"challenge/internal/service"
	"challenge/pkg"
	"encoding/json"
	"fmt"
)

func main() {
	fmt.Println("input:")
	//operations, err := pkg.ParseLines(pkg.ReadLines())
	operationsList, err := pkg.ParseLines(pkg.ReadLines())
	if err != nil {
		panic(err)
	}

	/*for _, operation := range operations {
		if operation.Type == "buy" {
			stockAccountService.Buy(operation)
		} else if operation.Type == "sell" {
			stockAccountService.Sell(operation)
		}
	}*/

	output := [][]domain.OutputLine{}
	for _, operations := range operationsList {
		stockAccountService := service.NewStockAccountService(&domain.StockAccount{})
		outputLine := []domain.OutputLine{}
		for _, operation := range operations {
			if operation.Type == "buy" {
				outputElement := stockAccountService.Buy(operation)
				outputLine = append(outputLine, outputElement)
			} else if operation.Type == "sell" {
				outputElement := stockAccountService.Sell(operation)
				outputLine = append(outputLine, outputElement)
			}
		}
		output = append(output, outputLine)
		stockAccountService.GetBalance()
	}

	serializedOutput, _ := json.Marshal(output)
	fmt.Println("output serialized:", string(serializedOutput))

	fmt.Println("output stdout:")
	fmt.Println(pkg.ParseOutputLines(output))
	
}

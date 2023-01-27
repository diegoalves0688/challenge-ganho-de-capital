package handler

import (
	"encoding/json"
	"strings"
	"testing"

	"challenge/internal/domain"

	"github.com/stretchr/testify/assert"
)

func TestStockAccountHandler(t *testing.T) {

	t.Run("01_-_test_stock_account_buy_operation_with_no_tax", func(t *testing.T) {

		operationInput := []domain.OperationInput{}
		operationInput = append(operationInput, domain.OperationInput{Type: domain.BuyOperationType, UnitCost: 20.11, Quantity: 15200})
		operationInput = append(operationInput, domain.OperationInput{Type: domain.BuyOperationType, UnitCost: 40.11, Quantity: 19550})

		serializedInput, _ := json.Marshal(operationInput)
		var serializedLines []string
		serializedLines = append(serializedLines, string(serializedInput))

		expectedOutput := []domain.OperationOutput{}
		expectedOutput = append(expectedOutput, domain.OperationOutput{Tax: 0})
		expectedOutput = append(expectedOutput, domain.OperationOutput{Tax: 0})
		expectedOutputSerialized, _ := json.Marshal(expectedOutput)

		rawOutput := NewStockAccountHandlerImpl().ProccessOperationsList(serializedLines)
		output := strings.Replace(string(rawOutput), "\n", "", -1)
		assert.Equal(t, string(expectedOutputSerialized), output)
	})

	t.Run("02_-_test_stock_account_sell_operation_with_tax", func(t *testing.T) {

		operationInput := []domain.OperationInput{}
		operationInput = append(operationInput, domain.OperationInput{Type: domain.SellOperationType, UnitCost: 20.11, Quantity: 15200})
		operationInput = append(operationInput, domain.OperationInput{Type: domain.SellOperationType, UnitCost: 40.11, Quantity: 19550})

		serializedInput, _ := json.Marshal(operationInput)
		var serializedLines []string
		serializedLines = append(serializedLines, string(serializedInput))

		expectedOutput := []domain.OperationOutput{}
		expectedOutput = append(expectedOutput, domain.OperationOutput{Tax: 61134.39})
		expectedOutput = append(expectedOutput, domain.OperationOutput{Tax: 156830.1})
		expectedOutputSerialized, _ := json.Marshal(expectedOutput)

		rawOutput := NewStockAccountHandlerImpl().ProccessOperationsList(serializedLines)
		output := strings.Replace(string(rawOutput), "\n", "", -1)
		assert.Equal(t, string(expectedOutputSerialized), output)
	})

}

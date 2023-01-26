package pkg

import (
	"encoding/json"
	"fmt"

	"challenge/internal/domain"
)

func ParseLines(serializedLines []string) (operationsList [][]domain.OperationInput, err error) {
	for _, serializedLine := range serializedLines {
		operationsLine := []domain.OperationInput{}
		err = json.Unmarshal([]byte(serializedLine), &operationsLine)
		operationsList = append(operationsList, operationsLine)
	}
	return
}

func ParseOutputLines(outputLineList [][]domain.OperationOutput) (content string) {
	for _, outputLine := range outputLineList {
		serializedOutput, _ := json.Marshal(outputLine)
		content += fmt.Sprintf("%s%s", serializedOutput, "\n")
	}
	return
}

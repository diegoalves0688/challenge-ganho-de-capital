package pkg

import (
	"encoding/json"
	"fmt"
	"strings"

	"challenge/internal/domain"
)

func ParseLiness(serializedLines []string) (lines []domain.InputOperation, err error) {
	segmentLineArr := []domain.InputOperation{}
	serializedArray := ""
	for _, serializedLine := range serializedLines {
		if strings.Contains(serializedLine, "]") {
			serializedArray += serializedLine
			fmt.Println(strings.Replace(serializedArray, "\n", "", -1))
			err = json.Unmarshal([]byte(serializedArray), &segmentLineArr)
			lines = append(lines, segmentLineArr...)
			serializedArray = ""
		} else {
			serializedArray += serializedLine
		}
	}
	return
}

func ParseLines(serializedLines []string) (operationsList [][]domain.InputOperation, err error) {
	for _, serializedLine := range serializedLines {
		operationsLine := []domain.InputOperation{}
		err = json.Unmarshal([]byte(serializedLine), &operationsLine)
		operationsList = append(operationsList, operationsLine)
	}
	return
}

func ParseOutputLines(outputLineList [][]domain.OutputLine) (content string) {
	for _, outputLine := range outputLineList {
		serializedOutput, _ := json.Marshal(outputLine)
		content += fmt.Sprintf("%s%s", serializedOutput, "\n")
	}
	return
}

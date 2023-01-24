package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/diegoalves0688/challenge/internal/domain"
)

func main() {
	fmt.Println("input:")
	lines, err := ParseLines(ReadLines(bufio.NewReader(os.Stdin)))
	if err != nil {
		panic(err)
	}

	for _, line := range lines {
		fmt.Println("output: " + line.Operation)
	}
}

func ReadLines(reader *bufio.Reader) (serializedLines []string) {
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		if len(strings.TrimSpace(line)) == 0 {
			break
		}
		serializedLines = append(serializedLines, line)
	}
	return
}

func ParseLines(serializedLines []string) (lines []domain.InputLine, err error) {
	segmentLineArr := []domain.InputLine{}
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

/*
fmt.Println("input text:")
	reader := bufio.NewReader(os.Stdin)

	var lines []string
	for {
		// read line from stdin using newline as separator
		line, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		// if line is empty, break the loop
		if len(strings.TrimSpace(line)) == 0 {
			break
		}

		//append the line to a slice
		lines = append(lines, line)
	}

	//print out all the lines
	println(len(lines))
	rawInput := ""
	serializedArray := ""

	for _, eachLine := range lines {
		if strings.Contains(eachLine, "]") {
			serializedArray += eachLine
			fmt.Println("output element:")
			fmt.Println(strings.Replace(serializedArray, "\n", "", -1))
			serializedArray = ""
		} else {
			rawInput += eachLine
			serializedArray += eachLine
		}
		rawInput += eachLine
	}
	outputString := strings.Replace(rawInput, "\n", "", -1)
	fmt.Println("output full:")
	fmt.Println(outputString)
*/

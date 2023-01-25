package pkg

import (
	"bufio"
	"os"
	"strings"
)

func ReadLines() (serializedLines []string) {
	scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
		if len(strings.TrimSpace(scanner.Text())) == 0 {
			break
		}
		serializedLines = append(serializedLines, scanner.Text())
    }
	return
}

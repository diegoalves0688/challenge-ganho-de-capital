package test

import (
	"encoding/json"
	"os"
	"os/exec"
	"path"
	"strings"
	"testing"

	"challenge/internal/domain"

	"github.com/stretchr/testify/assert"
)

func TestIntegration(t *testing.T) {

	dir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	t.Run("01_-_test_integration_case_1", func(t *testing.T) {

		expectedOutput := []domain.OperationOutput{}
		expectedOutput = append(expectedOutput, domain.OperationOutput{Tax: 0})
		expectedOutput = append(expectedOutput, domain.OperationOutput{Tax: 0})
		expectedOutput = append(expectedOutput, domain.OperationOutput{Tax: 0})
		expectedOutputSerialized, _ := json.Marshal(expectedOutput)

		cmd := exec.Command(path.Join(dir, "./main"))
		readFile, err := os.ReadFile("cases/case-1.txt")
		if err != nil {
			t.Fatal(err)
		}

		cmd.Stdin = strings.NewReader(string(readFile))
		rawOutput, err := cmd.CombinedOutput()
		output := strings.Replace(string(rawOutput), "\n", "", -1)
		if err != nil {
			t.Fatal(err)
		}

		assert.NoError(t, err)
		assert.Equal(t, string(expectedOutputSerialized), output)
	})

	t.Run("02_-_test_integration_case_2", func(t *testing.T) {

		expectedOutput := []domain.OperationOutput{}
		expectedOutput = append(expectedOutput, domain.OperationOutput{Tax: 0})
		expectedOutput = append(expectedOutput, domain.OperationOutput{Tax: 10000})
		expectedOutput = append(expectedOutput, domain.OperationOutput{Tax: 0})
		expectedOutputSerialized, _ := json.Marshal(expectedOutput)

		cmd := exec.Command(path.Join(dir, "./main"))
		readFile, err := os.ReadFile("cases/case-2.txt")
		if err != nil {
			t.Fatal(err)
		}

		cmd.Stdin = strings.NewReader(string(readFile))
		rawOutput, err := cmd.CombinedOutput()
		output := strings.Replace(string(rawOutput), "\n", "", -1)
		if err != nil {
			t.Fatal(err)
		}

		assert.NoError(t, err)
		assert.Equal(t, string(expectedOutputSerialized), output)
	})

	t.Run("03_-_test_integration_case_3", func(t *testing.T) {

		expectedOutput := []domain.OperationOutput{}
		expectedOutput = append(expectedOutput, domain.OperationOutput{Tax: 0})
		expectedOutput = append(expectedOutput, domain.OperationOutput{Tax: 0})
		expectedOutput = append(expectedOutput, domain.OperationOutput{Tax: 1000})
		expectedOutputSerialized, _ := json.Marshal(expectedOutput)

		cmd := exec.Command(path.Join(dir, "./main"))
		readFile, err := os.ReadFile("cases/case-3.txt")
		if err != nil {
			t.Fatal(err)
		}

		cmd.Stdin = strings.NewReader(string(readFile))
		rawOutput, err := cmd.CombinedOutput()
		output := strings.Replace(string(rawOutput), "\n", "", -1)
		if err != nil {
			t.Fatal(err)
		}

		assert.NoError(t, err)
		assert.Equal(t, string(expectedOutputSerialized), output)
	})

	t.Run("04_-_test_integration_case_4", func(t *testing.T) {

		expectedOutput := []domain.OperationOutput{}
		expectedOutput = append(expectedOutput, domain.OperationOutput{Tax: 0})
		expectedOutput = append(expectedOutput, domain.OperationOutput{Tax: 0})
		expectedOutput = append(expectedOutput, domain.OperationOutput{Tax: 0})
		expectedOutputSerialized, _ := json.Marshal(expectedOutput)

		cmd := exec.Command(path.Join(dir, "./main"))
		readFile, err := os.ReadFile("cases/case-4.txt")
		if err != nil {
			t.Fatal(err)
		}

		cmd.Stdin = strings.NewReader(string(readFile))
		rawOutput, err := cmd.CombinedOutput()
		output := strings.Replace(string(rawOutput), "\n", "", -1)
		if err != nil {
			t.Fatal(err)
		}

		assert.NoError(t, err)
		assert.Equal(t, string(expectedOutputSerialized), output)
	})

	t.Run("05_-_test_integration_case_5", func(t *testing.T) {

		expectedOutput := []domain.OperationOutput{}
		expectedOutput = append(expectedOutput, domain.OperationOutput{Tax: 0})
		expectedOutput = append(expectedOutput, domain.OperationOutput{Tax: 0})
		expectedOutput = append(expectedOutput, domain.OperationOutput{Tax: 0})
		expectedOutput = append(expectedOutput, domain.OperationOutput{Tax: 10000})
		expectedOutputSerialized, _ := json.Marshal(expectedOutput)

		cmd := exec.Command(path.Join(dir, "./main"))
		readFile, err := os.ReadFile("cases/case-5.txt")
		if err != nil {
			t.Fatal(err)
		}

		cmd.Stdin = strings.NewReader(string(readFile))
		rawOutput, err := cmd.CombinedOutput()
		output := strings.Replace(string(rawOutput), "\n", "", -1)
		if err != nil {
			t.Fatal(err)
		}

		assert.NoError(t, err)
		assert.Equal(t, string(expectedOutputSerialized), output)
	})

	t.Run("06_-_test_integration_case_6", func(t *testing.T) {

		expectedOutput := []domain.OperationOutput{}
		expectedOutput = append(expectedOutput, domain.OperationOutput{Tax: 0})
		expectedOutput = append(expectedOutput, domain.OperationOutput{Tax: 0})
		expectedOutput = append(expectedOutput, domain.OperationOutput{Tax: 0})
		expectedOutput = append(expectedOutput, domain.OperationOutput{Tax: 0})
		expectedOutput = append(expectedOutput, domain.OperationOutput{Tax: 3000})
		expectedOutputSerialized, _ := json.Marshal(expectedOutput)

		cmd := exec.Command(path.Join(dir, "./main"))
		readFile, err := os.ReadFile("cases/case-6.txt")
		if err != nil {
			t.Fatal(err)
		}

		cmd.Stdin = strings.NewReader(string(readFile))
		rawOutput, err := cmd.CombinedOutput()
		output := strings.Replace(string(rawOutput), "\n", "", -1)
		if err != nil {
			t.Fatal(err)
		}

		assert.NoError(t, err)
		assert.Equal(t, string(expectedOutputSerialized), output)
	})

	t.Run("07_-_test_integration_case_7", func(t *testing.T) {

		expectedOutput := []domain.OperationOutput{}
		expectedOutput = append(expectedOutput, domain.OperationOutput{Tax: 0})
		expectedOutput = append(expectedOutput, domain.OperationOutput{Tax: 0})
		expectedOutput = append(expectedOutput, domain.OperationOutput{Tax: 0})
		expectedOutput = append(expectedOutput, domain.OperationOutput{Tax: 0})
		expectedOutput = append(expectedOutput, domain.OperationOutput{Tax: 3000})
		expectedOutput = append(expectedOutput, domain.OperationOutput{Tax: 0})
		expectedOutput = append(expectedOutput, domain.OperationOutput{Tax: 0})
		expectedOutput = append(expectedOutput, domain.OperationOutput{Tax: 3700})
		expectedOutput = append(expectedOutput, domain.OperationOutput{Tax: 0})
		expectedOutputSerialized, _ := json.Marshal(expectedOutput)

		cmd := exec.Command(path.Join(dir, "./main"))
		readFile, err := os.ReadFile("cases/case-7.txt")
		if err != nil {
			t.Fatal(err)
		}

		cmd.Stdin = strings.NewReader(string(readFile))
		rawOutput, err := cmd.CombinedOutput()
		output := strings.Replace(string(rawOutput), "\n", "", -1)
		if err != nil {
			t.Fatal(err)
		}

		assert.NoError(t, err)
		assert.Equal(t, string(expectedOutputSerialized), output)
	})

	t.Run("08_-_test_integration_case_8", func(t *testing.T) {

		expectedOutput := []domain.OperationOutput{}
		expectedOutput = append(expectedOutput, domain.OperationOutput{Tax: 0})
		expectedOutput = append(expectedOutput, domain.OperationOutput{Tax: 80000})
		expectedOutput = append(expectedOutput, domain.OperationOutput{Tax: 0})
		expectedOutput = append(expectedOutput, domain.OperationOutput{Tax: 60000})
		expectedOutputSerialized, _ := json.Marshal(expectedOutput)

		cmd := exec.Command(path.Join(dir, "./main"))
		readFile, err := os.ReadFile("cases/case-8.txt")
		if err != nil {
			t.Fatal(err)
		}

		cmd.Stdin = strings.NewReader(string(readFile))
		rawOutput, err := cmd.CombinedOutput()
		output := strings.Replace(string(rawOutput), "\n", "", -1)
		if err != nil {
			t.Fatal(err)
		}

		assert.NoError(t, err)
		assert.Equal(t, string(expectedOutputSerialized), output)
	})

}

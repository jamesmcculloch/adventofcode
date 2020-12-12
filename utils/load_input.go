package utils

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// LoadStringsFromFile returns all lines in a file given the filepath
func LoadStringsFromFile(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return []string{}, err
	}
	defer file.Close()

	lines := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		lines = append(lines, strings.TrimSuffix(line, "\r"))
	}
	return lines, nil
}

// LoadBlankLineSeparatedStringsFromFile returns the blank line separated blocks in a file as a slice of strings
func LoadBlankLineSeparatedStringsFromFile(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return []string{}, err
	}
	defer file.Close()

	blocks := []string{}
	currentBlock := strings.Builder{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			blocks = append(blocks, currentBlock.String())
			currentBlock = strings.Builder{}
			continue
		}

		currentBlock.WriteString(line)
	}
	blocks = append(blocks, currentBlock.String())
	return blocks, nil
}

// LoadNumbersFromInput returns all ints in a file given the filepath
func LoadNumbersFromInput(filePath string) ([]int, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return []int{}, err
	}
	defer file.Close()

	numbers := []int{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSuffix(line, "\r")
		if line == "" {
			break
		}
		number, err := strconv.Atoi(line)
		if err != nil {
			return []int{}, err
		}
		numbers = append(numbers, number)
	}
	return numbers, nil
}

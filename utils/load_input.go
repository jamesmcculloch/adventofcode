package utils

import (
	"bufio"
	"io/ioutil"
	"os"
	"strings"
)

// LoadStringsFromFile returns all lines in a file given the filepath
func LoadStringsFromFile(filePath string) ([]string, error) {
	fileBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return []string{}, err
	}
	fileString := string(fileBytes)
	return strings.Split(strings.TrimSpace(fileString), "\r\n"), nil
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

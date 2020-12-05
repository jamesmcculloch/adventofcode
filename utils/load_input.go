package utils

import (
	"io/ioutil"
	"strings"
)

// LoadStringsFromFile returns all lines in a file given the filepath
func LoadStringsFromFile(filepath string) ([]string, error) {
	fileBytes, err := ioutil.ReadFile(filepath)
	if err != nil {
		return []string{}, err
	}
	fileString := string(fileBytes)
	return strings.Split(strings.TrimSpace(fileString), "\r\n"), nil
}

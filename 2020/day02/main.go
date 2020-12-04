package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type passwordValidator struct {
	password          string
	validationLetter  rune
	minLetterCount    int
	maxLetterCount    int
	firstLetterIndex  int
	secondLetterIndex int
}

func (v passwordValidator) isLetterCountValid() bool {
	letterCount := 0
	for _, letter := range v.password {
		if letter == v.validationLetter {
			letterCount++
		}
	}
	if letterCount >= v.minLetterCount && letterCount <= v.maxLetterCount {
		return true
	}
	return false
}

func (v passwordValidator) isLetterPositionValid() bool {
	letterInCorrectIndexCount := 0
	if rune(v.password[v.firstLetterIndex-1]) == v.validationLetter {
		letterInCorrectIndexCount++
	}

	if rune(v.password[v.secondLetterIndex-1]) == v.validationLetter {
		letterInCorrectIndexCount++
	}

	if letterInCorrectIndexCount != 1 {
		return false
	}
	return true
}

func validLetterCountPasswordCount(passwordsToValidate []passwordValidator) int {
	validPasswords := 0

	for _, candidate := range passwordsToValidate {
		if candidate.isLetterCountValid() {
			validPasswords++
		}
	}

	return validPasswords
}

func validPositionPasswordCount(passwordsToValidate []passwordValidator) int {
	validPasswords := 0

	for _, candidate := range passwordsToValidate {
		if candidate.isLetterPositionValid() {
			validPasswords++
		}
	}

	return validPasswords
}

func getPasswordsToValidate(candidates []string) ([]passwordValidator, error) {
	passwordsToValidate := make([]passwordValidator, len(candidates))
	for index, candidate := range candidates {
		parts := strings.Split(candidate, " ")
		if len(parts) != 3 {
			return []passwordValidator{}, fmt.Errorf("wrong number of parts in candidate: %v", parts)
		}
		letterCounts := parts[0]
		letter := rune(parts[1][0])
		password := parts[2]

		letterCountParts := strings.Split(letterCounts, "-")
		firstPolicyNumberString := letterCountParts[0]
		secondPolicyNumberString := letterCountParts[1]
		firstPolicyNumber, err := strconv.Atoi(firstPolicyNumberString)
		if err != nil {
			return []passwordValidator{}, fmt.Errorf("failed to parse first policy number %s", err.Error())
		}
		secondPolicyNumber, err := strconv.Atoi(secondPolicyNumberString)
		if err != nil {
			return []passwordValidator{}, fmt.Errorf("failed to parse second policy number %s", err.Error())
		}

		passwordToValidate := passwordValidator{
			password:          password,
			validationLetter:  letter,
			minLetterCount:    firstPolicyNumber,
			maxLetterCount:    secondPolicyNumber,
			firstLetterIndex:  firstPolicyNumber,
			secondLetterIndex: secondPolicyNumber,
		}
		passwordsToValidate[index] = passwordToValidate
	}
	return passwordsToValidate, nil
}

func readInput(inputFile string) ([]string, error) {
	bytes, err := ioutil.ReadFile(inputFile)
	if err != nil {
		return []string{}, err
	}
	return strings.Split(string(bytes), "\r\n"), nil
}

func main() {
	lines, err := readInput("input")
	if err != nil {
		fmt.Printf("failed to read input file %s", err.Error())
		return
	}

	candidates, err := getPasswordsToValidate(lines)
	if err != nil {
		fmt.Printf("failed to generate candidates %s", err.Error())
		return
	}

	letterCountPasswordCount := validLetterCountPasswordCount(candidates)
	fmt.Printf("valid password count using letter count method: %d\r\n", letterCountPasswordCount)

	positionPasswordCount := validPositionPasswordCount(candidates)
	fmt.Printf("valid password count using letter count method: %d\r\n", positionPasswordCount)
}

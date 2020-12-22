package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type passport struct {
	fields          map[string]string
	requiredFields  map[string]vaildator
	hairColourRegex *regexp.Regexp
	heightRegex     *regexp.Regexp
	pidRegex        *regexp.Regexp
}

func newPassport() *passport {
	p := &passport{}
	p.pidRegex = regexp.MustCompile(`^\d{9}$`)
	p.hairColourRegex = regexp.MustCompile(`^#[\d|a-f]{6}$`)
	p.heightRegex = regexp.MustCompile(`(\d*)(\w*)`)
	p.requiredFields = map[string]vaildator{
		"byr": p.isBirthYearValid,
		"iyr": p.isIssueYearValid,
		"eyr": p.isExpirationYearValid,
		"hgt": p.isHeightValid,
		"hcl": p.isHairColourValid,
		"ecl": p.isEyeColourValid,
		"pid": p.isPassprtValid,
	}
	p.fields = make(map[string]string)
	return p
}

type vaildator func(string) bool

func (p *passport) isBirthYearValid(birthYear string) bool {
	year, err := strconv.Atoi(birthYear)
	if err != nil {
		return false
	}
	if year < 1920 || year > 2002 {
		return false
	}
	return true
}

func (p *passport) isIssueYearValid(issueYear string) bool {
	year, err := strconv.Atoi(issueYear)
	if err != nil {
		return false
	}
	if year < 2010 || year > 2020 {
		return false
	}
	return true
}

func (p *passport) isExpirationYearValid(expirationYear string) bool {
	year, err := strconv.Atoi(expirationYear)
	if err != nil {
		return false
	}
	if year < 2020 || year > 2030 {
		return false
	}
	return true
}

func (p *passport) isHeightValid(height string) bool {
	matches := p.heightRegex.FindStringSubmatch(height)
	heightMagnitude, err := strconv.Atoi(matches[1])
	if err != nil {
		return false
	}
	heightUnit := matches[2]
	if heightUnit == "" {
		return false
	}
	if ((heightMagnitude < 150 || heightMagnitude > 193) && heightUnit == "cm") ||
		((heightMagnitude < 59 || heightMagnitude > 76) && heightUnit == "in") {
		return false
	}
	return true
}

func (p *passport) isPassprtValid(passportID string) bool {
	if !p.pidRegex.Match([]byte(passportID)) {
		return false
	}
	return true
}

func (p *passport) isHairColourValid(hairColour string) bool {
	if !p.hairColourRegex.Match([]byte(hairColour)) {
		return false
	}
	return true
}

func (p *passport) isEyeColourValid(eyeColour string) bool {
	if eyeColour != "amb" &&
		eyeColour != "blu" &&
		eyeColour != "brn" &&
		eyeColour != "gry" &&
		eyeColour != "grn" &&
		eyeColour != "hzl" &&
		eyeColour != "oth" {
		return false
	}
	return true
}

func (p *passport) isValid1() bool {
	for requiredField := range p.requiredFields {
		value, ok := p.fields[requiredField]
		if !ok {
			return false
		}
		if value == "" {
			return false
		}
	}
	return true
}

func (p *passport) isValid2() bool {
	for requiredField, isValid := range p.requiredFields {
		value, ok := p.fields[requiredField]
		if !ok {
			return false
		}
		if !isValid(value) {
			return false
		}
	}
	return true
}

func (p *passport) addField(key, value string) {
	p.fields[key] = value
}

func loadPassports(filepath string) ([]*passport, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return []*passport{}, err
	}
	defer file.Close()

	passports := []*passport{}
	currentPassport := newPassport()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		passportLine := scanner.Text()
		if passportLine == "" {
			passports = append(passports, currentPassport)
			currentPassport = newPassport()
			continue
		}

		parts := strings.Split(passportLine, " ")
		for _, part := range parts {
			field := strings.Split(part, ":")
			currentPassport.addField(field[0], field[1])
		}
	}
	passports = append(passports, currentPassport)
	return passports, nil
}

func numberOfValidPassports1(passports []*passport) int {
	validCount := 0
	for _, passport := range passports {
		if passport.isValid1() {
			validCount++
		}
	}
	return validCount
}

func numberOfValidPassports2(passports []*passport) int {
	validCount := 0
	for _, passport := range passports {
		if passport.isValid2() {
			validCount++
		}
	}
	return validCount
}

func main() {
	passports, err := loadPassports("input")
	if err != nil {
		panic(err)
	}

	fmt.Printf("part 1 valid passports: %d\n", numberOfValidPassports1(passports))
	fmt.Printf("part 2 valid passports: %d\n", numberOfValidPassports2(passports))
}

package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/jamesmcculloch/adventofcode/utils"
)

func loadRules(transmittedRules []string) []*regexp.Regexp {
	unparsedRules := make([]string, 140)
	for _, transmittedRule := range transmittedRules {
		parts := strings.Split(transmittedRule, ":")
		ruleID, err := strconv.Atoi(parts[0])
		if err != nil {
			panic(err)
		}

		unparsedRules[ruleID] = strings.TrimSpace(parts[1])
	}

	return parseRules(unparsedRules)
}

func parseRules(unparsedRules []string) []*regexp.Regexp {
	rules := make([]*regexp.Regexp, len(unparsedRules))
	for id, logic := range unparsedRules {
		if rules[id] != nil {
			continue
		}

		parseLogic(id, logic, unparsedRules, rules)
	}
	return rules
}

func parseLogic(id int, logic string, unparsedRules []string, rules []*regexp.Regexp) {
	if rules[id] != nil {
		return
	}

	strings.TrimSpace(logic)
	if strings.Contains(logic, "\"") {
		setAorB(id, logic, rules)
		return
	}

	if strings.Contains(logic, "|") {
		orParts := strings.Split(logic, "|")
		if !isIDInLogic(id, orParts[1]) {
			for _, orPart := range orParts {
				loadSubExpressionParts(orPart, unparsedRules, rules)
			}
			regexsForID := make([]string, len(orParts))
			for i, orPart := range orParts {
				regexsForID[i] = getRegexForSubExpression(orPart, rules)
			}
			rules[id] = regexp.MustCompile(`(` + strings.Join(regexsForID, "|") + `)`)
		} else {
			part2Hack(id, orParts, unparsedRules, rules)
		}
	} else {
		loadSubExpressionParts(logic, unparsedRules, rules)
		rules[id] = regexp.MustCompile(`(` + getRegexForSubExpression(logic, rules) + `)`)
	}
}

func loadSubExpressionParts(logic string, unparsedRules []string, rules []*regexp.Regexp) {
	parts := strings.Split(logic, " ")
	for _, idString := range parts {
		if idString == "" {
			continue
		}
		idValue, err := strconv.Atoi(idString)
		if err != nil {
			panic(err)
		}
		parseLogic(idValue, unparsedRules[idValue], unparsedRules, rules)
	}
}

func getRegexForSubExpression(logic string, rules []*regexp.Regexp) string {
	parts := strings.Split(logic, " ")
	regexForID := ""
	for _, idString := range parts {
		if idString == "" {
			continue
		}
		idValue, err := strconv.Atoi(idString)
		if err != nil {
			panic(err)
		}
		regexForID += rules[idValue].String()
	}
	return regexForID
}

func setAorB(id int, logic string, rules []*regexp.Regexp) {
	if strings.Contains(logic, "a") {
		rules[id] = regexp.MustCompile(`a`)
	} else {
		rules[id] = regexp.MustCompile(`b`)
	}
}

func messagesThatMatchRules(messages []string, rule *regexp.Regexp) int {
	completeMatch := regexp.MustCompile(`^` + rule.String() + `$`)
	count := 0
	for _, message := range messages {
		if completeMatch.Match([]byte(message)) {
			count++
		}
	}
	return count
}

func isIDInLogic(id int, logic string) bool {
	parts := strings.Split(logic, " ")
	for _, idString := range parts {
		if idString == "" {
			continue
		}
		idValue, err := strconv.Atoi(idString)
		if err != nil {
			panic(err)
		}
		if idValue == id {
			return true
		}
	}
	return false
}

func part2Hack(id int, orParts []string, unparsedRules []string, rules []*regexp.Regexp) {
	if id == 8 {
		loadSubExpressionParts(orParts[0], unparsedRules, rules)
		left := regexp.MustCompile(`(` + getRegexForSubExpression(orParts[0], rules) + `)`)

		right := left.String() + left.String() + `*`

		rules[id] = regexp.MustCompile(`(` + right + `)`)
	}
	if id == 11 {
		loadSubExpressionParts(orParts[0], unparsedRules, rules)
		left := regexp.MustCompile(`(` + getRegexForSubExpression(orParts[0], rules) + `)`)
		leftPartOne := regexp.MustCompile(`(` + rules[42].String() + `)`)
		leftPartTwo := regexp.MustCompile(`(` + rules[31].String() + `)`)

		recursions := []string{left.String()}

		for i := 1; i < 5; i++ {
			recursionString := ""
			for j := 0; j <= i; j++ {
				recursionString += leftPartOne.String()
			}
			for j := 0; j <= i; j++ {
				recursionString += leftPartTwo.String()
			}
			recursionString = `(` + recursionString + `)`
			recursions = append(recursions, recursionString)
		}

		rules[id] = regexp.MustCompile(`(` + strings.Join(recursions, "|") + `)`)
	}
}

func main() {
	transmission, err := utils.LoadBlankLineSeparatedBlocksOfStringsFromFile("input")
	if err != nil {
		panic(err)
	}

	rules := loadRules(transmission[0])

	fmt.Printf("part 1: %d\n", messagesThatMatchRules(transmission[1], rules[0]))

	transmission, err = utils.LoadBlankLineSeparatedBlocksOfStringsFromFile("input2")
	if err != nil {
		panic(err)
	}

	rules = loadRules(transmission[0])

	fmt.Printf("part 2: %d\n", messagesThatMatchRules(transmission[1], rules[0]))

}

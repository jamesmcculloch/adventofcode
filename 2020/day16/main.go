package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/jamesmcculloch/adventofcode/utils"
)

type rules map[string][]rule

type rule struct {
	upperLimit int
	lowerLimit int
}

func loadRules(ticketRules []string) (rules, error) {
	loadedRules := make(rules)
	for _, ticketRule := range ticketRules {
		parts := strings.Split(ticketRule, ":")
		ruleName := parts[0]
		loadedRules[ruleName] = []rule{}
		ranges := strings.Split(parts[1], "or")
		for _, valueRange := range ranges {
			valueRange = strings.TrimSpace(valueRange)
			limits := strings.Split(valueRange, "-")

			lowerLimit, err := strconv.Atoi(limits[0])
			if err != nil {
				return rules{}, err
			}
			upperLimit, err := strconv.Atoi(limits[1])
			if err != nil {
				return rules{}, err
			}

			loadedRules[ruleName] = append(loadedRules[ruleName], rule{
				lowerLimit: lowerLimit,
				upperLimit: upperLimit,
			})
		}
	}
	return loadedRules, nil
}

type tickets []ticket

type ticket []int

func loadInput(inputFile string) (rules, ticket, tickets, error) {
	ticketRules := make(rules)
	myTicket := ticket{}
	nearbyTickets := tickets{}
	var err error

	file, err := os.Open(inputFile)
	if err != nil {
		return ticketRules, myTicket, nearbyTickets, err
	}
	defer file.Close()

	rulesToLoad := []string{}
	mode := "rules"
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		blockEnd := line == ""
		if blockEnd {
			switch mode {
			case "rules":
				mode = "my ticket"
			case "my ticket":
				mode = "nearby ticket"
			case "nearby ticket":
				break
			}
			continue
		}

		switch mode {
		case "rules":
			rulesToLoad = append(rulesToLoad, line)
		case "my ticket":
			if !!!strings.HasPrefix(line, "your ticket") {
				myTicket, err = utils.GetInts(line, ",")
				if err != nil {
					return ticketRules, myTicket, nearbyTickets, err
				}
			}
		case "nearby ticket":
			if !!!strings.HasPrefix(line, "nearby tickets") {
				nearbyTicket, err := utils.GetInts(line, ",")
				if err != nil {
					return ticketRules, myTicket, nearbyTickets, err
				}
				nearbyTickets = append(nearbyTickets, nearbyTicket)
			}
		}
	}

	loadedRules, err := loadRules(rulesToLoad)
	if err != nil {
		return ticketRules, myTicket, nearbyTickets, err
	}

	return loadedRules, myTicket, nearbyTickets, nil
}

func scanningErrorRate(ticketRules rules, nearbyTickets tickets) (tickets, int) {
	scanningErrorRate := 0
	validTickets := tickets{}
	for _, ticket := range nearbyTickets {
		validTicket := true
		for _, value := range ticket {
			validValue := false
		valid:
			for _, rules := range ticketRules {
				for _, rule := range rules {
					if value >= rule.lowerLimit && value <= rule.upperLimit {
						validValue = true
						break valid
					}
				}
			}
			if !validValue {
				scanningErrorRate += value
				validTicket = false
			}
		}
		if validTicket {
			validTickets = append(validTickets, ticket)
		}
	}
	return validTickets, scanningErrorRate
}

func findRulePositions(ticketRules rules, validTickets tickets) map[string]int {
	rulePotentialPositions := newRulePotentialPositions(ticketRules)
	for valueIndex := 0; valueIndex <= len(validTickets[0])-1; valueIndex++ {
		potentialRulesForValue := newPotentialRules(ticketRules)
		for ticketIndex := 0; ticketIndex <= len(validTickets)-1; ticketIndex++ {
			for ruleName := range potentialRulesForValue {
				value := validTickets[ticketIndex][valueIndex]
				validRule := areRulesValid(value, ticketRules[ruleName])
				if validRule {
					potentialRulesForValue[ruleName]++
				}
			}
		}
		rulesForValue := findValidPotentialRules(potentialRulesForValue, len(validTickets))
		for _, rule := range rulesForValue {
			rulePotentialPositions[rule] = append(rulePotentialPositions[rule], valueIndex)
		}

	}
	finalPositions := map[string]int{}
	findFinalPositions(rulePotentialPositions, finalPositions)

	return finalPositions
}

func findFinalPositions(potentialPositions map[string][]int, rulePositions map[string]int) {
	indexToRemove := -1
	for rule, positions := range potentialPositions {
		if len(positions) == 0 {
			continue
		}
		if len(positions) == 1 {
			rulePositions[rule] = positions[0]
			indexToRemove = positions[0]
		}
	}
	if indexToRemove != -1 {
		for rule, positions := range potentialPositions {
			if len(positions) > 0 {
				potentialPositions[rule] = removeValue(potentialPositions[rule], indexToRemove)
			}
		}
		findFinalPositions(potentialPositions, rulePositions)
	}
}

func removeValue(numbers []int, value int) []int {
	for index, number := range numbers {
		if number == value {
			return append(numbers[:index], numbers[index+1:]...)
		}
	}
	return numbers
}

func areRulesValid(value int, rules []rule) bool {
	for _, rule := range rules {
		if value >= rule.lowerLimit && value <= rule.upperLimit {
			return true
		}
	}
	return false
}

func newRulePotentialPositions(ticketRules rules) map[string][]int {
	potentialRules := map[string][]int{}
	for ruleName := range ticketRules {
		potentialRules[ruleName] = []int{}
	}
	return potentialRules
}

func newPotentialRules(ticketRules rules) map[string]int {
	potentialRules := map[string]int{}
	for ruleName := range ticketRules {
		potentialRules[ruleName] = 0
	}
	return potentialRules
}

func findValidPotentialRules(potentialRules map[string]int, validRuleCount int) []string {
	rules := []string{}
	for rule, count := range potentialRules {
		if count == validRuleCount {
			rules = append(rules, rule)
		}
	}
	return rules
}

func part2(rulePositions map[string]int, myTicket ticket) int {
	result := 1
	for rule, position := range rulePositions {
		if strings.HasPrefix(rule, "departure") {
			result *= myTicket[position]
		}
	}
	return result
}

func main() {
	rules, myTicket, nearbyTickets, err := loadInput("input")
	if err != nil {
		panic(err)
	}

	validTickets, scanningErrorRate := scanningErrorRate(rules, nearbyTickets)
	fmt.Printf("part 1: %d\n", scanningErrorRate)

	validTickets = append(validTickets, myTicket)

	rulePositions := findRulePositions(rules, validTickets)

	fmt.Printf("part 2: %d\n", part2(rulePositions, myTicket))
}

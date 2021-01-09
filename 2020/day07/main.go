package main

import (
	"fmt"
	"strings"

	"github.com/jamesmcculloch/adventofcode/utils"
)

type luggageProcessor struct {
	rules             map[string]map[string]int
	containsTargetBag map[string]bool
	totalBagSize      map[string]int
}

func new() *luggageProcessor {
	l := &luggageProcessor{}
	l.rules = make(map[string]map[string]int)

	return l
}

func (l *luggageProcessor) populateRules(rules []string) {
	for _, rule := range rules {
		parts := strings.Split(rule, "contain")
		outerBag := getOuterBag(parts[0])
		innerBags := strings.Split(parts[1], ",")
		for _, innerBag := range innerBags {
			bag, count := getBagAndCount(innerBag)
			if _, ok := l.rules[outerBag]; !ok {
				l.rules[outerBag] = map[string]int{}
			}
			if bag != "" && count != 0 {
				l.rules[outerBag][bag] = count
			}
		}
	}
}

func (l *luggageProcessor) findBagsThatContainTarget(targetBag string) {
	l.containsTargetBag = make(map[string]bool)
	for outerBag, innerBags := range l.rules {
		if _, ok := l.containsTargetBag[outerBag]; ok {
			continue
		}
		found := l.containsTarget(outerBag, innerBags, targetBag)
		if found {
			l.containsTargetBag[outerBag] = true
		}
	}
}

func (l *luggageProcessor) containsTarget(outerBag string, bags map[string]int, targetBag string) bool {
	for bag := range bags {
		if bag == targetBag {
			l.containsTargetBag[outerBag] = true
			return true
		}
		found := l.containsTarget(bag, l.rules[bag], targetBag)
		if found {
			return true
		}
	}
	return false
}

func getOuterBag(ruleSection string) string {
	var adjective string
	var bagColour string
	fmt.Sscanf(ruleSection, "%s %s bags", &adjective, &bagColour)
	bag := adjective + " " + bagColour
	return bag
}

func getBagAndCount(ruleSection string) (string, int) {
	if ruleSection == "no other bags." {
		return "", 0
	}
	var adjective string
	var bagColour string
	var count int
	fmt.Sscanf(ruleSection, "%d %s %s bag", &count, &adjective, &bagColour)
	bag := adjective + " " + bagColour
	return bag, count
}

func (l *luggageProcessor) totalBagsContainingTarget(targetBag string) int {
	l.findBagsThatContainTarget(targetBag)
	total := 0
	for _, containsTargetBag := range l.containsTargetBag {
		if containsTargetBag {
			total++
		}
	}
	return total
}

func (l *luggageProcessor) findBagSizes() {
	l.totalBagSize = make(map[string]int)

	for outerBag, innerBags := range l.rules {
		if _, ok := l.totalBagSize[outerBag]; ok {
			continue
		}
		l.totalBagSize[outerBag] = l.bagSize(innerBags)
	}
}

func (l *luggageProcessor) bagSize(bags map[string]int) int {
	if len(bags) == 0 {
		return 0
	}
	totalSize := 0
	for bag, count := range bags {
		var size int
		if _, ok := l.totalBagSize[bag]; ok {
			size = l.totalBagSize[bag]
		} else {
			size = l.bagSize(l.rules[bag])
			l.totalBagSize[bag] = size
		}
		totalSize += count + (size * count)
	}
	return totalSize
}

func (l *luggageProcessor) totalBagSizeForTarget(targetBag string) int {
	return l.totalBagSize[targetBag]
}

func main() {
	rules, err := utils.LoadStringsFromFile("input")
	if err != nil {
		panic(err)
	}

	targetBag := "shiny gold"
	lp := new()
	lp.populateRules(rules)

	fmt.Printf("part 1: %d\n", lp.totalBagsContainingTarget(targetBag))
	lp.findBagSizes()
	fmt.Printf("part 2: %d\n", lp.totalBagSizeForTarget(targetBag))
}

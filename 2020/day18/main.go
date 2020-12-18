package main

import (
	"fmt"
	"strconv"

	"github.com/jamesmcculloch/adventofcode/utils"
	"github.com/jamesmcculloch/adventofcode/utils/stack"
)

type mode string

const operand mode = "operand"
const operator mode = "operator"
const calculate mode = "calculate"
const parenthesis mode = "parenthesis"

func sum(expressions []string, additionPreference bool) int {
	total := 0
	for _, expression := range expressions {
		total += evaluate(expression, additionPreference)
	}
	return total
}

func evaluate(expression string, additionPreference bool) int {
	operands := &stack.IntStack{}
	operators := &stack.StringStack{}
	elementBeforeOpenParenthesis := &stack.StringStack{}

	currentMode := operand
	lastElement := ""
	for _, element := range expression {
		if element == ' ' {
			continue
		}

		switch currentMode {
		case operand:
			if string(element) == "(" {
				elementBeforeOpenParenthesis.Push(lastElement)
				currentMode = operand
				lastElement = string(element)
				if additionPreference {
					operators.Push(string(element))
				}
				continue
			}
			value, err := strconv.Atoi(string(element))
			if err != nil {
				panic(err)
			}
			operands.Push(value)
			currentMode = operator
		case operator:
			switch string(element) {
			case "*":
				if !additionPreference {
					operators.Push(string(element))
					currentMode = calculate
					lastElement = string(element)
					continue
				}
				operators.Push(string(element))
				currentMode = operand
				lastElement = string(element)
				continue
			case "+":
				operators.Push(string(element))
				currentMode = calculate
				lastElement = string(element)
			case "(":
				elementBeforeOpenParenthesis.Push(lastElement)
				currentMode = operand
				lastElement = string(element)
				if additionPreference {
					operators.Push(string(element))
				}
				continue
			case ")":
				before, err := elementBeforeOpenParenthesis.Pop()
				if err != nil {
					panic(err)
				}
				if !additionPreference {
					if before != "(" {
						evaluateSubExpression(operands, operators)
					}
				} else {
					evaluateRemainingSubexpressions(operands, operators, false)
					if before == "+" {
						evaluateSubExpression(operands, operators)
					}
				}

			}
		case calculate:
			if string(element) == "(" {
				elementBeforeOpenParenthesis.Push(lastElement)
				currentMode = operand
				lastElement = string(element)
				if additionPreference {
					operators.Push(string(element))
				}
				continue
			}
			value, err := strconv.Atoi(string(element))
			if err != nil {
				panic(err)
			}
			operands.Push(value)
			evaluateSubExpression(operands, operators)
			currentMode = operator
		}
	}
	evaluateRemainingSubexpressions(operands, operators, true)
	result, err := operands.Pop()
	if err != nil {
		panic(err)
	}
	return result
}

func evaluateRemainingSubexpressions(operands *stack.IntStack, operators *stack.StringStack, final bool) {
	for {
		operator, err := operators.Pop()
		if err != nil {
			return
		}
		switch operator {
		case "+":
			left, err := operands.Pop()
			if err != nil {
				panic(err)
			}
			right, err := operands.Pop()
			if err != nil {
				operands.Push(left)
				operators.Push(operator)
				return
			}
			operands.Push(left + right)
		case "*":
			left, err := operands.Pop()
			if err != nil {
				panic(err)
			}
			right, err := operands.Pop()
			if err != nil {
				operands.Push(left)
				operators.Push(operator)
				return
			}
			operands.Push(left * right)
		case "(":
			if !final {
				return
			}
			continue
		}
	}
}

func evaluateSubExpression(operands *stack.IntStack, operators *stack.StringStack) {
	operator, err := operators.Pop()
	if err != nil {
		return
	}
	switch operator {
	case "+":
		left, err := operands.Pop()
		if err != nil {
			panic(err)
		}
		right, err := operands.Pop()
		if err != nil {
			operands.Push(left)
			operators.Push(operator)
			return
		}
		operands.Push(left + right)
	case "*":
		left, err := operands.Pop()
		if err != nil {
			panic(err)
		}
		right, err := operands.Pop()
		if err != nil {
			operands.Push(left)
			operators.Push(operator)
			return
		}
		operands.Push(left * right)
	case "(":
		return
	}
}

func main() {
	expressions, err := utils.LoadStringsFromFile("input")
	if err != nil {
		panic(err)
	}

	fmt.Printf("part 1: %d\n", sum(expressions, false))
	fmt.Printf("part 2: %d\n", sum(expressions, true))
}

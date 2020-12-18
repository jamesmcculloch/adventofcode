package stack

import "errors"

// IntStack push down stack of ints
type IntStack struct {
	head *intElement
}

type intElement struct {
	value int
	next  *intElement
}

// Push add value to stack
func (s *IntStack) Push(value int) {
	newElement := &intElement{
		value: value,
	}

	if s.head != nil {
		newElement.next = s.head
	}
	s.head = newElement
}

// Pop remove element from stack
func (s *IntStack) Pop() (int, error) {
	if s.head == nil {
		return 0, errors.New("empty stack")
	}
	returnValue := s.head.value
	s.head = s.head.next
	return returnValue, nil
}

// StringStack push down stack of ints
type StringStack struct {
	head *stringElement
}

type stringElement struct {
	value string
	next  *stringElement
}

// Push add value to stack
func (s *StringStack) Push(value string) {
	newElement := &stringElement{
		value: value,
	}

	if s.head != nil {
		newElement.next = s.head
	}
	s.head = newElement
}

// Pop remove element from stack
func (s *StringStack) Pop() (string, error) {
	if s.head == nil {
		return "", errors.New("empty stack")
	}
	returnValue := s.head.value
	s.head = s.head.next
	return returnValue, nil
}

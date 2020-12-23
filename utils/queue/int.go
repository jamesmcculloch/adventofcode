package queue

import (
	"errors"
	"strconv"
	"strings"
)

// IntQueue a queue of ints
type IntQueue struct {
	length int
	back   *element
	front  *element
}

type element struct {
	value int
	next  *element
	// prev  *element
}

// Enqueue add int to queue
func (q *IntQueue) Enqueue(value int) {
	q.length++
	newElement := &element{
		value: value,
	}

	if q.IsEmpty() {
		q.front = newElement
		q.back = newElement
		return
	}

	q.back.next = newElement
	q.back = newElement
}

// Dequeue remove int from queue
func (q *IntQueue) Dequeue() (int, error) {
	if q.IsEmpty() {
		return 0, errors.New("empty queue")
	}
	q.length--

	returnElement := q.front
	q.front = q.front.next
	if q.front == nil {
		q.back = nil
	}
	return returnElement.value, nil
}

// IsEmpty returns true if there are no elements in the queue
func (q *IntQueue) IsEmpty() bool {
	if q.front == nil {
		return true
	}
	return false
}

// Length returns the number of elements in the queue
func (q *IntQueue) Length() int {
	return q.length
}

// String returns a string representation of the queue
func (q *IntQueue) String() string {
	if q.IsEmpty() {
		return ""
	}

	var sb strings.Builder
	currentElement := q.front
	sb.WriteString(strconv.Itoa(currentElement.value))
	for {
		currentElement = currentElement.next
		if currentElement == nil {
			break
		}
		sb.WriteString("<-")
		sb.WriteString(strconv.Itoa(currentElement.value))
	}
	return sb.String()
}

// Duplicate creates a new queue containing the first n elements from this queue
func (q *IntQueue) Duplicate(n int) (*IntQueue, error) {
	newQueue := &IntQueue{}
	if n == 0 {
		return newQueue, nil
	}
	if q.IsEmpty() {
		return nil, errors.New("cannot duplicate empty qu")
	}

	currentElement := q.front
	newQueue.Enqueue(currentElement.value)
	count := 1
	if count == n {
		return newQueue, nil
	}
	for {
		currentElement = currentElement.next
		if currentElement == nil {
			return nil, errors.New("not enough elements to copy in queue")
		}
		newQueue.Enqueue(currentElement.value)
		count++
		if count == n {
			break
		}
	}

	return newQueue, nil
}

// Go (Golang) stack (LIFO) implantation. It can store any type of data type
// author: github.com/tiltwave

package stack

import (
	"errors"
)

type Stack []interface{}

func (stack *Stack) Push(data interface{}) {
	*stack = append(*stack, data)
}

func (stack *Stack) Pop() (interface{}, error) {
	if stack.IsEmpty() {
		return nil, errors.New("The stack is empty")
	}
	l := len(*stack)
	// Get the item from the top
	data := (*stack)[l-1]
	// Remove the top item
	*stack = (*stack)[:l-1]

	return data, nil
}

func (stack Stack) Peek() (interface{}, error) {
	if stack.IsEmpty() {
		return nil, errors.New("The stack is empty")
	}

	return stack[len(stack)-1], nil
}

func (stack *Stack) Len() int {
	return len(*stack)
}

func (stack *Stack) IsEmpty() bool {
	if len(*stack) < 1 {
		return true
	}

	return false
}

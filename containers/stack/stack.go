package stack

import "fmt"

type Stack[T any] struct {
	items []T
}

func (stack *Stack[T]) Push(value T) {
	stack.items = append(stack.items, value)
}

func (stack *Stack[T]) Pop() (*T, error) {
	if stack.IsEmpty() {
		return nil, fmt.Errorf("stack is empty")
	}
	item := stack.items[len(stack.items)-1]
	stack.items = stack.items[:len(stack.items)-1]
	return &item, nil
}

func (stack *Stack[T]) Peek() (*T, error) {
	if stack.IsEmpty() {
		return nil, fmt.Errorf("stack is empty")
	}
	return &stack.items[len(stack.items)-1], nil
}

func (stack *Stack[T]) IsEmpty() bool {
	return len(stack.items) == 0
}

func (stack *Stack[T]) Size() int {
	return len(stack.items)
}

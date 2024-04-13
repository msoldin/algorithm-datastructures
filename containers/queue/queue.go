package queue

import "fmt"

type Queue[T any] struct {
	items []T
}

func (queue *Queue[T]) Enqueue(item T) {
	queue.items = append(queue.items, item)
}

func (queue *Queue[T]) Dequeue() (*T, error) {
	if queue.IsEmpty() {
		return nil, fmt.Errorf("queue is empty")
	}
	item := queue.items[queue.Size()-1]
	queue.items = queue.items[:queue.Size()-1]
	return &item, nil
}

func (queue *Queue[T]) Peek() (*T, error) {
	if queue.IsEmpty() {
		return nil, fmt.Errorf("queue is empty")
	}
	return &queue.items[queue.Size()-1], nil
}

func (queue *Queue[T]) IsEmpty() bool {
	return len(queue.items) == 0
}

func (queue *Queue[T]) Size() int {
	return len(queue.items)
}

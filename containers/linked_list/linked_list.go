package linked_list

import (
	"fmt"
	"reflect"
)

type node[T any] struct {
	value T
	next  *node[T]
}

type LinkedList[T any] struct {
	head *node[T]
}

func (linkedList *LinkedList[T]) RemoveByIndex(index int) {
	current, previous := linkedList.head, linkedList.head

	for i := 0; current != nil; i++ {
		if i == index {
			if current == linkedList.head {
				linkedList.head = current.next
				return
			}
			previous.next = current.next
			break
		}
		index++
		current = current.next
	}
}

func (linkedList *LinkedList[T]) Remove(value T) {
	current, previous := linkedList.head, linkedList.head
	if current == nil {
		return
	}

	for current != nil {
		if reflect.DeepEqual(current.value, value) {
			if current == linkedList.head {
				linkedList.head = current.next
				return
			}
			previous.next = current.next
		}
		previous = current
		current = current.next
	}
}

func (linkedList *LinkedList[T]) Push(value T) {
	if linkedList.head == nil {
		linkedList.head = &node[T]{value, nil}
		return
	}
	newNode := &node[T]{value, nil}
	newNode.next = linkedList.head
	linkedList.head = newNode
}

func (linkedList *LinkedList[T]) PushBack(value T) {
	if linkedList.head == nil {
		linkedList.head = &node[T]{value, nil}
		return
	}
	linkedList.tail().next = &node[T]{value, nil}
}

func (linkedList *LinkedList[T]) tail() *node[T] {
	tail := linkedList.head
	for tail.next != nil {
		tail = tail.next
	}
	return tail
}

func (linkedList *LinkedList[T]) Tail() T {
	return linkedList.tail().value
}

func (linkedList *LinkedList[T]) ToString() string {
	r := "[ "
	for currentNode := linkedList.head; currentNode != nil; currentNode = currentNode.next {
		r += fmt.Sprintf("%v ", currentNode.value)
	}
	r += "]"

	return r
}

func New[T any]() *LinkedList[T] {
	return &LinkedList[T]{head: nil}
}

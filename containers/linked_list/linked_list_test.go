package linked_list

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLinkedList(t *testing.T) {
	// Create a new linked list
	list := New[int]()
	list.PushBack(1)
	list.PushBack(2)
	list.PushBack(3)
	list.PushBack(4)
	list.PushBack(5)
	list.PushBack(6)
	list.Push(99)
	list.Remove(4)
	list.RemoveByIndex(0)

	// Verify the linked list contents
	println(list.ToString())
	current := list.head
	expectedIntList := []int{1, 2, 3, 5, 6}
	for _, i := range expectedIntList {
		assert.Equal(t, i, current.value, "linked list does not contain expected element")
		current = current.next
	}
}

package hash_map

import (
	"fmt"
	"testing"
)

func TestHashMap(t *testing.T) {
	// Create a new linked list

	m := New[string, string]()
	m.Put("1", "hier")
	m.Put("2", "hier2")
	m.Put("3", "hier3")
	m.Put("4", "hier4")
	m.Put("5", "hier4")
	m.Put("6", "hier4")
	m.Put("7", "hier4")
	m.Put("8", "hier4")
	m.Put("9", "hier4")
	m.Put("10", "hier4")
	fmt.Printf("Tes")
}

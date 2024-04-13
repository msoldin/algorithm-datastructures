package binary_search

import "testing"

func TestSearch(t *testing.T) {
	tests := []struct {
		name     string
		arr      []int
		elem     int
		expected int
	}{
		{"Element found at the beginning", []int{1, 2, 3, 4, 5}, 1, 0},
		{"Element found at the end", []int{1, 2, 3, 4, 5}, 5, 4},
		{"Element found in the middle", []int{1, 2, 3, 4, 5}, 3, 2},
		{"Element not found", []int{1, 2, 3, 4, 5}, 6, -1},
		{"Empty slice", []int{}, 6, -1},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := search(test.arr, test.elem)
			if result != test.expected {
				t.Errorf("Expected index %d, but got %d", test.expected, result)
			}
		})
	}
}

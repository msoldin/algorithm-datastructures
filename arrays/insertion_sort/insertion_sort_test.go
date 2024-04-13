package insertion_sort

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	gosort "sort"
	"testing"
)

func Test_sort(t *testing.T) {
	//given
	arr := make([]int, 100)
	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(1000)
	}

	//when
	sort(arr)

	//then
	assert.True(t, gosort.IntsAreSorted(arr), "sorted array is not sorted")
}

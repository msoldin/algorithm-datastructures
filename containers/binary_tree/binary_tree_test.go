package binary_tree

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBinaryTree(t *testing.T) {
	tree := New[string]()
	tree.Add("1")
	tree.Add("2")
	tree.Add("3")
	tree.Print()
	tree.Invert()
	fmt.Sprintln(tree.Depth())

	assert.True(t, tree.root.value == "1")
}
